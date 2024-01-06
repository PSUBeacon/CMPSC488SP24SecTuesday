package main

import (
	"fmt"
)

// Appliance represents a simple appliance with an on/off state.
type Appliance struct {
	Name  string
	State bool
}

// NewAppliance creates a new Appliance instance with the given name and initial state.
func NewAppliance(name string, initialState bool) *Appliance {
	return &Appliance{
		Name:  name,
		State: initialState,
	}
}

// TurnOn turns the appliance on.
func (a *Appliance) TurnOn() {
	a.State = true
	fmt.Printf("%s is now turned ON\n", a.Name)
}

// TurnOff turns the appliance off.
func (a *Appliance) TurnOff() {
	a.State = false
	fmt.Printf("%s is now turned OFF\n", a.Name)
}

func main() {
	// Create a new light switch appliance
	lightSwitch := NewAppliance("Light Switch", false)

	// Use the appliance
	lightSwitch.TurnOn()
	lightSwitch.TurnOff()
}
