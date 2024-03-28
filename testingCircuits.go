package main

import (
	"log"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/spi"
	"time"

	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

// MAX7219 registers
const (
	NoOp        byte = 0x00
	Digit0      byte = 0x01
	Digit1      byte = 0x02
	Digit2      byte = 0x03
	Digit3      byte = 0x04
	Digit4      byte = 0x05
	Digit5      byte = 0x06
	Digit6      byte = 0x07
	Digit7      byte = 0x08
	DecodeMode  byte = 0x09
	Intensity   byte = 0x0A
	ScanLimit   byte = 0x0B
	ShutDown    byte = 0x0C
	DisplayTest byte = 0x0F
)

func main() {
	// Initialize periph.io library
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Open the first available SPI port
	port, err := spireg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	// Create a connection to the MAX7219
	conn, err := port.Connect(10*physic.MegaHertz, spi.Mode0, 8)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the MAX7219
	initMAX7219(conn)

	// Display a pattern on the matrix
	displayPattern(conn)

	// Keep the pattern displayed for 5 seconds
	time.Sleep(5 * time.Second)

	// Clear the display
	clearDisplay(conn)
}

// initMAX7219 initializes the MAX7219 LED matrix driver.
func initMAX7219(conn spi.Conn) {
	sendCommand(conn, ShutDown, 0x01)    // Exit shutdown mode
	sendCommand(conn, DisplayTest, 0x00) // Disable display test
	sendCommand(conn, DecodeMode, 0x00)  // No decode for digits
	sendCommand(conn, ScanLimit, 0x07)   // Display all digits
	sendCommand(conn, Intensity, 0x08)   // Set moderate intensity
	sendCommand(conn, ShutDown, 0x01)    // Exit shutdown mode
	clearDisplay(conn)                   // Clear display register
}

// sendCommand sends a command to the MAX7219.
func sendCommand(conn spi.Conn, register, data byte) {
	if err := conn.Tx([]byte{register, data}, nil); err != nil {
		log.Fatal(err)
	}
}

// displayPattern displays a pattern on the LED matrix.
func displayPattern(conn spi.Conn) {
	// Example pattern: diagonal line
	pattern := []byte{
		0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80,
	}
	for i, val := range pattern {
		sendCommand(conn, Digit0+byte(i), val)
	}
}

// clearDisplay clears the LED matrix display.
func clearDisplay(conn spi.Conn) {
	for i := 0; i < 8; i++ {
		sendCommand(conn, Digit0+byte(i), 0x00)
	}
}
