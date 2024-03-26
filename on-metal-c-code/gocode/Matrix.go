package gocode

import (
	"github.com/stianeikeland/go-rpio/v4"
)

// registers
const (
	Max7219RegNoop        = 0x00
	Max7219RegDigit0      = 0x01
	Max7219RegDigit1      = 0x02
	Max7219RegDigit2      = 0x03
	Max7219RegDigit3      = 0x04
	Max7219RegDigit4      = 0x05
	Max7219RegDigit5      = 0x06
	Max7219RegDigit6      = 0x07
	Max7219RegDigit7      = 0x08
	Max7219RegDecodeMode  = 0x09
	Max7219RegIntensity   = 0x0a
	Max7219RegScanLimit   = 0x0b
	Max7219RegShutdown    = 0x0c
	Max7219RegDisplayTest = 0x0f
)

var loadPin rpio.Pin
var buffer [8]byte

// Initialize initializes the LED matrix.
func Initialize(loadPinNum rpio.Pin) error {
	if err := rpio.Open(); err != nil {
		return err
	}

	err := rpio.SpiBegin(rpio.Spi0)
	if err != nil {
		return err
	}
	rpio.SpiSpeed(10000000)
	rpio.SpiMode(0, 0)

	loadPin = loadPinNum
	loadPin.Output()
	loadPin.Low()

	sendCommand(Max7219RegScanLimit, 0x07)
	sendCommand(Max7219RegDecodeMode, 0x00)
	sendCommand(Max7219RegShutdown, 0x01)
	sendCommand(Max7219RegDisplayTest, 0x00)
	Clear()
	sendCommand(Max7219RegIntensity, 0x0f)

	return nil
}

// Clear clears the LED matrix.
func Clear() {
	for i := 0; i < 8; i++ {
		buffer[i] = 0
		sendCommand(byte(Max7219RegDigit0+i), 0)
	}
}

// SetIntensity sets the intensity of the LED matrix.
func SetIntensity(intensity byte) {
	sendCommand(Max7219RegIntensity, intensity)
}

// SetPixel sets the state of a single pixel.
func SetPixel(x, y int, value bool) {
	if x < 0 || x >= 8 || y < 0 || y >= 8 {
		return
	}

	if value {
		buffer[y] |= 1 << uint(x)
	} else {
		buffer[y] &^= 1 << uint(x)
	}
	sendCommand(byte(Max7219RegDigit0+y), buffer[y])
}

// sendCommand sends a command to the LED matrix
func sendCommand(register, data byte) {
	loadPin.Low()
	rpio.SpiTransmit(register, data)
	loadPin.High()
}

// Cleanup releases the resources used
func Cleanup() {
	rpio.SpiEnd(rpio.Spi0)
	err := rpio.Close()
	if err != nil {
		return
	}
}
