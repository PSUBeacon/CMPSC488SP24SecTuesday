package gocode

import (
	"fmt"
)

// Switchable is an interface for switchable devices.
type Switchable interface {
	SwitchOn()
	SwitchOff()
}

// Buzzer is a struct that implements the Switchable interface.
type Buzzer struct {
	pin int
}

// NewBuzzer creates a new Buzzer instance.
func NewBuzzer(pin int) *Buzzer {
	return &Buzzer{pin: pin}
}

// SwitchOn turns the buzzer on.
func (b *Buzzer) SwitchOn() {
	fmt.Printf("Buzzer on pin %d is switched on.\n", b.pin)
}

// SwitchOff turns the buzzer off.
func (b *Buzzer) SwitchOff() {
	fmt.Printf("Buzzer on pin %d is switched off.\n", b.pin)
}
