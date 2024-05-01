package gocode

import (
	"fmt"
	"github.com/d2r2/go-dht"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

// ReadTemperature reads the temperature from the DHT sensor.
func ReadTempHum() string {
	// Redirect all log output to ioutil.Discard, effectively silencing all logs
	log.SetOutput(ioutil.Discard)

	// Initialize GPIO
	if err := rpio.Open(); err != nil {
		fmt.Println("Failed to open GPIO:", err)
		return "" // Exit if unable to open GPIO, using fmt.Println to output the error
	}

	// Define your DHT sensor type and GPIO pin
	sensorType := dht.DHT22
	pinNumber := 17 // Use the BCM pin number connected to your DHT22 sensor

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
		return "" // Exit if there was an error reading from the sensor, using fmt.Println to output the error
	}
	tempF := temperature*1.8 + 32
	TempHum := strconv.Itoa(int(tempF)) + "/" + strconv.Itoa(int(humidity))
	return TempHum
}
