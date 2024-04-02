package gocode

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func TurnServoTo90Degrees() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	// Define the GPIO pin
	servoPin := rpio.Pin(18) // Use the correct pin for your setup
	servoPin.Output()        // Set the pin to output mode

	// Manually generate the PWM signal for 90 degrees
	// Adjust the pulse width for your specific servo
	pulseWidth := 1500 * time.Microsecond // Pulse width for 90 degrees
	period := 20 * time.Millisecond       // Period of the PWM signal (20ms is common for servos)

	start := time.Now()
	for time.Since(start) < 1*time.Second { // Generate the signal for 1 second
		servoPin.High()                 // Set the pin high
		time.Sleep(pulseWidth)          // Wait for the pulse width duration
		servoPin.Low()                  // Set the pin low
		time.Sleep(period - pulseWidth) // Wait for the rest of the period
	}
}
