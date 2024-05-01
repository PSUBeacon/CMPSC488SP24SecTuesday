package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"os"
	"time"
)

func main() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	// Define the GPIO pin
	servoPin := rpio.Pin(18) // Use the correct pin for your setup
	servoPin.Mode(rpio.Pwm)  // Set the pin to PWM mode

	// Manually control PWM for servo
	const dutyCycle = 150               // Adjust this value for 90 degrees based on your servo
	servoPin.Freq(50 * 1000)            // Set frequency to 50Hz
	servoPin.DutyCycle(dutyCycle, 1000) // Set duty cycle

	// Wait for the servo to move
	time.Sleep(1 * time.Second)

	// Clean up the GPIO pin
	servoPin.Freq(0)          // Disable PWM
	servoPin.Mode(rpio.Input) // Set the pin back to input mode
}
