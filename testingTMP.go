package main

import (
	"fmt"
	"log"
	"time"

	"github.com/d2r2/go-dht"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	if err := rpio.Open(); err != nil {
		log.Fatalf("Failed to open GPIO: %s", err)
	}
	defer rpio.Close()

	pin := 4
	time.Sleep(2 * time.Second) // Allow sensor to stabilize

	sensorType := dht.DHT11

	temperature, humidity, _, err := dht.ReadDHTxxWithRetry(sensorType, pin, true, 10)
	if err != nil {
		log.Fatalf("Failed to read DHT11: %s", err)
	}

	fmt.Printf("Temperature: %.2f°C, Humidity: %.2f%%\n", temperature, humidity)
}
