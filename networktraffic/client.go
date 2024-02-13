package main

import (
	"log"
	"time"

	"github.com/ebusto/xbee"
)

func SendMessagesToServer() {
	// Open the XBee module for communication
	xbeeModule, err := xbee.Open("/dev/ttyUSB0", 9600)
	if err != nil {
		log.Fatal("Error opening XBee module:", err)
	}
	defer xbeeModule.Close()

	// Configure XBee module as a client

	for {
		// Send a message to the server
		message := "Hello from XBee client"
		xbeeModule.Send([]byte(message))

		time.Sleep(5 * time.Second) // Send message every 5 seconds
	}
}
