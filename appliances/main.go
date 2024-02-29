package appliances

import (
	"fmt"
)

// Appliance represents a simple appliance with an on/off state.
type Appliance struct {
	Name  string
	State bool
	Temp  int
}

// NewAppliance creates a new Appliance instance with the given name and initial state.
func NewAppliance(name string, initialState bool, temp int) *Appliance {
	return &Appliance{
		Name:  name,
		State: initialState,
		Temp:  temp,
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

func (a *Appliance) AdjustTemp(setTemp int) {
	a.Temp = setTemp
	fmt.Printf("%s temperature is now set to %d degrees farenheit\n", a.Name, a.Temp)
}
