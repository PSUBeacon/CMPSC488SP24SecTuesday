package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

// CheckForMotion initializes the GPIO then checks for motion
func CheckForMotion(pinNum uint8) (bool, error) {
	if err := rpio.Open(); err != nil {
		return false, err
	}
	defer func() {
		err := rpio.Close()
		if err != nil {

		}
	}()

	pin := rpio.Pin(pinNum)
	pin.Input()
	pin.PullDown()

	motionDetected := pin.Read() == rpio.High
	return motionDetected, nil
}

func main() {
	for {
		motion, err := CheckForMotion(11)
		if err != nil {
			fmt.Printf("Error checking for motion: %v\n", err)
			return
		}
		fmt.Printf("Motion detected: %t\n", motion)
		time.Sleep(1 * time.Second)
	}
}
