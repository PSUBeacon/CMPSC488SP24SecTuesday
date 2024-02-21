package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"github.com/joho/godotenv"
	"go.bug.st/serial"
	"io"
	"log"
	"os"
	"strings"
)

func decryptAES(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

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
		err = godotenv.Load()
		AesKey := os.Getenv("AES_KEY") //This key is for testing, will be switched later
		message = []byte(strings.Trim(string(message), "\n"))
		// Decrypt the message.
		decryptedText, err := decryptAES([]byte(AesKey), message)
		if err != nil {
			fmt.Println("Error decrypting:", err)
			return
		}
		fmt.Printf("Decrypted text: %s\n", decryptedText)
	}
}

func main() {
	ConfigureController()
}
