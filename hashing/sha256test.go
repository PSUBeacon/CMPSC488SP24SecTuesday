package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	input := "Hello, World!" // The string you want to hash

	// Record the start time
	startTime := time.Now()

	// Create a new SHA-256 hasher
	hasher := sha256.New()

	// Write the input string to the hasher
	hasher.Write([]byte(input))

	// Calculate the SHA-256 hash
	hash := hasher.Sum(nil)

	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hash)

	// Record the end time
	endTime := time.Now()

	// Calculate the duration
	duration := endTime.Sub(startTime)

	fmt.Printf("Input: %s\n", input)
	fmt.Printf("SHA-256 Hash: %s\n", hashString)
	fmt.Printf("Hashing took %s\n", duration)
}
