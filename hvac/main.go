package hvac

import (
	messaging "CMPSC488SP24SecTuesday/AES-BlockChain-Communication"
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

const temperaturePin = 4
const fanPin = 12

type Thermostat struct {
	UUID        string `json:"UUID"`
	CurrentTemp int    `json:"currentTemp"`
	Mode        string `json:"mode"`
	FanStatus   string `json:"fanStatus"`
	Status      string `json:"status"`
	FanSpeed    int    `json:"fanSpeed"`
	SetTemp     int    `json:"setTemp"`
}

// SetTemperature sets the desired temperature for the HVAC system.
func UpdateTemperature(newTemperature int, uuid string) {
	jsonThermData, err := os.ReadFile("hvac/thermostats.json")
	if err != nil {
		fmt.Println("Error reading thermostat data:", err)
		return
	}
	//fmt.Println("Thermostat data, ", jsonThermData)
	// Unmarshal the JSON data into a Thermostat struct
	var thermostat []Thermostat
	if err := json.Unmarshal(jsonThermData, &thermostat); err != nil {
		fmt.Println("Error unmarshalling thermostat data:", err)
		return
	}
	index := 0
	if uuid == "050336" {
		index = 1
	}
	if uuid != "050337" {
		index = 0
	}

	thermostat[index].SetTemp = newTemperature

	//currentTemp, err := gocode.ReadTemperature(temperaturePin, 22)
	currentTemp := thermostat[index].CurrentTemp
	if err != nil {
		fmt.Println("Error reading Temperature:", err)
		return
	}
	thermostat[index].CurrentTemp = int(currentTemp)

	// currentTemp = 76
	intCurrTemp := int(currentTemp)
	if newTemperature == intCurrTemp {
		gocode.TurnOffFan(fanPin)
		thermostat[index].FanStatus = "OFF"
		fmt.Printf("%s Temperature is set to %d°C\n", newTemperature)
	}
	if thermostat[index].Mode == "Cool" && newTemperature < intCurrTemp {
		gocode.SetFanSpeed(fanPin, thermostat[index].FanSpeed)
		//UpdateTemperature(newTemperature)
	}
	if thermostat[index].Mode == "Cool" && newTemperature > intCurrTemp {
		thermostat[index].FanStatus = "OFF"
		gocode.TurnOffFan(fanPin)
	}
	if thermostat[index].Mode == "Heat" && newTemperature > intCurrTemp {
		gocode.SetFanSpeed(fanPin, thermostat[index].FanSpeed)
		thermostat[index].FanStatus = "ON"
		//UpdateTemperature(newTemperature)
	}
	if thermostat[index].Mode == "Heat" && newTemperature < intCurrTemp {
		gocode.TurnOffFan(fanPin)
		thermostat[index].FanStatus = "OFF"
	}

	DisplayLCDHVAC(thermostat[index].Mode, thermostat[index].SetTemp, thermostat[index].FanStatus)
	thermostatJSON, err := json.MarshalIndent(thermostat, "", "	")
	if err != nil {
		fmt.Println("Error marshalling thermostat data:", err)
		return
	}

	err = os.WriteFile("hvac/thermostats.json", thermostatJSON, 0644)
	if err != nil {
		panic(err)
	}

}

// SetFanSpeed sets the fan speed for the HVAC system.
func UpdateFanSpeed(speed int, uuid string) {
	jsonThermData, err := os.ReadFile("hvac/thermostats.json")
	if err != nil {
		fmt.Println("Error reading thermostat data:", err)
		return
	}
	//fmt.Println("Thermostat data, ", jsonThermData)
	// Unmarshal the JSON data into a Thermostat struct
	var thermostat []Thermostat
	if err := json.Unmarshal(jsonThermData, &thermostat); err != nil {
		fmt.Println("Error unmarshalling thermostat data:", err)
		return
	}
	index := 0
	if uuid == "050336" {
		index = 1
	}
	if uuid != "050337" {
		index = 0
	}
	thermostat[index].FanStatus = "ON"

	DisplayLCDHVAC(thermostat[index].Mode, thermostat[index].SetTemp, thermostat[index].FanStatus)
	gocode.SetFanSpeed(fanPin, speed)
	fmt.Printf("%s fan speed is set to %s%%\n", speed)

	thermostatJSON, err := json.MarshalIndent(thermostat, "", "	")
	if err != nil {
		fmt.Println("Error marshalling thermostat data:", err)
		return
	}

	err = os.WriteFile("hvac/thermostats.json", thermostatJSON, 0644)
	if err != nil {
		panic(err)
	}
}

// SetStatus sets the status (e.g., "CoUpdateStatusol", "Heat", "Fan", "Off") for the HVAC system.
func UpdateStatus(status bool, uuid string) {

	jsonThermData, err := os.ReadFile("hvac/thermostats.json")
	if err != nil {
		fmt.Println("Error reading thermostat data:", err)
		return
	}
	//fmt.Println("Thermostat data, ", jsonThermData)
	// Unmarshal the JSON data into a Thermostat struct
	var thermostat []Thermostat
	if err := json.Unmarshal(jsonThermData, &thermostat); err != nil {
		fmt.Println("Error unmarshalling thermostat data:", err)
		return
	}
	index := 0
	if uuid == "050336" {
		index = 1
	}
	if uuid != "050337" {
		index = 0
	}

	if status == true {

		thermostat[index].Status = "ON"
		fmt.Printf("%s status is set to %s\n", status)

	}
	if status == false {
		thermostat[index].FanStatus = "OFF"
		fmt.Printf("%s status is set to %s\n", status)
		thermostat[index].Status = "OFF"
		thermostat[index].Mode = "OFF"
	}
	DisplayLCDHVAC(thermostat[index].Mode, thermostat[index].SetTemp, thermostat[index].FanStatus)

	thermostatJSON, err := json.MarshalIndent(thermostat, "", "	")
	if err != nil {
		fmt.Println("Error marshalling thermostat data:", err)
		return
	}

	err = os.WriteFile("hvac/thermostats.json", thermostatJSON, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Thermostat data updated successfully")

	fmt.Printf("%s status is set to %s\n", status)
	//gocode.DrawH(9, 4, 10)

}

func UpdateMode(mode string, uuid string) {
	if mode == "cool" {
		mode = "Cool"
	}
	if mode == "heat" {
		mode = "Heat"
	}
	jsonThermData, err := os.ReadFile("hvac/thermostats.json")
	if err != nil {
		fmt.Println("Error reading thermostat data:", err)
		return
	}
	//fmt.Println("Thermostat data, ", jsonThermData)
	// Unmarshal the JSON data into a Thermostat struct
	var thermostat []Thermostat
	if err := json.Unmarshal(jsonThermData, &thermostat); err != nil {
		fmt.Println("Error unmarshalling thermostat data:", err)
		return
	}
	index := 0
	if uuid == "050336" {
		index = 1
	}
	if uuid != "050337" {
		index = 0
	}
	thermostat[index].Mode = mode

	fmt.Printf("%s mode is set to %s\n", mode)

	DisplayLCDHVAC(thermostat[index].Mode, thermostat[index].SetTemp, thermostat[index].FanStatus)

	if thermostat[index].Mode == "OFF" {
		thermostat[index].FanStatus = "OFF"
		thermostat[index].Status = "OFF"
		gocode.TurnOffFan(fanPin)
	}
	if thermostat[index].Mode == "Heat" || thermostat[index].Mode == "Cool" && thermostat[index].Status == "OFF" {
		thermostat[index].Status = "ON"
	}

	thermostatJSON, err := json.MarshalIndent(thermostat, "", "	")
	if err != nil {
		fmt.Println("Error marshalling thermostat data:", err)
		return
	}

	err = os.WriteFile("hvac/thermostats.json", thermostatJSON, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Thermostat data updated successfully")
}

type DefaultHVAC struct {
	Mode      string
	TempToSet int
	FanSpeed  int
	FanStatus string
}

func SendTempToFE() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		currentTemp, err := gocode.ReadTemperature(temperaturePin, 22)
		if err != nil {
			fmt.Println("Error reading temperature:", err)
			continue
		}

		// Convert the float value to a string
		tempStr := strconv.FormatFloat(currentTemp, 'f', -1, 64)

		// Convert the string to bytes
		tempBytes := []byte(tempStr)

		messaging.BroadCastMessage(tempBytes)
	}
}

func DisplayLCDHVAC(mode string, tempToSet int, fanStatus string) {

	defaults := DefaultHVAC{
		Mode:      "Heat",
		TempToSet: 74,
		FanSpeed:  50,
		FanStatus: "OFF",
	}

	if mode == "" {
		mode = defaults.Mode
	}
	if tempToSet == 0 {
		tempToSet = defaults.TempToSet
	}
	if fanStatus == "" {
		fanStatus = defaults.FanStatus
	}

	var intCurrTemp int
	rawCurrTemp, err := gocode.ReadTemperature(temperaturePin, 22)
	if err != nil {
		intCurrTemp = 0
	}

	intCurrTemp = int(rawCurrTemp)

	gocode.WriteLCD(fmt.Sprintf("%-16s", "Now:"+strconv.Itoa(intCurrTemp)+" Mode:"+mode) + fmt.Sprintf("%-16s", "Set:"+strconv.Itoa(tempToSet)+" Fan:"+fanStatus))

}

func StartupLCDHVAC() {

	jsonThermData, err := os.ReadFile("hvac/thermostats.json")
	if err != nil {
		fmt.Println("Error reading thermostat data:", err)
		return
	}
	//fmt.Println("Thermostat data, ", jsonThermData)
	// Unmarshal the JSON data into a Thermostat struct
	var thermostat []Thermostat
	if err := json.Unmarshal(jsonThermData, &thermostat); err != nil {
		fmt.Println("Error unmarshalling thermostat data:", err)
		return
	}

	gocode.WriteLCD(fmt.Sprintf("%-16s", "Now:"+strconv.Itoa(thermostat[0].CurrentTemp)+" Mode:"+thermostat[0].Mode) + fmt.Sprintf("%-16s", "Set:"+strconv.Itoa(thermostat[0].SetTemp)+" Fan:"+strconv.Itoa(thermostat[0].FanSpeed)))

}
