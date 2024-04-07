package main

import (
	"fmt"
	"github.com/d2r2/go-dht"
	"log"
)

func main() {
	// Use GPIO pin 4, for example, change this to the pin you've connected your DHT22 sensor to.
	// Make sure to run your program with root permissions to access GPIO pins.
	sensorType := dht.DHT22

	// Read data from the sensor
	temperature, humidity, _, err := dht.ReadDHTxxWithRetry(sensorType, 4, false, 10)
	if err != nil {
		log.Fatalf("Failed to read from DHT22 sensor: %s", err)
	}

	fmt.Printf("Temperature: %.2f°F\n", (temperature*9/5)+32)
	fmt.Printf("Humidity: %.2f%%\n", humidity)
}
