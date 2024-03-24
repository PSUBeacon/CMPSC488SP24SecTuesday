package gocode

import (
	"errors"
	"log"
	"periph.io/x/periph/conn/physic"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

// MAX7219 registers
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

// MaxMatrix represents a MAX7219 LED matrix.
type MaxMatrix struct {
	port    spi.PortCloser
	conn    spi.Conn
	loadPin gpio.PinOut
	buffer  [8]byte
}

// NewMaxMatrix creates a new MaxMatrix instance.
func NewMaxMatrix(spiPort, loadPin string) (*MaxMatrix, error) {
	if _, err := host.Init(); err != nil {
		return nil, err
	}

	port, err := spireg.Open(spiPort)
	if err != nil {
		return nil, err
	}

	conn, err := port.Connect(10*physic.MegaHertz, spi.Mode0, 8)
	if err != nil {
		port.Close()
		return nil, err
	}

	pin := gpioreg.ByName(loadPin)
	if pin == nil {
		port.Close()
		return nil, errors.New("invalid load pin")
	}
	pin.Out(gpio.Low)

	return &MaxMatrix{
		port:    port,
		conn:    conn,
		loadPin: pin,
	}, nil
}

// Close releases the resources used by the MaxMatrix.
func (m *MaxMatrix) Close() error {
	if err := m.port.Close(); err != nil {
		return err
	}
	return nil
}

// Init initializes the MAX7219 LED matrix.
func (m *MaxMatrix) Init() {
	m.sendCommand(Max7219RegScanLimit, 0x07)
	m.sendCommand(Max7219RegDecodeMode, 0x00)
	m.sendCommand(Max7219RegShutdown, 0x01)
	m.sendCommand(Max7219RegDisplayTest, 0x00)
	m.Clear()
	m.sendCommand(Max7219RegIntensity, 0x0f)
}

// Clear clears the LED matrix.
func (m *MaxMatrix) Clear() {
	for i := 0; i < 8; i++ {
		m.buffer[i] = 0
		m.sendCommand(byte(Max7219RegDigit0+i), 0)
	}
}

// SetIntensity sets the intensity of the LED matrix.
func (m *MaxMatrix) SetIntensity(intensity byte) {
	m.sendCommand(Max7219RegIntensity, intensity)
}

// SetPixel sets the state of a single pixel.
func (m *MaxMatrix) SetPixel(x, y int, value bool) {
	if x < 0 || x >= 8 || y < 0 || y >= 8 {
		return
	}

	if value {
		m.buffer[y] |= 1 << uint(x)
	} else {
		m.buffer[y] &^= 1 << uint(x)
	}
	m.sendCommand(byte(Max7219RegDigit0+y), m.buffer[y])
}

// sendCommand sends a command to the MAX7219.
func (m *MaxMatrix) sendCommand(register, data byte) {
	m.loadPin.Out(gpio.Low)
	m.conn.Tx([]byte{register, data}, nil)
	m.loadPin.Out(gpio.High)
}

func main() {
	matrix, err := NewMaxMatrix("SPI0.0", "GPIO10")
	if err != nil {
		log.Fatalf("Failed to create MaxMatrix: %v", err)
	}
	defer matrix.Close()

	matrix.Init()
	matrix.SetIntensity(0x08)

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			matrix.SetPixel(x, y, true)
			time.Sleep(100 * time.Millisecond)
			matrix.SetPixel(x, y, false)
		}
	}

	// Display a pattern
	pattern := []byte{
		0b00111100,
		0b01000010,
		0b10011001,
		0b10100101,
		0b10000001,
		0b10000001,
		0b01000010,
		0b00111100,
	}
	for i, row := range pattern {
		matrix.sendCommand(byte(Max7219RegDigit0+i), row)
	}
	time.Sleep(5 * time.Second)

	matrix.Clear()
}
