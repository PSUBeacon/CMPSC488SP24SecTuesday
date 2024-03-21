package main

import (
	"encoding/json"
	"fmt"
)

// HVAC represents an HVAC system with temperature control, fan speed, and mode.
type HVAC struct {
	Name              string `json:"Name"`
	Temperature       int    `json:"Temperature"`       // Desired temperature in Celsius
	FanSpeed          string `json:"FanSpeed"`          // Fan speed ("Off" , "Low", "Medium", "High")
	Humidity          int    `json:"Humidity"`          // Humidity %
	Status            string `json:"Status"`            // HVAC mode (e.g., "Cool", "Heat", "Fan", "Off")
	Location          string `json:"Location"`          // Location of device
	EnergyConsumption int    `json:"EnergyConsumption"` // Energy Consumption
	LastChanged       string `json:"LastChanged"`       // Last Changed (date)
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

func serializeHvac(hvac *HVAC) (string, error) {
	hvacJSON, err := json.MarshalIndent(hvac, "", " ")
	if err != nil {
		return "", err
	}
	return string(hvacJSON), nil
}

func deserializeHvac(hvacJSON string) (*HVAC, error) {
	var hvac HVAC
	err := json.Unmarshal([]byte(hvacJSON), &hvac)
	if err != nil {
		return nil, err
	}
	return &hvac, nil
}

func main() {
	// Create a new HVAC instance
	hvac := NewHVAC("My HVAC")

	// Optional: Modify the HVAC instance as needed
	hvac.SetTemperature(22)
	hvac.SetFanSpeed("Medium")
	hvac.SetStatus("Cool")
	hvac.SetHumidity(45)

	// Serialize the Hvac instance to JSON
	hvacJSON, err := serializeHvac(hvac)
	if err != nil {
		fmt.Println("Error serializing HVAC to JSON:", err)
		return
	}

	// Print the serialized JSON string
	fmt.Println(string(hvacJSON))

	//Deserialize
	deserializeHvac, err := deserializeHvac(hvacJSON)
	if err != nil {
		fmt.Println("Error deserializing HVAC:", err)
		return
	}

	fmt.Printf("\nDeserialized HVAC:\n %+v\n", deserializeHvac)

}
