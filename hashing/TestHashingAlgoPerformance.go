package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	input := "Hello, World!" // The string you want to hash
	numIterations := 100000  // Number of iterations for performance testing
	warmUpIterations := 1000 // Number of warm-up iterations

	// Warm-up runs
	fmt.Println("Performing warm-up runs...")
	warmUp(warmUpIterations)

	// Perform the performance test for MD5 hashing
	totalTimeMD5, averageTimeMD5 := runPerformanceTest(hashMD5, input, numIterations)
	printResults("MD5", input, numIterations, totalTimeMD5, averageTimeMD5)

	// Perform the performance test for SHA1 hashing
	totalTimeSHA1, averageTimeSHA1 := runPerformanceTest(hashSHA1, input, numIterations)
	printResults("SHA1", input, numIterations, totalTimeSHA1, averageTimeSHA1)

	// Perform the performance test for SHA256 hashing
	totalTimeSHA256, averageTimeSHA256 := runPerformanceTest(hashSHA256, input, numIterations)
	printResults("SHA256", input, numIterations, totalTimeSHA256, averageTimeSHA256)
}

func warmUp(iterations int) {
	// To avoid skewed results.
	for i := 0; i < iterations; i++ {
		hashMD5("Warm-up")
		hashSHA1("Warm-up")
		hashSHA256("Warm-up")
	}
}

func runPerformanceTest(hashFunc func(string) string, input string, numIterations int) (float64, float64) {
	startTime := time.Now()
	for i := 0; i < numIterations; i++ {
		hashFunc(input)
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	averageTime := duration.Seconds() / float64(numIterations)
	totalTime := duration.Seconds()
	return totalTime, averageTime
}

func printResults(algo, input string, numIterations int, totalTime, averageTime float64) {
	fmt.Println("=====================", algo, "Hash==============================")
	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Number of Iterations: %d\n", numIterations)
	fmt.Printf("Total time taken: %.2f seconds\n", totalTime)
	fmt.Printf("Average time per iteration: %.10f seconds\n", averageTime)
	fmt.Println("===========================================================")
}

func hashMD5(input string) string {
	// Create an MD5 hasher
	hasher := md5.New()

	// Write the input string to the hasher
	hasher.Write([]byte(input))

	// Calculate the MD5 hash
	hash := hasher.Sum(nil)

	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hash)
	return hashString
}

func hashSHA1(input string) string {
	// Create an SHA1 hasher
	hasherSHA1 := sha1.New()

	// Write the input string to the hasher
	hasherSHA1.Write([]byte(input))

	// Calculate the SHA1 hash
	hash := hasherSHA1.Sum(nil)

	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hash)
	return hashString
}

func hashSHA256(input string) string {
	// Create an SHA1 hasher
	hasherSHA256 := sha256.New()

	// Write the input string to the hasher
	hasherSHA256.Write([]byte(input))

	// Calculate the SHA1 hash
	hash := hasherSHA256.Sum(nil)

	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hash)
	return hashString
}
