package main

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code"
	"fmt"
	"time"
)

func main() {
	matrix := cCode.NewMaxMatrix(9, 4, 10, 1) // Use the appropriate GPIO pins for data, load, and clock
	matrix.Init(1)                            // Initialize 1 matrix
	matrix.SetIntensity(8)                    // Set intensity to a medium level

	fmt.Println("Clearing matrix...")
	matrix.Clear() // Clear the matrix

	// Keep the program running for a while to observe the effects
	time.Sleep(10 * time.Second)
}
