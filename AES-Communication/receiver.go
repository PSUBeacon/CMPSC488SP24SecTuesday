package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"github.com/joho/godotenv"
	"go.bug.st/serial"
	"io"
	"log"
	"os"
)

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

	// Wrap the port in a bufio.Reader
	reader := bufio.NewReader(port)

	fmt.Println("Waiting for incoming messages...")
	for {
		message, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				log.Fatal("Error receiving message:", err)
			}
		}

		// Trim the newline character
		message = bytes.TrimRight(message, "\n")
		fmt.Println("The length after trimming is: ", len(message))

		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}

		AesKey := os.Getenv("AES_KEY")
		decryptedText, err := decryptAES([]byte(AesKey), message)
		if err != nil {
			fmt.Println("Error decrypting:", err)
			continue // Continue listening for new messages
		}
		fmt.Printf("Decrypted text: %s\n", decryptedText)
	}
}

func main() {
	ConfigureController()
}
