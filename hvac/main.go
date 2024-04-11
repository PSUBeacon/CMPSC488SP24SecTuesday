package hvac

import (
	messaging "CMPSC488SP24SecTuesday/AES-BlockChain-Communication"
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"fmt"
	"strconv"
	"time"
)

const temperaturePin = 4
const fanPin = 12

var mode string
var tempToSet = 0
var fanSpeed int = 50
var fanStatus string

// SetTemperature sets the desired temperature for the HVAC system.
func UpdateTemperature(newTemperature int) {
	currentTemp, err := gocode.ReadTemperature(temperaturePin, 22)
	if err != nil {
		fmt.Println("Error reading Temperature:", err)
		return
	}
	// currentTemp = 76
	intCurrTemp := int(currentTemp)
	if newTemperature == intCurrTemp {
		gocode.TurnOffFan(fanPin)
		fmt.Printf("%s Temperature is set to %d°C\n", newTemperature)
	}
	if mode == "Cool" && newTemperature < intCurrTemp {
		gocode.SetFanSpeed(fanPin, fanSpeed)
		//UpdateTemperature(newTemperature)
	}
	if mode == "Cool" && newTemperature > intCurrTemp {
		gocode.TurnOffFan(fanPin)
	}
	if mode == "Heat" && newTemperature > intCurrTemp {
		gocode.SetFanSpeed(fanPin, fanSpeed)
		//UpdateTemperature(newTemperature)
	}
	if mode == "Heat" && newTemperature < intCurrTemp {
		gocode.TurnOffFan(fanPin)
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
		gocode.SetFanSpeed(fanPin, fanSpeed)
		fmt.Printf("%s status is set to %s\n", status)
	}
	if status == false {
		gocode.TurnOffFan(fanPin)
		fmt.Printf("%s status is set to %s\n", status)
	}
}

func UpdateMode(mode string) {
	fmt.Printf("%s mode is set to %s\n", mode)
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
	//currentTemp, err := gocode.ReadTemperature(temperaturePin, 22)
	//if err != nil {
	//	return
	//}
	//if currentTemp > 999 {
	//	currentTemp = 999
	//}
	//intCurrTemp := int(currentTemp)
	// Hardcode current temp, Mode, Temperature to set, Status
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

	intCurrTemp := 76

	// gocode.WriteLCD("Now:" + fmt.Sprintf("%02d", intCurrTemp) + " Set:" + fmt.Sprintf("%02d", tempToSet) + "   Mode:" + mode + " Fan:" + fanStatus)
	gocode.WriteLCD("Now:" + fmt.Sprintf("%02d", intCurrTemp) + " Mode:" + mode + "Set:" + fmt.Sprintf("%02d", tempToSet) + " Fan:" + fanStatus)

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
