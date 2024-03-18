package energy

import (
	"encoding/json"
	"fmt"
	"time"
)

type Device struct {
	DeviceID int    `json:"DeviceID"` // 0 = Solar, 1 = Battery, 2 = Security Systems, 4 = Lighting, 5 = Other
	Name     string `json:"Name"`
	Power    int    `json:"Power"`
	State    bool   `json:"State"`
	Sector   int    `json:"Sector"`
}

func newDevice(deviceID int, name string, power int, state bool, sector int) *Device {
	return &Device{
		DeviceID: deviceID,
		Name:     name,
		Power:    power,
		State:    state,
		Sector:   sector,
	}
}

// SolarPanel represents a solar panel as a power source.
type SolarPanel struct {
	Name        string
	PowerOutput int // Power output in watts
}

// NewSolarPanel creates a new SolarPanel instance with the given name and power output.
func NewSolarPanel(name string, powerOutput int) *SolarPanel {
	return &SolarPanel{
		Name:        name,
		PowerOutput: powerOutput,
	}
}

// Battery represents an energy storage battery.
type Battery struct {
	Name     string
	Capacity int // Battery capacity in watt-hours (Wh)
	Charge   int // Current charge level in watt-hours (Wh)
}

// NewBattery creates a new Battery instance with the given name and capacity.
func NewBattery(name string, capacity int) *Battery {
	return &Battery{
		Name:     name,
		Capacity: capacity,
		Charge:   0, // Initialize with 0 charge
	}
}

// Appliance represents an appliance with energy consumption.
type Appliance struct {
	Name        string
	PowerRating int  // Power rating of the appliance in watts
	IsOn        bool // Whether the appliance is turned on
}

// NewAppliance creates a new Appliance instance with the given name and power rating.
func NewAppliance(name string, powerRating int) *Appliance {
	return &Appliance{
		Name:        name,
		PowerRating: powerRating,
		IsOn:        false,
	}
}

// TurnOn turns the appliance on.
func (a *Appliance) TurnOn() {
	a.IsOn = true
	fmt.Printf("%s is turned ON\n", a.Name)
}

// TurnOff turns the appliance off.
func (a *Appliance) TurnOff() {
	a.IsOn = false
	fmt.Printf("%s is turned OFF\n", a.Name)
}

func serializeDevice(device *Device) (string, error) {
	deviceJSON, err := json.MarshalIndent(device, "", " ")
	if err != nil {
		return "", err
	}
	return string(deviceJSON), nil
}

// deserializeDevice deserializes a JSON string into a Device instance.
func deserializeDevice(deviceJSON string) (*Device, error) {
	var device Device
	err := json.Unmarshal([]byte(deviceJSON), &device)
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func monitor() {
	solarPanel := newDevice(0, "Solar 1", 1500, false, -1)
	houseBattery := newDevice(1, "Battery 1", 3000, true, 0)
	solarEnergy := solarPanel.Power
	houseBattery.Power += solarEnergy
}

// !FOR DEMO ONLY NOT FOR FINAL PROD
func demoMonitoring() {
	// Create a solar panel, battery, and appliances
	solarPanel := NewSolarPanel("Solar Panel", 500)         // 500 watts of power output
	houseBattery := NewBattery("House Battery", 2000)       // 2000 watt-hours capacity
	fridge := NewAppliance("Fridge", 200)                   // 200 watts
	airConditioner := NewAppliance("Air Conditioner", 1500) // 1500 watts

	// TODO: showing gain and subtract per 10 sec
	i := 1
	for i <= 6 {

		time.Sleep(10000)
		i += 1
		// Simulate powering the appliances with solar energy
		solarEnergy := solarPanel.PowerOutput
		houseBattery.Charge += solarEnergy

		// Turn on appliances
		fridge.TurnOn()
		airConditioner.TurnOn()

		// Simulate appliance power consumption
		if fridge.IsOn {
			houseBattery.Charge -= fridge.PowerRating
		}
		if airConditioner.IsOn {
			houseBattery.Charge -= airConditioner.PowerRating
		}

		// Check battery charge level
		fmt.Printf("House Battery Charge Level: %d Wh\n", houseBattery.Charge)
	}

}

func main() {
	myDevice := newDevice(0, "Solar Panel", 1500, true, 1)

	// Serialize the Device instance to a JSON string
	deviceJSON, err := serializeDevice(myDevice)
	if err != nil {
		fmt.Println("Error serializing DEVICE to JSON:", err)
		return
	}

	// Print the serialized JSON string
	fmt.Println("\nSerialized device: \n" + deviceJSON)

	// Deserialize the JSON back into a Device object
	deserializedDevice, err := deserializeDevice(deviceJSON)
	if err != nil {
		fmt.Println("Error deserializing device:", err)
		return
	}

	fmt.Printf("\nDeserialized Device:\n %+v\n", deserializedDevice)

	// *This is using Goroutines, multithreading.
	// Though it isn't being put to actually use yet it will be useful for monitoring multiple devices.
	go demoMonitoring()
	demoMonitoring()
}
