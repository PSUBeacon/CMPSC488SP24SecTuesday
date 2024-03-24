package gocode

import (
	"errors"
	"time"

	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

// DHTType represents the type of DHT sensor.
type DHTType int

// Constants for different types of DHT sensors.
const (
	DHT11 DHTType = 11
	DHT22 DHTType = 22
)

// DHT represents a DHT sensor.
type DHT struct {
	Pin  string
	Type DHTType
}

// NewDHT creates a new DHT sensor instance.
func NewDHT(pin string, sensorType DHTType) *DHT {
	return &DHT{
		Pin:  pin,
		Type: sensorType,
	}
}

// ReadTemperature reads the temperature from the DHT sensor.
func (d *DHT) ReadTemperature() (float64, error) {
	data, err := d.readData()
	if err != nil {
		return 0, err
	}

	switch d.Type {
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
func (d *DHT) ReadHumidity() (float64, error) {
	data, err := d.readData()
	if err != nil {
		return 0, err
	}

	switch d.Type {
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
func (d *DHT) readData() ([5]byte, error) {
	var data [5]byte

	r := raspi.NewAdaptor()
	pin := gpio.NewDirectPinDriver(r, d.Pin)

	// Send start signal
	pin.DigitalWrite(0)
	time.Sleep(18 * time.Millisecond)
	pin.DigitalWrite(1)
	time.Sleep(20 * time.Microsecond)

	// Set pin to input and wait for response
	pin.DigitalRead()
	time.Sleep(80 * time.Microsecond)

	// Read data
	var bits [40]bool
	for i := 0; i < 40; i++ {
		// Wait for the pin to go high
		timeout := time.After(1 * time.Millisecond)
		for val, _ := pin.DigitalRead(); val == 0; val, _ = pin.DigitalRead() {
			select {
			case <-timeout:
				return data, errors.New("timeout waiting for high signal")
			default:
			}
		}

		// Measure the length of the high signal
		start := time.Now()
		timeout = time.After(1 * time.Millisecond)
		for val, _ := pin.DigitalRead(); val == 1; val, _ = pin.DigitalRead() {
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

//func main() {
//	// Create a new DHT22 sensor instance.
//	dhtSensor := NewDHT("4", DHT22) // Assuming the sensor is connected to GPIO pin 4
//
//	// Read temperature and humidity.
//	temp, err := dhtSensor.ReadTemperature()
//	if err != nil {
//		log.Fatal("Failed to read temperature:", err)
//	}
//	humidity, err := dhtSensor.ReadHumidity()
//	if err != nil {
//		log.Fatal("Failed to read humidity:", err)
//	}
//
//	fmt.Printf("Temperature: %.2f°C, Humidity: %.2f%%\n", temp, humidity)
//}
