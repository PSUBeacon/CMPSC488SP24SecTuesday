package gocode

import (
	"fmt"
)

// Switchable is an interface for switchable devices.

// Fan is a struct that implements the Switchable interface.
type Fan struct {
	pin int
}

// NewFan creates a new Fan instance.
func NewFan(pin int) *Fan {
	return &Fan{pin: pin}
}

// SwitchOn turns the fan on.
func (f *Fan) SwitchOn() {
	fmt.Printf("Fan on pin %d is switched on.\n", f.pin)
}

// SwitchOff turns the fan off.
func (f *Fan) SwitchOff() {
	fmt.Printf("Fan on pin %d is switched off.\n", f.pin)
}

func main() {

	fan := NewFan(5) // Assuming the fan is connected to pin 5
	fan.SwitchOn()
	fan.SwitchOff()
}
