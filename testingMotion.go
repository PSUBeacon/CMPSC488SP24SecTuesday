package main

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio/v4"
)

// PIR struct represents a PIR motion sensor
type PIR struct {
	signalPin rpio.Pin
}

// NewPIR creates a new PIR instance
func NewPIR(pinNumber int) *PIR {
	return &PIR{
		signalPin: rpio.Pin(pinNumber),
	}
}

// Read returns the current state of the PIR sensor (true for motion detected, false otherwise)
func (p *PIR) Read() bool {
	return p.signalPin.Read() == rpio.High
}

func main() {
	// Open and map memory to access GPIO, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println("Unable to open GPIO:", err)
		os.Exit(1)
	}
	defer rpio.Close()

	// Create a new PIR sensor instance connected to GPIO pin 17
	pirSensor := NewPIR(17)

	// Set the sensor pin to input mode
	pirSensor.signalPin.Input()

	fmt.Println("Monitoring for motion...")

	for {
		// Read the sensor state
		if pirSensor.Read() {
			fmt.Println("Motion detected!")
		} else {
			fmt.Println("No motion.")
		}
	}
}