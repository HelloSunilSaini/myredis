package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/hellosunilsaini/myredis/config"
)

var currentConn int
var connLock sync.Mutex

func AddConnection(conn net.Conn) {
	conf := config.GetConfig()
	connLock.Lock()
	if currentConn >= conf.MaxConnections {
		fmt.Println("Connection limit reached. Rejecting new connection.")
		conn.Close()
	} else {
		currentConn += 1
		go HandleConnection(conn)
	}
	connLock.Unlock()
}

// removeConnection removes a client connection from the list
func RemoveConnection(conn net.Conn) {
	connLock.Lock()
	conn.Close()
	currentConn -= 1
	connLock.Unlock()
}

// handleConnection handles a single client connection
func HandleConnection(conn net.Conn) {
	responseChan := make(chan string)
	eventChan := GetEventChan()
	defer func() {
		RemoveConnection(conn) // Remove connection from the list
	}()

	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("New connection established from %s\n", clientAddr)
	reader := bufio.NewReader(conn)

	for {
		// TODO: handle Idle connection timeout
		// TODO: changes required as per redis-cli inputs, read and parse connection commands
		// TODO: parsing have to be done here and for parsing errors result can be thrown directly from here
		// without sending it for processing
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Connection closed by %s\n", clientAddr)
			return
		}

		message = strings.TrimSpace(message)
		fmt.Printf("Received from %s: %s\n", clientAddr, message)

		// Sending message to the incoming channel
		event := Event{
			Message:  message,
			Response: responseChan,
		}
		eventChan <- event
		// will wait on channel for response in sync as for any single client request response have to be in sync
		response := <-responseChan
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Printf("Error sending response to %s: %v\n", clientAddr, err)
			return
		}
	}
}
