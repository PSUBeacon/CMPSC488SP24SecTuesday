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
	}(port)
	// Configure XBee module as a controller
	reader := bufio.NewReader(port)
	fmt.Println("Waiting for incoming messages...")
	for {
		// Receive messages from clients
		message := make([]byte, 128)
		n, err := reader.ReadString('\n')
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
