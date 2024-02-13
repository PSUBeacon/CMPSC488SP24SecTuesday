package main

import (
	"fmt"
	"go.bug.st/serial"
	"io"
	"log"
)

func ConfigureController() {
	// Open the XBee module for communication
	mode := &serial.Mode{
		BaudRate: 9600,
	}
	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		log.Fatal("Error opening XBee module:", err)
	}
	// Configure XBee module as a controller

	fmt.Println("Waiting for incoming messages...")
	for {
		// Receive messages from clients
		message := make([]byte, 128)
		n, err := port.Read(message)
		if err != nil {
			if err != io.EOF {
				log.Fatal("Error receiving message:", err)
			}
			continue
		}
		fmt.Printf("Received message from client: %s\n", string(message[:n]))
	}
}

func main() {
	ConfigureController()
}
