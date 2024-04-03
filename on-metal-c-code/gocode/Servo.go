package gocode

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"os"
	"time"
)

func TurnServo() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	// Define the GPIO pin
	servoPin := rpio.Pin(18) // Use the correct pin for your setup
	servoPin.Mode(rpio.Pwm)  // Set the pin to PWM mode

	// Set the PWM parameters for the servo
	const (
		period    = 20 * time.Millisecond
		fullCycle = 20 // This is equivalent to 50 Hz
		dutyCycle = 7  // Duty cycle for 90 degrees, adjust as needed
	)

	// Calculate the number of cycles for the given duty cycle
	cycles := uint32((fullCycle * dutyCycle) / 10)

	servoPin.DutyCycle(cycles, fullCycle)
	time.Sleep(300 * time.Millisecond)
}
