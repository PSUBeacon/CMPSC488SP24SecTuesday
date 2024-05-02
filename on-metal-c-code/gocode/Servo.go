package gocode

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"periph.io/x/host/v3"
	"syscall"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

const (
	servoPinName     = "GPIO8"
	pulseFrequency   = 20 * time.Millisecond // Common period for servo control
	minPulseWidth    = 600 * time.Microsecond
	maxPulseWidth    = 2400 * time.Microsecond
	rotationDuration = 4 * time.Second // Duration to send the signal, allowing full rotation
)

func TurnServo(stat bool) {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	servoPin := gpioreg.ByName(servoPinName)
	if servoPin == nil {
		log.Fatalf("Failed to find pin %s", servoPinName)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()
	if stat == true {
		setServoAngle(servoPin, maxPulseWidth) // Move to one side
	}
	if stat == false {
		setServoAngle(servoPin, minPulseWidth) // Move to the opposite side
	}
}

func setServoAngle(pin gpio.PinIO, pulseWidth time.Duration) {
	end := time.Now().Add(rotationDuration)
	for time.Now().Before(end) {
		// Send the pulse
		pin.Out(gpio.High)
		time.Sleep(pulseWidth)
		pin.Out(gpio.Low)
		time.Sleep(pulseFrequency - pulseWidth)
	}
}
