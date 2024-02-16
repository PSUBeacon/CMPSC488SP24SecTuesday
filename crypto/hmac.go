package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// generate an HMAC for the given message and key [sha256]
func generateHMAC(message, key []byte) string {
	hmacHash := hmac.New(sha256.New, key)
	hmacHash.Write(message)
	hashedMessage := hmacHash.Sum(nil)
	return hex.EncodeToString(hashedMessage)
}

// add an HMAC to the message and return the combined payload
func addHMAC(message, key []byte) []byte {
	hmacValue := generateHMAC(message, key)
	return append(message, []byte(hmacValue)...)
}

// verify the integrity of the message and returns true if it's valid
func verifyHMAC(payload, key []byte) bool {
	if len(payload) < 64 {
		// The payload must have at least 64 characters for the HMAC
		return false
	}

	message := payload[:len(payload)-64]
	receivedHMAC := payload[len(payload)-64:]

	// recalculate HMAC to compare them later
	calculatedHMAC := generateHMAC(message, key)

	// compare recalculated and received HMACs
	return hmac.Equal([]byte(calculatedHMAC), receivedHMAC)
}

/*
func main() {
	// Example message and secret key
	message := []byte("Hello, HMAC!")
	secretKey := []byte("mySecretKey")

	// Add HMAC to the message
	payload := addHMAC(message, secretKey)

	// Transmit the payload (simulate transport)

	// On the receiving side, verify the integrity
	isValid := verifyHMAC(payload, secretKey)

	// Print the verification result
	if isValid {
		fmt.Println("Message integrity verified successfully.")
	} else {
		fmt.Println("Message integrity verification failed.")
	}
}
*/
