package main

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

func main() {
	dinPinNumber := 9  // GPIO pin for DIN (MOSI)
	csPinNumber := 4   // GPIO pin for CS
	clkPinNumber := 10 // GPIO pin for CLK

	// Set up the pins
	dinPin := rpio.Pin(dinPinNumber)
	csPin := rpio.Pin(csPinNumber)
	clkPin := rpio.Pin(clkPinNumber)

	// Display a pattern on the matrix

	// You can replace this with any pattern or function you want to display
	gocode.DrawLightbulb(dinPin, csPin, clkPin, 15)
	// Keep the matrix on for some time
	time.Sleep(5 * time.Second)

	// Turn off the matrix
	gocode.TurnOffMatrix(dinPin, csPin, clkPin)
}
