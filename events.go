package main

import "fmt"

type Event struct {
	Message  string // TODO: change message type
	Response chan string
}

var EventChan chan Event

func init() {
	EventChan = make(chan Event)
}

func GetEventChan() chan<- Event {
	return EventChan
}

// ProcessEvents only listening form EventChan so will be processing events sequentially
// So, Isolation property of ACID will be taken care here
func ProcessEvents() {
	for {
		event := <-EventChan
		fmt.Printf("Incoming Message %s", event.Message)
		event.Response <- "OK"
	}
}
