package main

import (
	"fmt"
)

// HVAC represents an HVAC system with temperature control, fan speed, and mode.
type HVAC struct {
	Name        string
	Temperature int    // Desired temperature in Celsius
	FanSpeed    int    // Fan speed (0-100%)
	Mode        string // HVAC mode (e.g., "Cool", "Heat", "Fan", "Off")
}

// NewHVAC creates a new HVAC instance with the given name and initial settings.
func NewHVAC(name string) *HVAC {
	return &HVAC{
		Name:        name,
		Temperature: 25,    // Initial temperature setting
		FanSpeed:    50,    // Initial fan speed setting (50%)
		Mode:        "Off", // Initial mode is Off
	}
}

// SetTemperature sets the desired temperature for the HVAC system.
func (h *HVAC) SetTemperature(temperature int) {
	h.Temperature = temperature
	fmt.Printf("%s temperature is set to %d°C\n", h.Name, h.Temperature)
}

// SetFanSpeed sets the fan speed for the HVAC system.
func (h *HVAC) SetFanSpeed(speed int) {
	if speed < 0 {
		speed = 0
	} else if speed > 100 {
		speed = 100
	}
	h.FanSpeed = speed
	fmt.Printf("%s fan speed is set to %d%%\n", h.Name, h.FanSpeed)
}

// SetMode sets the mode (e.g., "Cool", "Heat", "Fan", "Off") for the HVAC system.
func (h *HVAC) SetMode(mode string) {
	h.Mode = mode
	fmt.Printf("%s mode is set to %s\n", h.Name, h.Mode)
}

func main() {
	// Create a new HVAC system
	livingRoomHVAC := NewHVAC("Living Room HVAC")

	// Use the HVAC system
	livingRoomHVAC.SetMode("Cool")
	livingRoomHVAC.SetTemperature(22)
	livingRoomHVAC.SetFanSpeed(75)

	// Display the current settings
	fmt.Printf("Current %s settings: Mode: %s, Temperature: %d°C, Fan Speed: %d%%\n",
		livingRoomHVAC.Name, livingRoomHVAC.Mode, livingRoomHVAC.Temperature, livingRoomHVAC.FanSpeed)
}
