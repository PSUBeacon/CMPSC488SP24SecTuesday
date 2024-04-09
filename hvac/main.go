package hvac

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"fmt"
)

const temperaturePin = 4
const fanPin = uint8(18)

var mode string
var tempToSet int
var fanStatus string

// SetTemperature sets the desired temperature for the HVAC system.
func UpdateTemperature(newTemperature int) {
	currentTemp, err := gocode.ReadTemperature(temperaturePin, 22)
	if err != nil {
		fmt.Println("Error reading Temperature:", err)
		return
	}
	intCurrTemp := int(currentTemp)
	if newTemperature == intCurrTemp {
		gocode.FanStatus(fanPin, false)
		fmt.Printf("%s Temperature is set to %d°C\n", newTemperature)
	}
	if mode == "Cool" && newTemperature < intCurrTemp {
		gocode.FanStatus(fanPin, true)
		UpdateTemperature(newTemperature)
	}
	if mode == "Cool" && newTemperature > intCurrTemp {
		gocode.FanStatus(fanPin, false)
	}
	if mode == "Heat" && newTemperature > intCurrTemp {
		gocode.FanStatus(fanPin, true)
		UpdateTemperature(newTemperature)
	}
	if mode == "Heat" && newTemperature < intCurrTemp {
		gocode.FanStatus(fanPin, false)
	}

}

// SetFanSpeed sets the fan speed for the HVAC system.
func UpdateFanSpeed(speed int) {
	gocode.SetFanSpeed(fanPin, speed)
	fmt.Printf("%s fan speed is set to %s%%\n", speed)
}

// SetStatus sets the status (e.g., "Cool", "Heat", "Fan", "Off") for the HVAC system.
func UpdateStatus(status bool) {
	if status == true {
		gocode.FanStatus(fanPin, true)
		fmt.Printf("%s status is set to %s\n", status)
	}
	if status == false {
		gocode.FanStatus(fanPin, false)
		fmt.Printf("%s status is set to %s\n", status)
	}
}

func UpdateMode(mode string) {
	fmt.Printf("%s mode is set to %s\n", mode)
}

func DisplayLCDHVAC() {
	//currentTemp, err := gocode.ReadTemperature(temperaturePin, 22)
	//if err != nil {
	//	return
	//}
	//if currentTemp > 999 {
	//	currentTemp = 999
	//}
	//intCurrTemp := int(currentTemp)
	// Hardcode current temp, Mode, Temperature to set, Status
	intCurrTemp := 76
	mode = "Heat"
	tempToSet = 74
	fanStatus = "ON"

	gocode.WriteLCD("Now:" + fmt.Sprintf("%03d", intCurrTemp) + " Set:" + fmt.Sprintf("%02d", tempToSet) + "  Mode:" + mode + " Fan:" + fanStatus)
}

//func main() {
//	// Create a new HVAC instance
//	hvac := NewHVAC("My HVAC")
//
//	// Optional: Modify the HVAC instance as needed
//	hvac.SetTemperature(22)
//	hvac.SetFanSpeed("Medium")
//	hvac.SetStatus("Cool")
//	hvac.SetHumidity(45)
//
//	// Serialize the HVAC instance to JSON
//	jsonData, err := json.Marshal(hvac)
//	if err != nil {
//		fmt.Println("Error serializing HVAC to JSON:", err)
//		return
//	}
//
//	// Print the serialized JSON string
//	fmt.Println(string(jsonData))
//}
