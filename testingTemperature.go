package main

import (
	"bytes"
	"fmt"
	"github.com/d2r2/go-dht"
	"log"
)

func main() {
	pinNumber := 4 // GPIO pin connected to DHT22

	// Use the ReadDHT22SensorData function to get temperature and humidity
	temperature, humidity, err := ReadDHT22SensorData(pinNumber)
	if err != nil {
		log.Fatalf("Error reading sensor data: %v", err)
	}

	// Print temperature and humidity
	fmt.Printf("Temperature: %.2f°F\n", (temperature*9/5)+32)
	fmt.Printf("Humidity: %.2f%%\n", humidity)
}

func ReadDHT22SensorData(pinNumber int) (float64, float64, error) {
	// Temporarily redirect log output to suppress DEBUG and WARN messages
	originalLogger := log.Default()
	defer log.SetOutput(originalLogger.Writer())

	var logBuffer bytes.Buffer
	log.SetOutput(&logBuffer)

	// Initialize the DHT22 sensor
	sensorType := dht.DHT22

	// Read data from the sensor
	temp, hum, _, err := dht.ReadDHTxxWithRetry(sensorType, pinNumber, false, 10)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read from DHT22 sensor: %w", err)
	}

	// Convert float32 to float64
	temperature := float64(temp)
	humidity := float64(hum)

	// Process the log buffer if necessary (e.g., checking for critical errors)
	// Example: log.Println(logBuffer.String())

	// Return the temperature and humidity readings as float64
	return temperature, humidity, nil
}
