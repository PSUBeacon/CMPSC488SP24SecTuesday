package main

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// HVAC represents an HVAC system with temperature control, fan speed, and mode.
type HVAC struct {
	ID                primitive.ObjectID
	UUID              string
	Location          string
	Temperature       string // Desired temperature in Celsius
	Humidity          string
	FanSpeed          string // Fan speed (0-100%)
	Status            bool   // System status (on/off)
	EnergyConsumption int
	LastChanged       time.Time
}

// NewHVAC creates a new HVAC instance with the given location and initial settings.
func NewHVAC(location string) *HVAC {
	return &HVAC{
		ID:                primitive.NewObjectID(),
		UUID:              "", // UUID should be generated or assigned elsewhere if necessary
		Location:          location,
		Temperature:       "25",       // Initial temperature setting in Celsius
		Humidity:          "50",       // Initial humidity setting
		FanSpeed:          "50%",      // Initial fan speed setting
		Status:            false,      // Initial system status is off
		EnergyConsumption: 10,         // Initial energy consumption value
		LastChanged:       time.Now(), // Use the current time for the last changed value
	}
}

func (h *HVAC) SetHumidity(humidity string) {
	h.Humidity = humidity
	h.LastChanged = time.Now()
	fmt.Printf(" Humidity is set to %s\n", h.Humidity)
}

// SetTemperature sets the desired temperature for the HVAC system.
func (h *HVAC) SetTemperature(temp string) {
	h.Temperature = temp
	h.LastChanged = time.Now()
	fmt.Printf(" Temperature is set to %s\n", h.Temperature)
}

// SetFanSpeed updates the HVAC's fan speed.
func (h *HVAC) SetFanSpeed(speed string) {
	h.FanSpeed = speed
	h.LastChanged = time.Now()
}

// TurnOn turns the HVAC system on.
func (h *HVAC) TurnOn() {
	h.Status = true
	h.LastChanged = time.Now()
}

// TurnOff turns the HVAC system off.
func (h *HVAC) TurnOff() {
	h.Status = false
	h.LastChanged = time.Now()
}
func (h *HVAC) SetStatus(status bool) {
	h.Status = status
	h.LastChanged = time.Now()
	fmt.Printf(" Status is set to %t\n", h.Status)
}
func main() {
	// Create a new HVAC system
	hvac := NewHVAC("Living Room")
	fmt.Println("Initial HVAC system:", hvac)

	// Change settings
	hvac.SetTemperature("22")
	hvac.SetFanSpeed("75%")
	hvac.TurnOn()

	fmt.Println("Updated HVAC system:", hvac)
	// Serialize the HVAC instance to JSON
	jsonData, err := json.Marshal(hvac)
	if err != nil {
		fmt.Println("Error serializing HVAC to JSON:", err)
		return
	}

	// Print the serialized JSON strin
	fmt.Println(string(jsonData))
}
