package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	pinNumber = 4
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println("Error opening GPIO:", err)
		os.Exit(1)
	}
	defer rpio.Close()

	pin := rpio.Pin(pinNumber)
	pin.Input() // Start with the pin in input mode

	triggerSensor(pin)

	data := readData(pin)

	fmt.Printf("Read data: %v\n", data)
}

func triggerSensor(pin rpio.Pin) {
	pin.Output()
	pin.Low()
	time.Sleep(18 * time.Millisecond) // Required to trigger the sensor
	pin.High()
	time.Sleep(20 * time.Microsecond) // Short delay before setting to input mode
	pin.Input()
}

func readData(pin rpio.Pin) []byte {
	var data []byte
	time.Sleep(40 * time.Millisecond) // Wait for sensor response
	for i := 0; i < 40; i++ {
		state := pin.Read()

		if state == rpio.High {
			data = append(data, 1)
		} else {
			data = append(data, 0)
		}
	}
	return data
}
