package main

import (
	"CMPSC488SP24SecTuesday/blockchain"
	"CMPSC488SP24SecTuesday/crypto"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func decryptAES(key, ciphertext []byte) ([]byte, error) {
	//ciphertext = []byte(strings.Trim(string(ciphertext), "*"))
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

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

func BlockReceiver() {
	// Open the XBee module for communication
	var chain blockchain.Blockchain
	var block blockchain.Block

	fmt.Println("Waiting for incoming messages...")
	// Use ReadBytes or ReadString to dynamically handle incoming data

	//loads file and pulls the key from there
	err := godotenv.Load()
	AesKey := os.Getenv("AES_KEY")

	//Decrypt the message.
	decryptedText, err := decryptAES([]byte(AesKey), message)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}
	//fmt.Printf("Decrypted text: %s\n", decryptedText)

	//Verify the HMAC
	isValid, receivedMessage := crypto.VerifyHMAC(decryptedText)
	//if valid will check the blockchain
	if isValid {
		fmt.Println("Message integrity verified successfully.")
		jsonChainData, err := os.ReadFile("chain.json")
		chainlen := len(jsonChainData)
		if err != nil {
			panic(err)
		}

		//Checks if there is an existing chain or if this is the start of the chain
		if chainlen == 0 {
			chainTojson := json.Unmarshal(receivedMessage, &chain)
			if chainTojson != nil {
				// if error is not nil
				fmt.Println(chainTojson)
			}
			// Marshal the chain struct to JSON
			jsonChainData, err = json.MarshalIndent(chain, "", "    ")
			if err != nil {
				panic(err)
			}
			// Write the JSON data to a file
			err = os.WriteFile("chain.json", jsonChainData, 0644)
			if err != nil {
				panic(err)
			}
			continue

		}

		//Checks if the incoming block is not the first block in a chain
		if chainlen > 0 {
			blockTojson := json.Unmarshal(receivedMessage, &block)
			if blockTojson != nil {
				fmt.Println(blockTojson)
			}
			verify := verifyBlockchain(block)
			if verify == true {

				err := json.Unmarshal(jsonChainData, &chain)
				if err != nil {
					return
				}

				chain.Chain = append(chain.Chain, block)
				// Marshal the chain struct to JSON
				jsonChainData, err = json.MarshalIndent(chain, "", "    ")
				//fmt.Println("This is the chain after its been appended: ", string(jsonChainData))
				if err != nil {
					panic(err)
				}
				// Write the JSON data to a file
				err = os.WriteFile("chain.json", jsonChainData, 0644)
				if err != nil {
					panic(err)
				}
			}
			if verify == false {
				fmt.Println("Invalid Block")
			}
		}
	} else {
		fmt.Println("Message integrity verification failed.")
		continue

	}
}

func verifyBlockchain(currentblock blockchain.Block) bool {
	// Read the JSON file
	//fmt.Println("verfiy got this: ", currentblock)
	jsonChainData, err := os.ReadFile("chain.json")
	if err != nil {
		panic(err)
	}

	var readBlockchain blockchain.Blockchain
	err = json.Unmarshal(jsonChainData, &readBlockchain)
	if err != nil {
		panic(err)
	}

	if readBlockchain.Chain[len(readBlockchain.Chain)-1].Hash == currentblock.PrevHash {
		// Verify the rest of the hashes
		for i := 1; i < len(readBlockchain.Chain); i++ {
			currBlock := readBlockchain.Chain[i]
			prevBlock := readBlockchain.Chain[i-1]

			if currBlock.PrevHash != prevBlock.Hash { //invalid hash
				return false
			}
		}
	}
	fmt.Println("block and chain is valid")
	return true
}

//
//func main() {
//	BlockReceiver()
//}
