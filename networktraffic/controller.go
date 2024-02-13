package main

import (
	"fmt"
	"io"
	"log"

	//"github.com/ebusto/xbee"
	"go.bug.st/serial"
)

func ConfigureController() {
	// Open the XBee module for communication
	mode := &serial.Mode{
		BaudRate: 9600,
	}
	xbeeModule, err := serial.Open("/dev/ttyUSB1", mode)
	if err != nil {
		log.Fatal("Error opening XBee module:", err)
	}
	// Configure XBee module as a controller

	fmt.Println("Waiting for incoming messages...")
	for {
		// Receive messages from clients
		message := make([]byte, 128)
		n, err := xbeeModule.Read(message)
		if err != nil {
			if err != io.EOF {
				log.Fatal("Error receiving message:", err)
			}
			continue
		}
		fmt.Printf("Received message from client: %s\n", string(message[:n]))
	}
}
