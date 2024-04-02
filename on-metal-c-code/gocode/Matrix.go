package gocode

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	dinPinNumber = 9  // GPIO pin for DIN (MOSI)
	csPinNumber  = 4  // GPIO pin for CS
	clkPinNumber = 10 // GPIO pin for CLK
)

func TurnOff() {

	dinPin := rpio.Pin(dinPinNumber)
	csPin := rpio.Pin(csPinNumber)
	clkPin := rpio.Pin(clkPinNumber)
	initializeMatrix(dinPin, csPin, clkPin)
	dinPin.Output()
	csPin.Output()
	clkPin.Output()

	for row := 0; row < 8; row++ {
		sendData(csPin, dinPin, clkPin, byte(row+1), 0x00)
	}
}

// Initialize the LED matrix
func initializeMatrix(dinPin, csPin, clkPin rpio.Pin) {
	if err := rpio.Open(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Unable to open GPIO: %v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
	defer func() {
		err := rpio.Close()
		if err != nil {

		}
	}()
	// Set the scan-limit to show all 8 digits
	sendData(csPin, dinPin, clkPin, 0x0B, 0x07)

	// Use normal operation mode (not test mode)
	sendData(csPin, dinPin, clkPin, 0x0F, 0x00)

	// Set the intensity (brightness) of the display
	sendData(csPin, dinPin, clkPin, 0x0A, 0x0F)

	// Turn on the display
	sendData(csPin, dinPin, clkPin, 0x0C, 0x01)
}

// Set the intensity (brightness) of the LED matrix
func setIntensity(intensity byte) {
	dinPin := rpio.Pin(dinPinNumber)
	csPin := rpio.Pin(csPinNumber)
	clkPin := rpio.Pin(clkPinNumber)

	initializeMatrix(dinPin, csPin, clkPin)

	if intensity > 0x0F {
		intensity = 0x0F // Maximum intensity value is 0x0F
	}
	sendData(csPin, dinPin, clkPin, 0x0A, intensity)
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

// Drawing functions
func drawLightbulb(dinPin, csPin, clkPin rpio.Pin) {
	clearMatrix(csPin, dinPin, clkPin)
	lightbulbPattern := []byte{
		0b00111100,
		0b01111110,
		0b01111110,
		0b01111110,
		0b01111110,
		0b00111100,
		0b00011000,
		0b00011000,
	}
	for row, pattern := range lightbulbPattern {
		sendData(csPin, dinPin, clkPin, byte(row+1), pattern)
	}
}

func drawLock(dinPin, csPin, clkPin rpio.Pin) {
	clearMatrix(csPin, dinPin, clkPin)
	lockPattern := []byte{
		0b00111100,
		0b00100100,
		0b01100110,
		0b01111110,
		0b01111110,
		0b01111110,
		0b01111110,
		0b01111110,
	}
	for row, pattern := range lockPattern {
		sendData(csPin, dinPin, clkPin, byte(row+1), pattern)
	}
}

func drawH(dinPin, csPin, clkPin rpio.Pin) {
	clearMatrix(csPin, dinPin, clkPin)
	hPattern := []byte{
		0b10000001,
		0b10000001,
		0b10000001,
		0b11111111,
		0b10000001,
		0b10000001,
		0b10000001,
		0b10000001,
	}
	for row, pattern := range hPattern {
		sendData(csPin, dinPin, clkPin, byte(row+1), pattern)
	}
}

func drawA(dinPin, csPin, clkPin rpio.Pin) {
	clearMatrix(csPin, dinPin, clkPin)
	hPattern := []byte{
		0b00011000,
		0b00111100,
		0b01100110,
		0b01100110,
		0b01111110,
		0b01100110,
		0b01100110,
		0b01100110,
	}
	for row, pattern := range hPattern {
		sendData(csPin, dinPin, clkPin, byte(row+1), pattern)
	}
}

func MatrixStatus(dinPin, csPin, clkPin rpio.Pin, status bool) {
	if err := rpio.Open(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Unable to open GPIO: %v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
	defer func() {
		err := rpio.Close()
		if err != nil {

		}
	}()

	dinPin.Output()
	csPin.Output()
	clkPin.Output()

	// Set the scan-limit to show all 8 digits
	sendData(csPin, dinPin, clkPin, 0x0B, 0x07)

	// Use normal operation mode (not test mode)
	sendData(csPin, dinPin, clkPin, 0x0F, 0x00)

	// Set the intensity (brightness) of the display
	sendData(csPin, dinPin, clkPin, 0x0A, 0x0F)

	// Turn on the display
	sendData(csPin, dinPin, clkPin, 0x0C, 0x01)

	if status == false {
		clearMatrix(csPin, dinPin, clkPin)
	}
	if status == true {
		OnPattern := []byte{
			0b11111111,
			0b11111111,
			0b11111111,
			0b11111111,
			0b11111111,
			0b11111111,
			0b11111111,
			0b11111111,
		}
		for row, pattern := range OnPattern {
			sendData(csPin, dinPin, clkPin, byte(row+1), pattern)
		}
	}
}

// Clear the LED matrix
func clearMatrix(csPin, dinPin, clkPin rpio.Pin) {
	for row := 0; row < 8; row++ {
		sendData(csPin, dinPin, clkPin, byte(row+1), 0x00)
	}
}
