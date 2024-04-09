package gocode

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	LowSpeed    = 30
	MediumSpeed = 50
	HighSpeed   = 90
)

// SetFanSpeed controls the fan speed using software PWM
func SetFanSpeed(pin rpio.Pin, speed int) {
	if err := rpio.Open(); err != nil {
		fmt.Println("Unable to open GPIO:", err)
		os.Exit(1)
	}
	defer rpio.Close()

	pin.Output()

	if speed == 0 {
		TurnOffFan(pin)
		return
	}

	onTime := time.Duration(speed) * time.Millisecond
	offTime := 100*time.Millisecond - onTime
	for i := 0; i < 100; i++ { // Run PWM for a short period
		pin.High()
		time.Sleep(onTime)
		pin.Low()
		time.Sleep(offTime)
	}
}

// TurnOffFan turns off the fan
func TurnOffFan(pin rpio.Pin) {
	if err := rpio.Open(); err != nil {
		fmt.Println("Unable to open GPIO:", err)
		os.Exit(1)
	}
	defer rpio.Close()

	pin.Output()
	pin.Low()
}

//func main() {
//	// Open and map memory to access GPIO, check for errors
//	if err := rpio.Open(); err != nil {
//		fmt.Println("Unable to open GPIO:", err)
//		os.Exit(1)
//	}
//	defer rpio.Close()
//
//	// Set pin to output mode
//	pin := rpio.Pin(18)
//	pin.Output()
//
//	// Set the fan to low speed
//	fancontrol.SetFanSpeed(pin, fancontrol.LowSpeed)
//	fmt.Println("Fan LOW")
//	time.Sleep(2 * time.Second)
//
//	// Set the fan to medium speed
//	fancontrol.SetFanSpeed(pin, fancontrol.MediumSpeed)
//	fmt.Println("Fan MEDIUM")
//	time.Sleep(2 * time.Second)
//
//	// Set the fan to high speed
//	fancontrol.SetFanSpeed(pin, fancontrol.HighSpeed)
//	fmt.Println("Fan HIGH")
//	time.Sleep(2 * time.Second)
//
//	// Turn the fan off
//	fancontrol.TurnOffFan(pin)
//	fmt.Println("Fan OFF")
//}
