package appliances

import (
	"encoding/json"
	"fmt"
)

// Appliance represents a simple appliance with an on/off state.
type Appliance struct {
	Name  string `json:"Name"`
	State bool   `json:"State"`
	Temp  int    `json:"Temp"`
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
func serializeAppliance(appliance *Appliance) (string, error) {
	applianceJSON, err := json.MarshalIndent(appliance, "", " ")
	if err != nil {
		return "", err
	}
	return string(applianceJSON), nil
}

func deserializeAppliance(applianceJSON string) (*Appliance, error) {
	var appliance Appliance
	err := json.Unmarshal([]byte(applianceJSON), &appliance)
	if err != nil {
		return nil, err
	}
	return &appliance, nil
}

func main() {
	myAppliance := NewAppliance("Toaster", false, 85)

	//on
	myAppliance.TurnOn()

	//off
	myAppliance.TurnOff()

	//set temp
	myAppliance.AdjustTemp(100)

	// Serialize the Appliance instance to JSON
	applianceJSON, err := serializeAppliance(myAppliance)
	if err != nil {
		fmt.Println("Error serializing APPLIANCES to JSON:", err)
		return
	}

	// Print the serialized JSON string
	fmt.Println("\nSerialized appliance: \n" + applianceJSON)

	// Deserialize the JSON back into a Appliance object
	deserializeAppliance, err := deserializeAppliance(applianceJSON)
	if err != nil {
		fmt.Println("Error deserializing appliance:", err)
		return
	}

	fmt.Printf("\nDeserialized Appliance:\n %+v\n", deserializeAppliance)
}
