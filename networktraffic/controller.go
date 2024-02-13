package main

import (
	"bufio"
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
	defer func(port serial.Port) {
		err := port.Close()
		if err != nil {

		}
	}(port) // Ensure the port is closed when the function returns

	// Wrap the port in a bufio.Reader
	reader := bufio.NewReader(port)

	fmt.Println("Waiting for incoming messages...")
	for {
		// Use ReadBytes or ReadString to dynamically handle incoming data
		// For example, reading until a newline character (adjust as needed)
		message, err := reader.ReadBytes('\n') // or reader.ReadString('\n')       // The controller will search until it finds a /n character in the message string
		if err != nil {
			if err == io.EOF {
				// End of file (or stream) reached, could handle differently if needed
				continue
			} else {
				log.Fatal("Error receiving message:", err)
			}
		}

		// Process the received message
		fmt.Printf("Received message from client: %s", message) // Adjust printing based on ReadBytes or ReadString
	}
}

func main() {
	ConfigureController()
}
