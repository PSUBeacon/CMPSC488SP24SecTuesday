package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const sensorPin = 11

func main() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println("Unable to open GPIO:", err)
		os.Exit(1)
	}
	defer rpio.Close()

	// Set pin to input mode
	pin := rpio.Pin(sensorPin)
	pin.Input()

	fmt.Println("Monitoring for motion...")

	for {
		// Read pin state
		if pin.Read() == rpio.High {
			fmt.Println("Motion detected!")
		} else {
			fmt.Println("No motion.")
		}

		// Sleep for a while to reduce CPU usage
		time.Sleep(1 * time.Second)
	}
}
