package main

import (
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/host/v3"
)

// Start starts the Fan.
func Start(pin gpio.PinIO) {
	// Generate a 33% duty cycle 10KHz signal.
	if err := pin.PWM(gpio.DutyMax/3, 440*physic.Hertz); err != nil {
		log.Fatal(err)
	}
}

// Stop stops the Fan.
func Stop(pin gpio.PinIO) {
	if err := pin.Halt(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Use gpioreg GPIO pin registry to find a GPIO pin by name.
	pin := gpioreg.ByName("GPIO4")
	if pin == nil {
		log.Fatalf("Failed to find GPIO4")
	}

	// Start the fan
	Start(pin)

	// Stop the fan after 10 seconds
	time.Sleep(10 * time.Second)
	Stop(pin)
}
