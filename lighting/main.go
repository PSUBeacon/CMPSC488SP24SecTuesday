package lighting

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	dinPinNumber = 9  // GPIO pin for DIN (MOSI)
	csPinNumber  = 4  // GPIO pin for CS
	clkPinNumber = 10 // GPIO pin for CLK
)

type light struct {
	Status     string
	Brightness int
}

// TurnOn turns the lighting on.
func UpdateStatus(newStatus bool) {

	jsonlightData, err := os.ReadFile("lighting/lighting.json")
	if err != nil {
		fmt.Println("Error reading key data:", err)
		return
	}
	//fmt.Println("Thermostat data, ", jsonThermData)
	// Unmarshal the JSON data into a Thermostat struct
	var bright light
	if err := json.Unmarshal(jsonlightData, &bright); err != nil {
		fmt.Println("Error unmarshalling lighting data:", err)
		return
	}
	if newStatus {
		bright.Status = "On"
	} else {
		bright.Status = "Off"
	}

	lightJSON, err := json.MarshalIndent(bright, "", "	")
	if err != nil {
		fmt.Println("Error marshalling lighting data:", err)
		return
	}

	err = os.WriteFile("lighting/lighting.json", lightJSON, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s is now turned \n", newStatus)
	gocode.MatrixStatus(9, 4, 10, newStatus, bright.Brightness)

}

// SetBrightness sets the brightness of the lighting.
func SetBrightness(brightness int) {
	if brightness < 0 {
		brightness = 0
	} else if brightness > 15 {
		brightness = 15
	}
	jsonlightData, err := os.ReadFile("lighting/lighting.json")
	if err != nil {
		fmt.Println("Error reading key data:", err)
		return
	}
	//fmt.Println("Thermostat data, ", jsonThermData)
	// Unmarshal the JSON data into a Thermostat struct
	var bright light
	if err := json.Unmarshal(jsonlightData, &bright); err != nil {
		fmt.Println("Error unmarshalling security data:", err)
		return
	}
	bright.Brightness = brightness

	gocode.DrawLightbulb(9, 4, 10, bright.Brightness)
	time.Sleep(3 * time.Second)
	gocode.TurnOffMatrix(9, 4, 10)
	gocode.TurnOnMatrix(9, 4, 10)

	gocode.SetIntensity(9, 4, 10, bright.Brightness)
	fmt.Printf("%s brightness is set to %s\n", bright.Brightness)

	lightJSON, err := json.MarshalIndent(bright, "", "	")
	if err != nil {
		fmt.Println("Error marshalling lighting data:", err)
		return
	}

	err = os.WriteFile("lighting/lighting.json", lightJSON, 0644)
	if err != nil {
		panic(err)
	}

}

func FlashSymbol(Symbol string) {
	jsonlightData, err := os.ReadFile("lighting/lighting.json")
	if err != nil {
		fmt.Println("Error reading key data:", err)
		return
	}
	//fmt.Println("Thermostat data, ", jsonThermData)
	// Unmarshal the JSON data into a Thermostat struct
	var bright light
	if err := json.Unmarshal(jsonlightData, &bright); err != nil {
		fmt.Println("Error unmarshalling security data:", err)
		return
	}
	if Symbol == "Bulb" {
		gocode.DrawLightbulb(9, 4, 10, bright.Brightness)
	}
	if Symbol == "HVAC" {
		gocode.DrawH(9, 4, 10, bright.Brightness)
	}
	if Symbol == "Lock" {
		gocode.DrawLock(9, 4, 10, bright.Brightness)
	}
	if Symbol == "App" {
		gocode.DrawA(9, 4, 10, bright.Brightness)
	}
	time.Sleep(3 * time.Second)
	gocode.TurnOffMatrix(9, 4, 10)

	if bright.Status == "On" {
		gocode.TurnOnMatrix(9, 4, 10)

		gocode.SetIntensity(9, 4, 10, bright.Brightness)
		fmt.Printf("%s brightness is set to %s\n", bright.Brightness)
	}

	lightJSON, err := json.MarshalIndent(bright, "", "	")
	if err != nil {
		fmt.Println("Error marshalling lighting data:", err)
		return
	}

	err = os.WriteFile("lighting/lighting.json", lightJSON, 0644)
	if err != nil {
		panic(err)
	}
}

//func drawBulblTimer(dinPin, csPin, clkPin rpio.Pin) {
//	gocode.DrawLightbulb(dinPin, csPin, clkPin)
//	time.Sleep(3 * time.Second)
//	gocode.ClearMatrix(dinPin, csPin, clkPin)
//
//}
