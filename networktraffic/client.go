package networking

import (
	"bufio"
	"fmt"
	"go.bug.st/serial"
	"log"
	"os"
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	message, err := reader.ReadString('\n')
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, "Error reading from input:", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
	for {

		// Send a message to the server
		//The controller will search until it finds a /n character in the message string
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
