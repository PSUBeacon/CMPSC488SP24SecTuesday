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
	pin.Input()

	triggerSensor(pin)

	data := readData(pin)

	fmt.Printf("Read data: %v\n", data)
}

func triggerSensor(pin rpio.Pin) {
	pin.Output()
	pin.Low()
	time.Sleep(18 * time.Millisecond)
	pin.High()
	pin.Input()
}

func readData(pin rpio.Pin) []byte {
	var data []byte
	for i := 0; i < 40; i++ {
		state := pin.Read()

		if state == rpio.High {
			data = append(data, 1)
		} else {
			data = append(data, 0)
		}

		// Make sure this brace closes the for loop before the else statement.
		// An incorrectly placed brace here might cause the syntax error.
	}
	return data
}
