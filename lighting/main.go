package main

import (
	"fmt"
)

// Lighting represents a lighting fixture with on/off and brightness control.
type Lighting struct {
	Name       string
	State      bool
	Brightness int // 0-100 (0% to 100% brightness)
}

// NewLighting creates a new Lighting instance with the given name and initial state.
func NewLighting(name string, initialState bool) *Lighting {
	return &Lighting{
		Name:       name,
		State:      initialState,
		Brightness: 0, // Initialize with 0% brightness
	}
}

// TurnOn turns the lighting on.
func (l *Lighting) TurnOn() {
	l.State = true
	fmt.Printf("%s is now turned ON\n", l.Name)
}

// TurnOff turns the lighting off.
func (l *Lighting) TurnOff() {
	l.State = false
	fmt.Printf("%s is now turned OFF\n", l.Name)
}

// SetBrightness sets the brightness of the lighting.
func (l *Lighting) SetBrightness(brightness int) {
	if brightness < 0 {
		brightness = 0
	} else if brightness > 100 {
		brightness = 100
	}
	l.Brightness = brightness
	fmt.Printf("%s brightness is set to %d%%\n", l.Name, l.Brightness)
}

func main() {
	// Create a new lighting fixture
	livingRoomLight := NewLighting("Living Room Light", false)

	// Use the lighting fixture
	livingRoomLight.TurnOn()
	livingRoomLight.SetBrightness(75)
	livingRoomLight.TurnOff()
}
