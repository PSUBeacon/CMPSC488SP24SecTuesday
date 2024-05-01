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
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	// Define the GPIO pin
	servoPin := rpio.Pin(12) // Use the correct pin for your setup

	// Set the pin to PWM mode
	servoPin.Mode(rpio.Pwm)

	// Set PWM frequency to 50Hz
	servoPin.Freq(50)

	// Define the duty cycle for 90 degrees (adjust as needed for your servo)
	const dutyCycle90 = 1500 // Duty cycle in microseconds

	// Set the duty cycle to move the servo to 90 degrees
	servoPin.DutyCycle(dutyCycle90, 20000) // 20000 microseconds = 20ms period

	// Wait for the servo to move
	time.Sleep(1 * time.Second)

	// Clean up the GPIO pin
	servoPin.Mode(rpio.Output) // Set the pin to output mode
	servoPin.Low()             // Set the pin to low state
	servoPin.Mode(rpio.Input)  // Set the pin back to input mode
}
