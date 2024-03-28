package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	dinPinNumber = 9  // GPIO pin for DIN (MOSI)
	csPinNumber  = 4  // GPIO pin for CS
	clkPinNumber = 10 // GPIO pin for CLK
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open GPIO: %v\n", err)
		os.Exit(1)
	}
	defer rpio.Close()

	// Set up the pins
	dinPin := rpio.Pin(dinPinNumber)
	csPin := rpio.Pin(csPinNumber)
	clkPin := rpio.Pin(clkPinNumber)

	dinPin.Output()
	csPin.Output()
	clkPin.Output()

	// Initialize the LED matrix
	initializeMatrix(dinPin, csPin, clkPin)

	// Turn on all LEDs
	for row := 0; row < 8; row++ {
		sendData(csPin, dinPin, clkPin, byte(row+1), 0xFF)
	}
	time.Sleep(5 * time.Second)

	// Turn off all LEDs
	for row := 0; row < 8; row++ {
		sendData(csPin, dinPin, clkPin, byte(row+1), 0x00)
	}
}

// Initialize the LED matrix
func initializeMatrix(dinPin, csPin, clkPin rpio.Pin) {
	// Set the scan-limit to show all 8 digits
	sendData(csPin, dinPin, clkPin, 0x0B, 0x07)

	// Use normal operation mode (not test mode)
	sendData(csPin, dinPin, clkPin, 0x0F, 0x00)

	// Set the intensity (brightness) of the display
	sendData(csPin, dinPin, clkPin, 0x0A, 0x0F)

	// Turn on the display
	sendData(csPin, dinPin, clkPin, 0x0C, 0x01)
}

// Send data to the LED matrix
func sendData(csPin, dinPin, clkPin rpio.Pin, address, data byte) {
	csPin.Low()
	sendByte(dinPin, clkPin, address)
	sendByte(dinPin, clkPin, data)
	csPin.High()
}

// Send a single byte of data
func sendByte(dinPin, clkPin rpio.Pin, data byte) {
	for i := 7; i >= 0; i-- {
		clkPin.Low()
		if (data & (1 << i)) != 0 {
			dinPin.High()
		} else {
			dinPin.Low()
		}
		clkPin.High()
	}
}
