package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

// function to get  HMAC secret key from an environment variable
func getHMACSecretKey() []byte {
	// get the secret key from an environment variable
	key := os.Getenv("HMAC_SECRET_KEY")

	// generate a random key if the key is not available in an environment variable
	if key == "" {
		// if environment variable doesn't work, use this key - fallback value
		key = "405196174e9ac81268eadaba6bacea0fcbdd6446bccea9642c6bc3b53a47b8b3"
		err := os.Setenv("HMAC_SECRET_KEY", key)
		if err != nil {
			fmt.Println("Error setting environment variable:", err)
		}
	}

	return []byte(key)
}

// generate an HMAC for the given message and key [sha256]
func generateHMAC(message []byte) string {
	key := getHMACSecretKey()
	hmacHash := hmac.New(sha256.New, key)
	hmacHash.Write(message)
	hashedMessage := hmacHash.Sum(nil)
	return hex.EncodeToString(hashedMessage)
}

// add an HMAC to the message and return the combined payload
func AddHMAC(message []byte) []byte {
	hmacValue := generateHMAC(message)
	return append(message, []byte(hmacValue)...)
}

// verify the integrity of the message and returns true if it's valid
func VerifyHMAC(payload []byte) (bool, []byte) {
	if len(payload) < 64 {
		// The payload must have at least 64 characters for the HMAC
		return false, nil
	}

	message := payload[:len(payload)-64]
	receivedHMAC := payload[len(payload)-64:]

	// recalculate HMAC to compare them later
	calculatedHMAC := generateHMAC(message)

	// compare recalculated and received HMACs
	// return comparison boolean value, as well as the message itself
	return hmac.Equal([]byte(calculatedHMAC), receivedHMAC), message
}

//func main() {
//	// load .env file which is in gitignore
//	err := godotenv.Load(".env")
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//
//	// Example message and secret key
//	message := []byte("HMAC verification message")
//
//	// Add HMAC to the message
//	payload := AddHMAC(message)
//
//	// On the receiving side, verify the integrity
//	isValid, receivedMessage := VerifyHMAC(payload)
//
//	// Print the verification result
//	if isValid {
//		fmt.Println("Message integrity verified successfully.")
//		fmt.Println(string(receivedMessage))
//	} else {
//		fmt.Println("Message integrity verification failed.")
//	}
//}
