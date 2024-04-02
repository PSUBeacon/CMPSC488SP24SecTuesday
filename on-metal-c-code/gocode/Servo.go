package gocode

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func TurnServoMotor90Degrees() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	pin := rpio.Pin(18)
	pin.Mode(rpio.Pwm)
	pin.Freq(50)

	// Set the servo to the 90-degree position
	pin.DutyCycle(0, 7.5)

	// Wait for the servo to reach the desired position
	time.Sleep(300 * time.Millisecond)
}
