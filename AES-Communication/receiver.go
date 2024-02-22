package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.bug.st/serial"
	"io"
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

func decryptAES(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	fmt.Printf("Ciphertext length (including IV): %d\n", len(ciphertext))

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	fmt.Printf("Ciphertext length (excluding IV): %d\n", len(ciphertext))

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Remove padding
	padding := int(ciphertext[len(ciphertext)-1])
	if padding < 1 || padding > aes.BlockSize {
		return nil, fmt.Errorf("invalid padding")
	}
	for i := len(ciphertext) - padding; i < len(ciphertext); i++ {
		if ciphertext[i] != byte(padding) {
			return nil, fmt.Errorf("invalid padding")
		}
	}
	ciphertext = ciphertext[:len(ciphertext)-padding]

	return ciphertext, nil
}

func ConfigureController() {
	// Open the XBee module for communication
	var block []Block
	mode := &serial.Mode{
		BaudRate: 9600,
	}
	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		log.Fatal("Error opening XBee module:", err)
	}
	defer port.Close()

	// Wrap the port in a bufio.Reader
	reader := bufio.NewReader(port)

	fmt.Println("Waiting for incoming messages...")
	for {
		// Read until the delimiter
		message, err := reader.ReadBytes('*')
		if err != nil {
			if err != io.EOF {
				log.Fatal("Error receiving message:", err)
			}
			// Handle EOF if necessary
		}

		// Trim the delimiter and newline character
		message = bytes.TrimRight(message, "*")

		// Process the message
		fmt.Println("Received message:", string(message))

		// Decrypt the message
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}

		AesKey := os.Getenv("AES_KEY")
		decryptedText, err := decryptAES([]byte(AesKey), message)
		if err != nil {
			log.Fatal("Error decrypting:", err)
		}
		fmt.Printf("Decrypted text: %s\n", decryptedText)

		// Unmarshal the JSON object possibly
		err = json.Unmarshal(decryptedText, &block)
		if err != nil {
			log.Fatal("Error unmarshaling JSON:", err)
		}

		for _, b := range block {
			fmt.Printf("%d - %s - %s - %s - %s\n", b.Index, b.Timestamp, b.Data, b.PrevHash, b.Hash)
		}
	}
}

func main() {
	ConfigureController()
}
