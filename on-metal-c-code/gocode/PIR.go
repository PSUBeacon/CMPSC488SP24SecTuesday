package gocode

import (
	"errors"
	"log"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

// PIR represents a passive infrared (PIR) sensor.
type PIR struct {
	signalPin gpio.PinIn
}

// NewPIR creates a new PIR instance.
func NewPIR(pinName string) (*PIR, error) {
	if _, err := host.Init(); err != nil {
		return nil, err
	}

	pin := gpioreg.ByName(pinName)
	if pin == nil {
		return nil, errors.New("invalid GPIO pin")
	}

	if err := pin.In(gpio.PullDown, gpio.NoEdge); err != nil {
		return nil, err
	}

	return &PIR{
		signalPin: pin,
	}, nil
}

// Read reads the state of the PIR sensor.
func (p *PIR) Read() bool {
	return p.signalPin.Read() == gpio.High
}

func main() {
	pir, err := NewPIR("GPIO17") // Adjust the GPIO pin according to your setup
	if err != nil {
		log.Fatalf("Failed to initialize PIR sensor: %v", err)
	}

	for {
		motionDetected := pir.Read()
		if motionDetected {
			log.Println("Motion detected!")
		} else {
			log.Println("No motion.")
		}
		time.Sleep(1 * time.Second)
	}
}
