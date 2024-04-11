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

// Set the intensity (brightness) of the LED matrix
func SetIntensity(dinPin, csPin, clkPin rpio.Pin, intensity int) {

	if err := rpio.Open(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open GPIO: %v\n", err)
		os.Exit(1)
	}
	defer rpio.Close()

	dinPin.Output()
	csPin.Output()
	clkPin.Output()

	initializeMatrix(dinPin, csPin, clkPin)

	intensityByte := byte(intensity)
	if intensity > 0x0F {
		intensity = 0x0F // Maximum intensity value is 0x0F
	}
	sendData(csPin, dinPin, clkPin, 0x0A, intensityByte)
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
func DrawLightbulb(dinPin, csPin, clkPin rpio.Pin, brightness int) {

	ClearMatrix(csPin, dinPin, clkPin)
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
		sendData(csPin, dinPin, clkPin, 0x0A, byte(brightness))
	}
	time.Sleep(2 * time.Second)
	ClearMatrix(csPin, dinPin, clkPin)
}

func drawLock(dinPin, csPin, clkPin rpio.Pin) {
	if err := rpio.Open(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open GPIO: %v\n", err)
		os.Exit(1)
	}
	defer rpio.Close()

	dinPin.Output()
	csPin.Output()
	clkPin.Output()
	initializeMatrix(dinPin, csPin, clkPin)

	ClearMatrix(csPin, dinPin, clkPin)
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
	time.Sleep(2 * time.Second)
	ClearMatrix(csPin, dinPin, clkPin)
}

func drawH(dinPin, csPin, clkPin rpio.Pin) {
	if err := rpio.Open(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open GPIO: %v\n", err)
		os.Exit(1)
	}
	defer rpio.Close()

	dinPin.Output()
	csPin.Output()
	clkPin.Output()
	initializeMatrix(dinPin, csPin, clkPin)

	ClearMatrix(csPin, dinPin, clkPin)
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
	if err := rpio.Open(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open GPIO: %v\n", err)
		os.Exit(1)
	}
	defer rpio.Close()

	dinPin.Output()
	csPin.Output()
	clkPin.Output()
	initializeMatrix(dinPin, csPin, clkPin)

	ClearMatrix(csPin, dinPin, clkPin)
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

func MatrixStatus(dinPin, csPin, clkPin rpio.Pin, status bool, brightness int) {
	if err := rpio.Open(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open GPIO: %v\n", err)
		os.Exit(1)
	}
	defer rpio.Close()

	dinPin.Output()
	csPin.Output()
	clkPin.Output()
	initializeMatrix(dinPin, csPin, clkPin)

	DrawLightbulb(dinPin, csPin, clkPin, brightness)

	if status == false {
		ClearMatrix(csPin, dinPin, clkPin)
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

func TurnOffMatrix(dinPin, csPin, clkPin rpio.Pin) {
	if err := rpio.Open(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open GPIO: %v\n", err)
		os.Exit(1)
	}
	defer rpio.Close()

	dinPin.Output()
	csPin.Output()
	clkPin.Output()
	initializeMatrix(dinPin, csPin, clkPin)
	OnPattern := []byte{
		0b00000000,
		0b00000000,
		0b00000000,
		0b00000000,
		0b00000000,
		0b00000000,
		0b00000000,
		0b00000000,
	}
	for row, pattern := range OnPattern {
		sendData(csPin, dinPin, clkPin, byte(row+1), pattern)
	}
}

// Clear the LED matrix
func ClearMatrix(csPin, dinPin, clkPin rpio.Pin) {

	for row := 0; row < 8; row++ {
		sendData(csPin, dinPin, clkPin, byte(row+1), 0x00)
	}
}

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

	// Initialize and turn on the matrix
	initializeMatrix(dinPin, csPin, clkPin)

	// Set the intensity of the LED matrix
	SetIntensity(dinPin, csPin, clkPin, 15)

	// Display a pattern on the matrix
	// You can replace this with any pattern or function you want to display
	DrawLightbulb(dinPin, csPin, clkPin, 15)

	// Keep the matrix on for some time
	time.Sleep(5 * time.Second)

	// Turn off the matrix
	TurnOffMatrix(dinPin, csPin, clkPin)
}
