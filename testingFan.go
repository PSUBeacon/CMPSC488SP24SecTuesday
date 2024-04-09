//package main
//
//import (
//	"log"
//	"time"
//
//	"periph.io/x/conn/v3/gpio"
//	"periph.io/x/conn/v3/gpio/gpioreg"
//	"periph.io/x/conn/v3/physic"
//)
//
//// SetFanSpeed sets the fan speed to low, medium, or high.
//func SetFanSpeed(pin gpio.PinIO, speed string) {
//	var duty gpio.Duty
//	switch speed {
//	case "low":
//		duty = gpio.DutyMax / 4 // 25% duty cycle for low speed
//	case "medium":
//		duty = gpio.DutyMax / 2 // 50% duty cycle for medium speed
//	case "high":
//		duty = gpio.DutyMax * 3 / 4 // 75% duty cycle for high speed
//	default:
//		log.Fatalf("Invalid speed setting: %s", speed)
//	}
//
//	// Generate signal with specified duty cycle at 10KHz
//	if err := pin.PWM(duty, 440*physic.Hertz); err != nil {
//		log.Fatal(err)
//	}
//}
//
//// TurnFanOn starts the fan at a specified speed.
//func TurnFanOn(pin gpio.PinIO, speed string) {
//	SetFanSpeed(pin, speed)
//}
//
//// TurnFanOff stops the fan.
//func TurnFanOff(pin gpio.PinIO) {
//	if err := pin.Halt(); err != nil {
//		log.Fatal(err)
//	}
//}
//
//func main() {
//	pin := gpioreg.ByName("GPIO0")
//	if pin == nil {
//		log.Fatalf("Failed to find GPIO0")
//	}
//
//	// Example: Turn the fan on at medium speed.
//	TurnFanOn(pin, "medium")
//
//	// Example: Turn the fan off after 10 seconds.
//	time.Sleep(10 * time.Second)
//	TurnFanOff(pin)
//}

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	// Open and map memory to access GPIO, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println("Unable to open GPIO:", err)
		os.Exit(1)
	}
	defer rpio.Close()

	// Set pin to output mode
	pin := rpio.Pin(14)
	pin.Output()

	for i := 0; i < 5; i++ {
		// Turn the fan on
		pin.High()
		fmt.Println("Fan ON")
		time.Sleep(2 * time.Second)

		// Turn the fan off
		pin.Low()
		fmt.Println("Fan OFF")
		time.Sleep(2 * time.Second)
	}
}
