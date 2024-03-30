package gocode

import (
	"errors"
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

// Constants for different types of DHT sensors.
const (
	DHT11 int = 11
	DHT22 int = 22
)

// ReadTemperature reads the temperature from the DHT sensor.
func ReadTemperature(pin int, inType int) (float64, error) {
	data, err := readData(pin)
	if err != nil {
		return 0, err
	}

	switch inType {
	case DHT11:
		return float64(data[2]), nil
	case DHT22:
		rawTemp := int16(data[2]&0x7F)<<8 | int16(data[3])
		if data[2]&0x80 > 0 {
			rawTemp *= -1
		}
		return float64(rawTemp) / 10.0, nil
	default:
		return 0, errors.New("unknown sensor type")
	}
}

// ReadHumidity reads the humidity from the DHT sensor.
func ReadHumidity(pin int, inType int) (float64, error) {
	data, err := readData(pin)
	if err != nil {
		return 0, err
	}

	switch inType {
	case DHT11:
		return float64(data[0]), nil
	case DHT22:
		rawHumidity := int16(data[0])<<8 | int16(data[1])
		return float64(rawHumidity) / 10.0, nil
	default:
		return 0, errors.New("unknown sensor type")
	}
}

// readData handles the communication with the DHT sensor and returns the raw data.
func readData(inPin int) ([5]byte, error) {
	var data [5]byte

	if err := rpio.Open(); err != nil {
		return data, fmt.Errorf("unable to open GPIO: %w", err)
	}

	defer func() {
		err := rpio.Close()
		if err != nil {

		}
	}()

	pin := rpio.Pin(inPin)
	pin.Output()

	// Send start signal
	pin.Low()
	time.Sleep(18 * time.Millisecond)
	pin.High()
	time.Sleep(20 * time.Microsecond)

	// Set pin to input and wait for response
	pin.Input()
	time.Sleep(80 * time.Microsecond)

	// Read data
	var bits [40]bool
	for i := 0; i < 40; i++ {
		// Wait for the pin to go high
		timeout := time.After(1 * time.Millisecond)
		for pin.Read() == rpio.Low {
			select {
			case <-timeout:
				return data, errors.New("timeout waiting for high signal")
			default:
			}
		}

		// Measure the length of the high signal
		start := time.Now()
		timeout = time.After(1 * time.Millisecond)
		for pin.Read() == rpio.High {
			select {
			case <-timeout:
				return data, errors.New("timeout waiting for low signal")
			default:
			}
		}
		duration := time.Since(start)

		// The DHT22 sensor uses a high signal of about 70 microseconds to indicate a '1'
		// and a high signal of about 28 microseconds to indicate a '0'.
		// We use 50 microseconds as a threshold.
		bits[i] = duration > 50*time.Microsecond
	}

	// Convert bit array to bytes
	for i := 0; i < 40; i++ {
		data[i/8] <<= 1
		if bits[i] {
			data[i/8] |= 1
		}
	}

	// Verify checksum
	checksum := data[0] + data[1] + data[2] + data[3]
	if data[4] != checksum&0xFF {
		return data, errors.New("checksum mismatch")
	}

	return data, nil
}
