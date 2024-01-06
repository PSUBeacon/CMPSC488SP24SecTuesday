package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Simulated smart home devices
	lights := false
	thermostat := 70

	// Create a reader for user input
	reader := bufio.NewReader(os.Stdin)

	// Main menu loop
	for {
		fmt.Println("\nSmart Home Control System")
		fmt.Println("1. Toggle Lights")
		fmt.Println("2. Adjust Thermostat")
		fmt.Println("3. Exit")
		fmt.Print("Select an option: ")

		// Read user input
		choice, _ := reader.ReadString('\n')

		switch choice {
		case "1\r\n":
			// Toggle lights
			lights = !lights
			if lights {
				fmt.Println("Lights turned on.")
			} else {
				fmt.Println("Lights turned off.")
			}

		case "2\r\n":
			// Adjust thermostat
			fmt.Print("Enter new temperature: ")
			_, err := fmt.Scan(&thermostat)
			if err != nil {
				fmt.Println("Invalid input for thermostat.")
			} else {
				fmt.Printf("Thermostat set to %d°F\n", thermostat)
			}

		case "3\r\n":
			// Exit
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Please select a valid option.")
		}
	}
}
