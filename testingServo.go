package main

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

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	servoPin := gpioreg.ByName(servoPinName)
	if servoPin == nil {
		log.Fatalf("Failed to find pin %s", servoPinName)
	}

	fmt.Println("Servo door hinge control started. Type 'open' to open or 'close' to close the door. Press Ctrl+C to exit.")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()

	for {
		var command string
		fmt.Print("Command (open/close): ")
		fmt.Scanln(&command)

		switch command {
		case "open":
			setServoAngle(servoPin, maxPulseWidth) // Move to one side
		case "close":
			setServoAngle(servoPin, minPulseWidth) // Move to the opposite side
		default:
			fmt.Println("Invalid command. Please type 'open' or 'close'.")
		}
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
