package main

import (
	"fmt"
	"log"

	"github.com/ebusto/xbee"
)

func ConfigureController() {
	// Open the XBee module for communication
	xbeeModule, err := xbee.Open("/dev/ttyUSB1", 9600)
	if err != nil {
		log.Fatal("Error opening XBee module:", err)
	}
	defer xbeeModule.Close()

	// Configure XBee module as a controller

	fmt.Println("Waiting for incoming messages...")
	for {
		// Receive messages from clients
		message, err := xbeeModule.Receive()
		if err != nil {
			log.Fatal("Error receiving message:", err)
		}
		fmt.Printf("Received message from client: %s\n", string(message.Data))
	}
}
