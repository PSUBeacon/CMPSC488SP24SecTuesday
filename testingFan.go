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

func main() {
	// Open and map memory to access GPIO, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println("Unable to open GPIO:", err)
		os.Exit(1)
	}
	defer rpio.Close()

	// Set pin to PWM mode
	pin := rpio.Pin(18)
	pin.Mode(rpio.Pwm)
	pin.Freq(19200000) // Set PWM frequency
	pin.DutyCycle(0, 1)

	// Set the fan to low speed
	pin.DutyCycle(1920000, 19200000) // 10% duty cycle
	fmt.Println("Fan LOW")
	time.Sleep(2 * time.Second)

	// Set the fan to medium speed
	pin.DutyCycle(9600000, 19200000) // 50% duty cycle
	fmt.Println("Fan MEDIUM")
	time.Sleep(2 * time.Second)

	// Set the fan to high speed
	pin.DutyCycle(17280000, 19200000) // 90% duty cycle
	fmt.Println("Fan HIGH")
	time.Sleep(2 * time.Second)

	// Turn the fan off
	pin.DutyCycle(0, 1)
	fmt.Println("Fan OFF")
}
