package main

import (
	"CMPSC488SP24SecTuesday/blockchain"
	"CMPSC488SP24SecTuesday/crypto"
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

	//fmt.Printf("Ciphertext length after encryption: %d\n", len(ciphertext))

	return ciphertext, nil
}

func BroadCastMessage(messageToSend []byte) {
	// The key should be 16, 24, or 32 bytes long for AES-128, AES-192, or AES-256, respectively.
	err := godotenv.Load()
	AesKey := os.Getenv("AES_KEY")

	jsonChainData, err := os.ReadFile("chain.json")
	chainlen := len(jsonChainData)
	if err != nil {
		panic(err)
	}
	fmt.Println("This is the length of whats in the file: ", len(jsonChainData))

	//Checks if there is an existing chain or if this is the start of the chain
	if chainlen == 0 {
		// Create a new blockchain
		chain := blockchain.NewBlockchain()
		theChain, err := json.MarshalIndent(chain, "", "  ")
		if err != nil {
			panic(err)
		}
		err = os.WriteFile("chain.json", theChain, 0644)
		if err != nil {
			panic(err)
		}
		// Create the first block with messageToSend
		firstBlock := blockchain.CreateBlock(string(messageToSend), chain.Chain[0].Hash, chain.Chain[0].Index+1)

		// Add the first block to the chain
		chain.Chain = append(chain.Chain, firstBlock)

		// Marshal the entire blockchain with the new block added
		updatedChain, err := json.MarshalIndent(chain, "", "  ")
		if err != nil {
			panic(err)
		}

		// Write the updated chain data to a file
		err = os.WriteFile("chain.json", updatedChain, 0644)
		if err != nil {
			panic(err)
		}

		//creating the HMAC with unencrypted data
		hmacChain := crypto.AddHMAC(updatedChain)
		fmt.Println("This is the updated chain with hmac", hmacChain)
		// Encrypt and send the new block
		encryptedBlock, err := encryptAES([]byte(AesKey), hmacChain)
		if err != nil {
			log.Fatal("Error encrypting block:", err)
		}

		send(encryptedBlock)
		return
	}

	if chainlen > 0 {
		var jsonChain blockchain.Blockchain
		err := json.Unmarshal(jsonChainData, &jsonChain)
		if err != nil {
			log.Fatal("Error unmarshalling chain data:", err)
		}

		// Pass the hash of the last block in the chain as the prevHash for the new block
		lastBlock := jsonChain.Chain[len(jsonChain.Chain)-1]
		newBlock := blockchain.CreateBlock(string(messageToSend), lastBlock.Hash, lastBlock.Index+1)

		// Add the new block to the chain
		jsonChain.Chain = append(jsonChain.Chain, newBlock)

		// Marshal the entire blockchain with the new block added
		jsonBlock, err := json.MarshalIndent(newBlock, "", "  ")
		if err != nil {
			panic(err)
		}

		updatedChain, err := json.MarshalIndent(jsonChain, "", "  ")
		if err != nil {
			panic(err)
		}

		err = os.WriteFile("chain.json", updatedChain, 0644)
		if err != nil {
			panic(err)
		}
		//fmt.Println("This is the block before encryption: ", jsonBlock)

		//creating the HMAC with unencrypted data
		hmacBlock := crypto.AddHMAC(jsonBlock)

		fmt.Println("This is the updated block with hmac", hmacBlock)
		encryptedBlock, err := encryptAES([]byte(AesKey), hmacBlock)
		if err != nil {
			log.Fatal("Error encrypting block:", err)
		}

		send(encryptedBlock)
		return
	}
}
func send(message []byte) {
	// Open the XBee module for communication
	mode := &serial.Mode{
		BaudRate: 9600,
	}
	//fmt.Println("This is in the send function", message)
	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		log.Fatal("Error opening XBee module:", err)
	}
	defer func(port serial.Port) {
		err := port.Close()
		if err != nil {

		}
	}(port)

	delimiter := []byte{0xE2, 0x99, 0xB4}
	message = append(message, delimiter...)

	// Send a message to the server
	fmt.Println(len(message))
	_, err = port.Write(message)
	fmt.Printf("Sent \n")
	if err != nil {
		log.Println("Error sending message:", err)
	}
	return
}
func main() {

	BroadCastMessage([]byte("testing the stuff"))
	BroadCastMessage([]byte("this is another one"))

}
