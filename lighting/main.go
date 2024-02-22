package main

import (
	"fmt"
	"time"
)

// Lighting represents a lighting fixture with on/off and brightness control.
type Lighting struct {
	Name              string
	Location          string // Location of the light bulb
	Brightness        string // 0-100 (0% to 100% brightness)
	Status            bool   // on/off
	EnergyConsumption int    // How much energy did the lights use (in kilowatts)
	LastChanged       string // log the time of last used
}

// NewLighting creates a new Lighting instance with the given name and initial state.
func NewLighting(name string, initialState bool) *Lighting {
	return &Lighting{
		Name:        name,
		Status:      initialState,
		Brightness:  "0",                             // Initialize with 0% brightness
		LastChanged: time.Now().Format(time.RFC3339), // Capture the creation time
	}
}

// TurnOn turns the lighting on.
func (l *Lighting) TurnOn() {
	l.Status = true
	l.LastChanged = time.Now().Format(time.RFC3339)
	fmt.Printf("%s is now turned ON\n", l.Name)
}

// TurnOff turns the lighting off.
func (l *Lighting) TurnOff() {
	l.Status = false
	l.LastChanged = time.Now().Format(time.RFC3339)
	fmt.Printf("%s is now turned OFF\n", l.Name)
}

// SetBrightness sets the brightness of the lighting.
func (l *Lighting) SetBrightness(brightness int) {
	if brightness < 0 {
		brightness = 0
	} else if brightness > 100 {
		brightness = 100
	}
	l.Brightness = fmt.Sprintf("%d%%", brightness)
	l.LastChanged = time.Now().Format(time.RFC3339)
	fmt.Printf("%s brightness is set to %s\n", l.Name, l.Brightness)
}

func main() {
	// Create a new lighting fixture
	livingRoomLight := NewLighting("Living Room Light", false)

	// Use the lighting fixture
	livingRoomLight.TurnOn()
	livingRoomLight.SetBrightness(75)
	livingRoomLight.TurnOff()
}
