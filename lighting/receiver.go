package main

import (
	"bufio"
	"fmt"
	"go.bug.st/serial"
	"io"
	"log"
	"strings"
)

func receiverController() {
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
			log.Println("Error closing serial port:", err)
		}
	}(port) // Ensure the port is closed when the function returns

	// Wrap the port in a bufio.Reader for easier data handling
	reader := bufio.NewReader(port)

	fmt.Println("Waiting for incoming messages...")
	for {
		// Reading until a newline character
		message, err := reader.ReadString('\n') // Adjust according to your data format
		if err != nil {
			if err == io.EOF {
				// Stream ended (not typical for serial ports but can happen)
				continue
			} else {
				log.Fatal("Error receiving message:", err)
			}
		}

		// Trim the newline character for easier processing
		message = strings.TrimSpace(message)

		// Parse and act on the message
		// Here you would add your logic to handle different commands
		fmt.Printf("Received message: %s\n", message) // Placeholder for actual handling logic
	}
}

func main() {
	receiverController() // Start the receiver controller to listen for messages
}
