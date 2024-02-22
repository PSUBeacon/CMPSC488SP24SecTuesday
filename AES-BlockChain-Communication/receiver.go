package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.bug.st/serial"
	"log"
	"os"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

type Blockchain struct {
	Chain []Block
}

func blockReceiver() {
	// Open the XBee module for communication
	var chain Blockchain
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
	const bufferSize = 4096 // Adjust this value as needed
	reader := bufio.NewReaderSize(port, bufferSize)

	fmt.Println("Waiting for incoming messages...")
	// Use ReadBytes or ReadString to dynamically handle incoming data
	for {
		// Read and parse the data manually
		var message []byte
		for {
			b, err := reader.ReadByte()
			if err != nil {
				log.Fatal("Error reading byte:", err)
			}
			// Check for the UTF-8 encoding of '♄' the hex value is (E2 99 B4)
			if len(message) >= 2 && message[len(message)-2] == 0xE2 && message[len(message)-1] == 0x99 && b == 0xB4 {
				message = message[:len(message)-2] // Remove the delimiter from the message
				break
			}
			message = append(message, b)
		}

		err = godotenv.Load()
		AesKey := os.Getenv("AES_KEY") //This key is for testing, will be switched later
		//Decrypt the message.
		decryptedText, err := decryptAES([]byte(AesKey), message)
		//decryptedText := message

		if err != nil {
			fmt.Println("Error decrypting:", err)
			return
		}
		fmt.Printf("Decrypted text: %s\n", decryptedText)

		tojson := json.Unmarshal(decryptedText, &chain)
		if tojson != nil {

			// if error is not nil
			// print error
			fmt.Println(tojson)
		}

		fmt.Println(chain)
		//for i := range chain {
		//	fmt.Println(string(rune(block[i].Index)) + " - " + block[i].Timestamp + " - " + block[i].Data + " - " + block[i].PrevHash + " - " + block[i].Hash)
		//}
	}
}

func main() {
	blockReceiver()
}
