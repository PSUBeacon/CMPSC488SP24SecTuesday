package main

import (
	"encoding/json"
	"fmt"
)

// HVAC represents an HVAC system with temperature control, fan speed, and mode.
type HVAC struct {
	Name              string
	Temperature       int    // Desired temperature in Celsius
	FanSpeed          string // Fan speed ("Off" , "Low", "Medium", "High")
	Humidity          int    // Humidity %
	Status            string // HVAC mode (e.g., "Cool", "Heat", "Fan", "Off")
	Location          string // Location of device
	EnergyConsumption int
	LastChanged       string
}

// NewHVAC creates a new HVAC instance with the given name and initial settings.
func NewHVAC(name string) *HVAC {
	return &HVAC{
		Name:              name,
		Temperature:       25, // Initial temperature setting
		Humidity:          80,
		FanSpeed:          "Low", // Initial fan speed setting (50%)
		Status:            "Off", // Initial mode is Off
		Location:          "Kitchen",
		EnergyConsumption: 10,
		LastChanged:       "2023-10-01T18:30:00Z",
	}
}

// SetTemperature sets the desired temperature for the HVAC system.
func (h *HVAC) SetTemperature(temperature int) {
	h.Temperature = temperature
	fmt.Printf("%s temperature is set to %d°C\n", h.Name, h.Temperature)
}

// SetFanSpeed sets the fan speed for the HVAC system.
func (h *HVAC) SetFanSpeed(speed string) {
	h.FanSpeed = speed
	fmt.Printf("%s fan speed is set to %s%%\n", h.Name, h.FanSpeed)
}

// SetStatus sets the status (e.g., "Cool", "Heat", "Fan", "Off") for the HVAC system.
func (h *HVAC) SetStatus(status string) {
	h.Status = status
	fmt.Printf("%s mode is set to %s\n", h.Name, h.Status)
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
func main() {
	// Create a new HVAC instance
	hvac := NewHVAC("My HVAC")

	// Optional: Modify the HVAC instance as needed
	hvac.SetTemperature(22)
	hvac.SetFanSpeed("Medium")
	hvac.SetStatus("Cool")
	hvac.SetHumidity(45)

	// Serialize the HVAC instance to JSON
	jsonData, err := json.Marshal(hvac)
	if err != nil {
		fmt.Println("Error serializing HVAC to JSON:", err)
		return
	}

	// Print the serialized JSON string
	fmt.Println(string(jsonData))
}
