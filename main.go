package main

import (
	"CMPSC488SP24SecTuesday/appliances"
	"CMPSC488SP24SecTuesday/blockchain"
	"CMPSC488SP24SecTuesday/crypto"
	"CMPSC488SP24SecTuesday/dal"
	"CMPSC488SP24SecTuesday/hvac"
	"CMPSC488SP24SecTuesday/lighting"
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"CMPSC488SP24SecTuesday/security"
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.bug.st/serial"
	"log"
	"os"
	"strconv"
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

	//ticker := time.NewTicker(60 * time.Second)

	// Use ReadBytes or ReadString to dynamically handle incoming data
	for {
		// Read and parse the data manually
		var message []byte
		for {
			b, err := reader.ReadByte()
			if err != nil {
				log.Fatal("Error reading byte:", err)
			}
			//if ticker == nil {
			//go messaging.BroadCastMessage([]byte("pi # connected"))
			//ticker.Reset(60 * time.Second)

			//}

			// Check for the UTF-8 encoding of '♄' the hex value is (E2 99 B4)
			if len(message) >= 2 && message[len(message)-2] == 0xE2 && message[len(message)-1] == 0x99 && b == 0xB4 {
				//fmt.Println(message)
				message = message[:len(message)-2] // Remove the delimiter from the message
				break
			}
			message = append(message, b)
		}

		//loads file and pulls the key from there
		err = godotenv.Load()
		AesKey := os.Getenv("AES_KEY")

		//Decrypt the message.
		decryptedText, err := decryptAES([]byte(AesKey), message)
		if err != nil {
			fmt.Println("Error decrypting:", err)
			//return
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
				//fmt.Println("Got to functionality")
				go handleFunctionality()
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
					//fmt.Println("Got to functionality")
					go handleFunctionality()
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
		//for i := 1; i < len(readBlockchain.Chain); i++ {
		//currBlock := readBlockchain.Chain[i]
		//prevBlock := readBlockchain.Chain[i-1]

		//if currBlock.PrevHash != prevBlock.Hash { //invalid hash

		//checks just the previous block
		fmt.Println("block and chain is valid")
		return true
	}
	return false
}

func handleFunctionality() {

	jsonChainData, err := os.ReadFile("chain.json")
	if err != nil {
		panic(err)
	}

	var readBlockchain blockchain.Blockchain
	err = json.Unmarshal(jsonChainData, &readBlockchain)
	if err != nil {
		panic(err)
	}
	chainlen := len(readBlockchain.Chain)
	latestBlockData := readBlockchain.Chain[chainlen-1].Data
	//fmt.Println(latestBlockData)

	var messageData dal.MessagingStruct
	var UUIDsData dal.UUIDsConfig

	err = json.Unmarshal([]byte(latestBlockData), &messageData)
	fmt.Println(messageData)
	jsonconfigData, err := os.ReadFile("config.json")

	err = json.Unmarshal(jsonconfigData, &UUIDsData)
	if err != nil {
		panic(err)
	}

	messageChange, _ := strconv.Atoi(messageData.Change)

	//fmt.Println("This is the uuids", UUIDsData)
	if messageData.Name == "Lighting" {
		//fmt.Println("Got past the name")
		for _, Pi := range UUIDsData.Lighting {
			if Pi.UUID == messageData.UUID {
				//fmt.Println("got past the loop")
				if messageData.Function == "Status" {
					if messageData.Change == "false" {
						//fmt.Println("inside the status false")
						lighting.UpdateStatus(false)
					}
					if messageData.Change == "true" {
						//fmt.Println("inside the status false")
						lighting.UpdateStatus(true)
					}
				}
				if messageData.Function == "Brightness" {
					lighting.SetBrightness(messageChange)
				}
			}
		}
	}
	if messageData.Name == "HVAC" {
		for _, group := range [][]dal.Pi{UUIDsData.Hvac} {
			for _, Pi := range group {
				if Pi.UUID == messageData.UUID {
					if messageData.Function == "Status" {
						if messageData.Change == "false" {
							hvac.UpdateStatus(false, messageData.UUID)
							//hvac.DisplayLCDHVAC("", 0, "OFF")
						}
						if messageData.Change == "true" {
							hvac.UpdateStatus(true, messageData.UUID)
							//hvac.DisplayLCDHVAC("", 0, "ON")
						}
					}
					if messageData.Function == "FanSpeed" {
						hvac.UpdateFanSpeed(messageChange, messageData.UUID)
						//hvac.DisplayLCDHVAC("", 0, "ON")

					}
					if messageData.Function == "Temperature" {
						hvac.UpdateTemperature(messageChange, messageData.UUID)
						//hvac.DisplayLCDHVAC("", messageChange, "")
					}
					if messageData.Function == "Mode" {
						hvac.UpdateMode(messageData.Change, messageData.UUID)
						//hvac.DisplayLCDHVAC(messageData.Change, 0, "")
					}
				}
			}
		}
	}

	if messageData.Name == "Security" {
		for _, group := range [][]dal.Pi{UUIDsData.Security} {
			for _, Pi := range group {
				if Pi.UUID == messageData.UUID {
					if messageData.Function == "Status" {
						if messageData.Change == "false" {
							security.UpdateAlarmStatus(false)
							security.DisplayLCDSecurity("Disarmed", "OFF")
						}
						if messageData.Change == "true" {
							security.UpdateAlarmStatus(true)
							security.DisplayLCDSecurity("Armed", "ON")
						}
					}
					if messageData.Function == "LockStatus" {
						if messageData.Change == "false" {
							security.LockOrUnlock(false)
						}
						if messageData.Change == "true" {
							security.LockOrUnlock(true)
						}
					}
				}
			}
		}
	}
	if messageData.Name == "Appliances" {
		for _, group := range [][]dal.Pi{UUIDsData.Appliances} {
			for _, Pi := range group {
				if Pi.UUID == messageData.UUID {
					if messageData.Function == "Status" {
						if messageData.Change == "false" {
							appliances.UpdateStatus(messageData.AppType, false)
							fmt.Println("got here")
						}
						if messageData.Change == "true" {
							appliances.UpdateStatus(messageData.AppType, true)
							fmt.Println("got here")
						}
					}
					if messageData.Function == "Temperature" {
						appliances.UpdateTemperature(messageChange)
					}
					if messageData.Function == "TimerStopTime" {
						appliances.UpdateTimeStopTime(messageChange)
					}
					if messageData.Function == "Power" {
						appliances.UpdatePower(messageChange)
					}
					if messageData.Function == "EnergySaveMode" {
						if messageData.Change == "false" {
							appliances.UpdateStatus(messageData.AppType, false)
						}
						if messageData.Change == "true" {
							appliances.UpdateStatus(messageData.AppType, true)
						}
					}
					if messageData.Function == "WashTime" {
						appliances.UpdateWashTime(messageChange)
					}
				}
			}
		}
	}
	if messageData.Name == "Energy" {
		for _, group := range [][]dal.Pi{UUIDsData.Energy} {
			for _, Pi := range group {
				if Pi.UUID == messageData.UUID {
					if messageData.Function == "Status" {
						//energy.,UpdateAlarmStatus(messageData.StatusChange)
					}
				}
			}
		}

	}
	return
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	assignedPiNum := os.Getenv("PI_NUM")
	piNum, _ := strconv.Atoi(assignedPiNum)
	if piNum == 13 {
		go gocode.InitKeypad()
	}
	if piNum == 16 {
		//hvac.DisplayLCDHVAC("", 0, "")
		//go hvac.SendTempToFE()
	}
	if piNum == 22 {
		security.DisplayLCDSecurity("", "")
		go gocode.InitKeypad()
	}
	BlockReceiver()
}
