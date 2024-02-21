package main

import (
	"CMPSC488SP24SecTuesday/blockchain"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.bug.st/serial"
	"io"
	"log"
	"os"
	"time"
)

func encryptAES(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// PKCS#7 padding
	padding := block.BlockSize() - len(plaintext)%block.BlockSize()
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintext = append(plaintext, padText...)

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	fmt.Printf("Ciphertext length after encryption: %d\n", len(ciphertext))

	return ciphertext, nil
}

func SendMessagesToServer() {

	// The key should be 16, 24, or 32 bytes long for AES-128, AES-192, or AES-256, respectively.

	err := godotenv.Load()
	AesKey := os.Getenv("AES_KEY")

	// The message to be encrypted.
	// Create a new blockchain and add a block
	blockMessage := blockchain.NewBlockchain()
	//blockMessage.CreateBlock("This used block chain")

	// Convert the blockchain to JSON
	blockchainJSON, err := json.MarshalIndent(blockMessage, "", "  ")
	if err != nil {
		log.Fatal("Error marshalling blockchain:", err)
	}
	// Encrypt the blockchain JSON
	encryptedBlock, err := encryptAES([]byte(AesKey), blockchainJSON)
	if err != nil {
		log.Fatal("Error encrypting block:", err)
	}
	//encryptedBlock := blockchainJSON
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
	//sendmessage := append(encryptedBlock, '\n')
	for {
		// Send a message to the server
		fmt.Println(len(encryptedBlock))
		_, err := port.Write(encryptedBlock)
		_, err = port.Write([]byte("\n"))
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
