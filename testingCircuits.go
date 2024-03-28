package main

import (
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"time"
)

//
//import (
//	"flag"
//	"fmt"
//	"golang.org/x/exp/io/i2c"
//	"os"
//	"periph.io/x/periph/conn/i2c/i2creg"
//	"periph.io/x/periph/experimental/devices/hd44780"
//	"periph.io/x/periph/host"
//	"log"
//	"strings"
//	"time"
//	"periph.io/x/conn/v3/gpio"
//	"periph.io/x/conn/v3/gpio/gpioreg"
//	"periph.io/x/devices/v3/hd44780"
//	"periph.io/x/host/v3"
//)

//func main() {
//	gocode.Switchable(17, true)
//
//}

func main() {
	// Open/initialize the RPi's GPIO memory range for use
	err := rpio.Open()
	if err != nil {
		log.Fatalf("Error initializing GPIO: %v", err)
	}
	defer rpio.Close()

	// Set up pin 4 as output (Chip Select)
	csPin := rpio.Pin(4)
	csPin.Output()
	csPin.High() // Set pin high (inactive) initially

	// Set up SPI communication
	err = rpio.SpiBegin(rpio.Spi0)
	if err != nil {
		log.Fatalf("Error initializing SPI: %v", err)
	}
	defer rpio.SpiEnd(rpio.Spi0)

	rpio.SpiSpeed(1000000) // Set SPI speed to 1MHz (adjust as needed)
	rpio.SpiChipSelect(0)  // Use CS0 (this will be overridden by manual control of the CS pin)
	rpio.SpiMode(0, 0)     // Set SPI mode (adjust as needed)

	// Map physical pins to SPI function
	rpio.Pin(9).Mode(rpio.Alt0)  // Set pin 9 as SPI0 MOSI (DIN)
	rpio.Pin(10).Mode(rpio.Alt0) // Set pin 10 as SPI0 SCLK (CLK)

	// Initialize the MAX7219
	initMax7219(csPin)

	// Turn on all LEDs in the 8x8 matrix
	for row := byte(1); row <= 8; row++ {
		sendDataToMax7219(csPin, row, 0x01) // 0xFF turns on all LEDs in the row
	}

	// Add a delay to see the effect
	time.Sleep(5 * time.Second)
}

// Initialize the MAX7219 LED driver
func initMax7219(csPin rpio.Pin) {
	sendDataToMax7219(csPin, 0x0F, 0x00) // Display test register: normal operation
	sendDataToMax7219(csPin, 0x09, 0x00) // Decode mode: no decode
	sendDataToMax7219(csPin, 0x0B, 0x07) // Scan limit: display all digits
	sendDataToMax7219(csPin, 0x0A, 0x0F) // Intensity: maximum
	sendDataToMax7219(csPin, 0x0C, 0x01) // Shutdown register: normal operation
}

// Send data to the MAX7219
func sendDataToMax7219(csPin rpio.Pin, address, data byte) {
	csPin.Low()                     // Activate the CS line
	rpio.SpiTransmit(address, data) // Send the address and data
	csPin.High()                    // Deactivate the CS line
}
