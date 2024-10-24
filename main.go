package main

import (
	"fmt"
	"log" // TODO: change logger
	"net"

	"github.com/hellosunilsaini/myredis/config"
	"github.com/hellosunilsaini/myredis/connections"
	"github.com/hellosunilsaini/myredis/events"
)

func main() {
	conf := config.GetConfig()
	// Running events processor
	go events.ProcessEvents()
	// Create a new server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.ServerPort))
	if err != nil {
		log.Fatalf("Error starting TCP server: %v", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")

	for {
		// Accept a new connection
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		connections.AddConnection(conn)
	}
}
