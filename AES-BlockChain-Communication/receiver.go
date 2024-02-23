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
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	Data      string `json:"data"`
	PrevHash  string `json:"prevHash"`
	Hash      string `json:"hash"`
}

type Blockchain struct {
	Chain []Block `json:"chain"`
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

	//fmt.Printf("Ciphertext length (including IV): %d\n", len(ciphertext))

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	//fmt.Printf("Ciphertext length (excluding IV): %d\n", len(ciphertext))

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

func blockReceiver() {
	// Open the XBee module for communication
	var chain Blockchain
	var block Block
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

		if len(chain.Chain) == 0 {
			chainTojson := json.Unmarshal(decryptedText, &chain)
			if chainTojson != nil {
				// if error is not nil
				fmt.Println(chainTojson)
				verifyBlock()
			}
		}

		if len(chain.Chain) > 0 {
			blockTojson := json.Unmarshal(decryptedText, &block)
			if blockTojson != nil {
				fmt.Println(blockTojson)
			}
			chain.Chain = append(chain.Chain, block)
		}
		// Marshal the chain struct to JSON
		jsonChainData, err := json.MarshalIndent(chain, "", "    ")
		if err != nil {
			panic(err)
		}

		// Write the JSON data to a file
		err = os.WriteFile("chain.json", jsonChainData, 0644)
		if err != nil {
			panic(err)
		}
		// Read the JSON file
		jsonChainData, err = os.ReadFile("chain.json")
		if err != nil {
			panic(err)
		}

		var readBlockchain Blockchain
		err = json.Unmarshal(jsonChainData, &readBlockchain)
		if err != nil {
			panic(err)
		}
		fmt.Println("This is the current block: ", readBlockchain.Chain)
		// Print the data
		fmt.Printf("%+v\n", readBlockchain)

		//fmt.Println(chain)
		//for i := range chain.Chain {
		//	fmt.Println(string(rune(chain.Chain[i].Index)) + " - " + chain.Chain[i].Timestamp + " - " + chain.Chain[i].Data + " - " + chain.Chain[i].PrevHash + " - " + chain.Chain[i].Hash)
		//}
	}
}

func verifyBlock() {
	// Read the JSON file
	jsonChainData, err := os.ReadFile("chain.json")
	if err != nil {
		panic(err)
	}

	var readBlockchain Blockchain
	err = json.Unmarshal(jsonChainData, &readBlockchain)
	if err != nil {
		panic(err)
	}

	// Verify the hashes
	for i := 1; i < len(readBlockchain.Chain); i++ {
		currentBlock := readBlockchain.Chain[i]
		previousBlock := readBlockchain.Chain[i-1]

		if currentBlock.PrevHash != previousBlock.Hash {
			fmt.Printf("Block %d has an invalid previous hash\n", currentBlock.Index)
			return
		}
	}

	fmt.Println("All blocks have valid previous hashes")
}

func main() {
	blockReceiver()
}
