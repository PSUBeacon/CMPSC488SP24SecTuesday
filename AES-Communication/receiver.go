package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
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

func decryptAES(key, ciphertext []byte) ([]byte, error) {
	//ciphertext = []byte(strings.Trim(string(ciphertext), "*"))
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
	defer func(port serial.Port) {
		err := port.Close()
		if err != nil {

		}
	}(port) // Ensure the port is closed when the function returns

	// Wrap the port in a bufio.Reader
	const bufferSize = 4096 // Adjust this value as needed
	reader := bufio.NewReaderSize(port, bufferSize)

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
		//Decrypt the message.
		decryptedText, err := decryptAES([]byte(AesKey), []byte(message))
		//decryptedText := message

		if err != nil {
			fmt.Println("Error decrypting:", err)
			return
		}
		fmt.Printf("Decrypted text: %s\n", decryptedText)

		tojson := json.Unmarshal(decryptedText, &block)
		if tojson != nil {

			// if error is not nil
			// print error
			fmt.Println(tojson)
		}
		//fmt.Println(tojson)
		for i := range block {
			fmt.Println(string(rune(block[i].Index)) + " - " + block[i].Timestamp + " - " + block[i].Data + " - " + block[i].PrevHash + " - " + block[i].Hash)
		}
	}
}

func main() {
	ConfigureController()
}