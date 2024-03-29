package gocode

import (
	"fmt"
	"log"

	"periph.io/x/periph/conn/i2c"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

// TMP102 represents a TMP102 temperature sensor.
type TMP102 struct {
	bus     i2c.Bus
	address uint16
}

// NewTMP102 creates a new TMP102 instance.
func NewTMP102(bus i2c.Bus, address uint16) *TMP102 {
	return &TMP102{
		bus:     bus,
		address: address,
	}
}

// Read reads the raw temperature data from the TMP102 sensor.
func (t *TMP102) Read() (int, error) {
	read := make([]byte, 2)
	err := t.bus.Tx(t.address, nil, read)
	if err != nil {
		return 0, err
	}
	result := int(read[0])<<8 | int(read[1])
	return result >> 4, nil // TMP102 result is 12 bits
}

// ReadCelsius reads the temperature in Celsius from the TMP102 sensor.
func (t *TMP102) ReadCelsius() (float64, error) {
	raw, err := t.Read()
	if err != nil {
		return 0, err
	}
	return float64(raw) * 0.0625, nil
}

// ReadFahrenheit reads the temperature in Fahrenheit from the TMP102 sensor.
func (t *TMP102) ReadFahrenheit() (float64, error) {
	celsius, err := t.ReadCelsius()
	if err != nil {
		return 0, err
	}
	return celsius*1.8 + 32, nil
}

func main() {
	// Initialize periph.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Open the I2C bus.
	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	// Create a new TMP102 instance.
	sensor := NewTMP102(bus, 0x48)

	// Read the temperature in Celsius.
	tempC, err := sensor.ReadCelsius()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Temperature: %.2f°C\n", tempC)

	// Read the temperature in Fahrenheit.
	tempF, err := sensor.ReadFahrenheit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Temperature: %.2f°F\n", tempF)
}
