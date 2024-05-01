package gocode

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	gpioPin = 26 // GPIO pin connected to the buzzer
)

func BuzzerStatus() {

	// Open /dev/gpiomem using the go-rpio library
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	// Configure the pin connected to the buzzer as output
	pin := rpio.Pin(gpioPin)
	pin.Output()

	fmt.Println("Playing pulsing tone...")

	// Run the pulsing tone loop
	playPulsingTone(&pin, 10)

	// Create a channel to receive OS signals for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Wait for a termination signal
	<-c
	fmt.Println("\nReceived termination signal. Exiting.")
}

// playPulsingTone plays a pulsing tone for the given duration (in seconds)
func playPulsingTone(pin *rpio.Pin, duration int) {
	toggleDuration := 500 * time.Millisecond // 500 milliseconds = 0.5 second

	endTime := time.Now().Add(time.Duration(duration) * time.Second)
	for time.Now().Before(endTime) {
		// Toggle the pin high
		pin.High()
		time.Sleep(toggleDuration)

		// Toggle the pin low
		pin.Low()
		time.Sleep(toggleDuration)
	}
}
