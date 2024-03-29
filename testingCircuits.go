package main

import (
	"log"
	"periph.io/x/periph/host"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

const (
	csPin  = "GPIO4"
	dinPin = "GPIO9"
	clkPin = "GPIO10"
)

var (
	cs  gpio.PinOut
	din gpio.PinOut
	clk gpio.PinOut
)

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	cs = gpioreg.ByName(csPin)
	din = gpioreg.ByName(dinPin)
	clk = gpioreg.ByName(clkPin)

	if cs == nil || din == nil || clk == nil {
		log.Fatal("Failed to find GPIO pins")
	}

	if err := cs.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}
	if err := din.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}
	if err := clk.Out(gpio.Low); err != nil {
		log.Fatal(err)
	}

	initializeMatrix()
	clearMatrix()
	drawA()
	time.Sleep(30 * time.Second)
	clearMatrix()
}

func sendByte(data byte) {
	for i := 0; i < 8; i++ {
		din.Out(gpio.Level(data&0x80 != 0))
		clk.Out(gpio.High)
		time.Sleep(100 * time.Microsecond)
		clk.Out(gpio.Low)
		data <<= 1
	}
}

func sendCommand(register, data byte) {
	cs.Out(gpio.Low)
	sendByte(register)
	sendByte(data)
	cs.Out(gpio.High)
}

func initializeMatrix() {
	sendCommand(0x0B, 0x07)
	sendCommand(0x0A, 0x04)
	sendCommand(0x0C, 0x01)
	sendCommand(0x0F, 0x00)
}

func clearMatrix() {
	for i := 1; i <= 8; i++ {
		sendCommand(byte(i), 0x00)
	}
}

func drawA() {
	a := []byte{
		0b00011000,
		0b00111100,
		0b01100110,
		0b01100110,
		0b01111110,
		0b01100110,
		0b01100110,
		0b01100110,
	}
	for row, pattern := range a {
		sendCommand(byte(row+1), pattern)
	}
}
func drawLightbulb() {
	clearMatrix()
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
		sendCommand(byte(row+1), pattern)
	}
}

func drawLock() {
	clearMatrix()
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
		sendCommand(byte(row+1), pattern)
	}
}

func drawH() {
	clearMatrix()
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
		sendCommand(byte(row+1), pattern)
	}
}
