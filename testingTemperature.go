package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

type DHT struct {
	PinName string
	Pin     gpio.PinIO
}

func NewDHT(pinName string) *DHT {
	// Initialize periph.io host for GPIO access
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	pin := gpioreg.ByName(pinName)
	if pin == nil {
		log.Fatalf("Failed to find pin %s", pinName)
	}

	return &DHT{
		PinName: pinName,
		Pin:     pin,
	}
}

func (d *DHT) ReadTemperatureAndHumidity() (temperature float64, humidity float64, err error) {
	// Send start signal
	d.Pin.Out(gpio.Low)
	time.Sleep(18 * time.Millisecond) // hold for at least 18 ms
	d.Pin.Out(gpio.High)
	time.Sleep(20 * time.Microsecond) // back to high, wait 20-40 us

	// Switch to input to read data
	d.Pin.In(gpio.PullUp, gpio.NoEdge)

	// Read the sensor response and the 40 bits of data
	var data [5]byte
	for i := range data {
		for j := 7; j >= 0; j-- {
			// Wait for the pin to go high
			if !waitForPinState(d.Pin, gpio.High, 50*time.Microsecond) {
				return 0, 0, errors.New("timeout waiting for response")
			}
			// Measure the length of the high state
			length := measureHighState(d.Pin)
			if length > 28 { // If the high state is long, it's a 1
				data[i] |= 1 << j
			}
			// else it's a 0, and we do nothing because the bit is already 0
		}
	}

	// Validate checksum
	if data[4] != (data[0]+data[1]+data[2]+data[3])&0xFF {
		return 0, 0, errors.New("checksum mismatch")
	}

	// Convert to temperature and humidity

	humidity = float64(uint16(data[0])<<8|uint16(data[1])) * 0.1
	temperature = float64(uint16(data[2]&0x7F)<<8|uint16(data[3])) * 0.1
	if data[2]&0x80 > 0 {
		temperature *= -1
	}

	return temperature, humidity, nil
}

// waitForPinState waits for the pin to reach the desired state until the timeout.
// Returns true if the desired state was reached, false on timeout.
func waitForPinState(pin gpio.PinIO, state gpio.Level, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if pin.Read() == state {
			return true
		}
		time.Sleep(1 * time.Microsecond)
	}
	return false
}

// measureHighState measures the duration of the high state of the pin.
// This is a simplified approximation and may not be precise.
func measureHighState(pin gpio.PinIO) time.Duration {
	start := time.Now()
	for pin.Read() == gpio.High {
		// Spin wait
	}
	return time.Since(start)
}

func main() {
	dhtSensor := NewDHT("GPIO22") // Replace GPIO22 with your actual GPIO pin name

	for {
		temp, humidity, err := dhtSensor.ReadTemperatureAndHumidity()
		if err != nil {
			log.Printf("Error reading from DHT sensor: %v", err)
			continue
		}

		fmt.Printf("Temperature: %.2f°C, Humidity: %.2f%%\n", temp, humidity)
		time.Sleep(2 * time.Second)
	}
}
