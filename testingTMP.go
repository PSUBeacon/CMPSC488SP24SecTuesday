package main

import (
	"fmt"
	"github.com/d2r2/go-dht"
	"log"
)

func main() {
	// Use GPIO pin 4, for example, change this to the pin you've connected your sensor to.
	// Make sure to run your program with root permissions to access GPIO pins.
	pin := 4

	// Specify the sensor type (DHT11 or DHT22)
	sensorType := dht.DHT11

	// Read data from the sensor
	temperature, humidity, _, err := dht.ReadDHTxxWithRetry(sensorType, pin, false, 10)
	if err != nil {
		log.Fatalf("Failed to read from sensor: %s", err)
	}

	// Print the results based on the sensor type
	if sensorType == dht.DHT11 {
		fmt.Printf("Temperature: %.0f°F\n", (temperature*9/5)+32)
		fmt.Printf("Humidity: %.0f%%\n", humidity)
	} else if sensorType == dht.DHT22 {
		fmt.Printf("Temperature: %.2f°F\n", (temperature*9/5)+32)
		fmt.Printf("Humidity: %.2f%%\n", humidity)
	}
}
