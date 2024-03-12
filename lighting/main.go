package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Lighting represents the state of a smart light in the system.
type Lighting struct {
	UUID              string `json:"uuid"`              // Unique identifier for the light
	Location          string `json:"location"`          // Location of the light (e.g., "Living Room", "Kitchen").
	Brightness        string `json:"brightness"`        // Brightness level as a string to include numeric value.
	Status            bool   `json:"status"`            // true for on, false for off.
	EnergyConsumption int    `json:"energyConsumption"` // Energy consumption in kilowatts.
	LastChanged       string `json:"lastChanged"`       // Timestamp of the last change in ISODate format.
}

// NewLighting creates a new Lighting instance with the given parameters.
func NewLighting(uuid, location string, status bool, brightness string, energyConsumption int) *Lighting {
	return &Lighting{
		UUID:              uuid,
		Location:          location,
		Status:            status,
		Brightness:        brightness,
		EnergyConsumption: energyConsumption,
		LastChanged:       time.Now().Format(time.RFC3339), // Capture the creation time in ISO 8601 format
	}
}

// serializeLight converts a Lighting object to a JSON string.
func serializeLight(light *Lighting) (string, error) {
	lightJSON, err := json.MarshalIndent(light, "", " ")
	if err != nil {
		return "", err
	}
	return string(lightJSON), nil
}

// deserializeLight converts a JSON string back into a Lighting object.
func deserializeLight(lightJSON string) (*Lighting, error) {
	var light Lighting
	err := json.Unmarshal([]byte(lightJSON), &light)
	if err != nil {
		return nil, err
	}
	return &light, nil
}

func main() {
	// Example usage:
	lightUUID := "123e4567-e89b-12d3-a456-426614174000" // Example UUID
	light := NewLighting(lightUUID, "Living Room", false, "50", 10)

	// Serialize the Lighting object
	lightJSON, err := serializeLight(light)
	if err != nil {
		fmt.Println("Error serializing light:", err)
		return
	}
	fmt.Println("Serialized light:", lightJSON)

	// Deserialize the JSON back into a Lighting object
	deserializedLight, err := deserializeLight(lightJSON)
	if err != nil {
		fmt.Println("Error deserializing light:", err)
		return
	}
	fmt.Printf("Deserialized light: %+v\n", deserializedLight)
}
