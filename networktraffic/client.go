package main

import (
	"fmt"
	"go.bug.st/serial"
	"log"
	"time"
)

func SendMessagesToServer() {
	// Open the XBee module for communication
	mode := &serial.Mode{
		BaudRate: 9600,
	}
	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		log.Fatal("Error opening XBee module:", err)
	}

	//sender := xbee.NewSender(port)
	// Configure XBee module as a client

	for {
		// Send a message to the server
		message := "Hello\n" // The controller will search until it finds a /n character in the message string
		_, err := port.Write([]byte(message))
		fmt.Printf("Sent \n")
		if err != nil {
			log.Println("Error sending message:", err)
		}
		time.Sleep(5 * time.Second) // Send message every 5 seconds
	}
}

func main() {
	SendMessagesToServer()
}
