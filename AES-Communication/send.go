package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"go.bug.st/serial"
	"io"
	"log"
	"time"
)

func encryptAES(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func SendMessagesToServer() {

	// The key should be 16, 24, or 32 bytes long for AES-128, AES-192, or AES-256, respectively.
	key := []byte("890fa9277f40e9394dc80e53b203f952") //This key is for testing, will be switched later

	// The message to be encrypted.
	plaintext := []byte("This encrypted everything")

	message, _ := encryptAES(key, plaintext)

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

		_, err := port.Write(message)
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
