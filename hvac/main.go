package main

import (
	"fmt"
)

// HVAC represents an HVAC system with temperature control, fan speed, and mode.
type HVAC struct {
	Name        string
	Temperature int // Desired temperature in Celsius
	FanSpeed    int // Fan speed (0-100%)
	Humidity    int
	Mode        string // HVAC mode (e.g., "Cool", "Heat", "Fan", "Off")
	Zone        int
}

// NewHVAC creates a new HVAC instance with the given name and initial settings.
func NewHVAC(name string) *HVAC {
	return &HVAC{
		Name:        name,
		Temperature: 25, // Initial temperature setting
		Humidity:    80,
		FanSpeed:    50,    // Initial fan speed setting (50%)
		Mode:        "Off", // Initial mode is Off
		Zone:        1,
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

func (h *HVAC) SetHumidity(humidity int) {
	if humidity < 0 {
		humidity = 0 // Ensure humidity is not set below 0%
	} else if humidity > 100 {
		humidity = 100 // Ensure humidity does not exceed 100%
	}
	h.Humidity = humidity
	fmt.Printf("%s humidity is set to %d%%\n", h.Name, h.Humidity)
}
