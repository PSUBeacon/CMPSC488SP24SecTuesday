//package main
//
//import (
//	"fmt"
//	"os"
//	"time"
//
//	"github.com/stianeikeland/go-rpio/v4"
//)
//
//func main() {
//	// Open and map memory to access GPIO, check for errors
//	if err := rpio.Open(); err != nil {
//		fmt.Println("Unable to open GPIO:", err)
//		os.Exit(1)
//	}
//	defer rpio.Close()
//
//	// Set pin to output mode
//	pin := rpio.Pin(14)
//	pin.Output()
//
//	for i := 0; i < 5; i++ {
//		// Turn the fan on
//		pin.High()
//		fmt.Println("Fan ON")
//		time.Sleep(2 * time.Second)
//
//		// Turn the fan off
//		pin.Low()
//		fmt.Println("Fan OFF")
//		time.Sleep(2 * time.Second)
//	}
//}

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	lowSpeed    = 10
	mediumSpeed = 50
	highSpeed   = 90
)

func main() {
	// Open and map memory to access GPIO, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println("Unable to open GPIO:", err)
		os.Exit(1)
	}
	defer rpio.Close()

	// Set pin to output mode
	pin := rpio.Pin(18)
	pin.Output()

	// Set the fan to low speed
	setFanSpeed(pin, lowSpeed)
	fmt.Println("Fan LOW")
	time.Sleep(2 * time.Second)

	// Set the fan to medium speed
	setFanSpeed(pin, mediumSpeed)
	fmt.Println("Fan MEDIUM")
	time.Sleep(2 * time.Second)

	// Set the fan to high speed
	setFanSpeed(pin, highSpeed)
	fmt.Println("Fan HIGH")
	time.Sleep(2 * time.Second)

	// Turn the fan off
	pin.Low()
	fmt.Println("Fan OFF")
}

// setFanSpeed controls the fan speed using software PWM
func setFanSpeed(pin rpio.Pin, speed int) {
	onTime := time.Duration(speed) * time.Millisecond
	offTime := 100*time.Millisecond - onTime
	for i := 0; i < 50; i++ { // Run PWM for a short period
		pin.High()
		time.Sleep(onTime)
		pin.Low()
		time.Sleep(offTime)
	}
}
