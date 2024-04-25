package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/d2r2/go-dht"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	// Redirect all log output to ioutil.Discard, effectively silencing all logs
	log.SetOutput(ioutil.Discard)

	// Initialize GPIO
	if err := rpio.Open(); err != nil {
		fmt.Println("Failed to open GPIO:", err)
		return // Exit if unable to open GPIO, using fmt.Println to output the error
	}
	defer func() {
		err := rpio.Close()
		if err != nil {
			fmt.Println("Failed to close GPIO:", err)
		}
	}()

	// Define your DHT sensor type and GPIO pin
	sensorType := dht.DHT22
	pinNumber := 4 // Use the BCM pin number connected to your DHT22 sensor

	// Prepare the GPIO pin (optional, since go-dht handles GPIO)
	pin := rpio.Pin(pinNumber)
	pin.Output()                // Set pin to output mode
	pin.High()                  // Set pin high (DHT22 requires pull-up resistor)
	time.Sleep(1 * time.Second) // Stabilize sensor

	// Switch back to input mode before reading (go-dht manages this internally)
	pin.Input()

	// Now, read from the DHT22 sensor using go-dht
	temperature, humidity, _, err := dht.ReadDHTxxWithRetry(sensorType, pinNumber, false, 10)
	if err != nil {
		fmt.Println("Failed to read from DHT22 sensor:", err)
		return // Exit if there was an error reading from the sensor, using fmt.Println to output the error
	}

	// Output the results
	fmt.Printf("Temperature: %.2f°C\n", temperature)
	fmt.Printf("Humidity: %.2f%%\n", humidity)
}
