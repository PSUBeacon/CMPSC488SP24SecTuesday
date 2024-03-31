package hvac

import (
	"fmt"
)

// SetTemperature sets the desired temperature for the HVAC system.
func UpdateTemperature(temperature int) {
	fmt.Printf("%s temperature is set to %d°C\n", temperature)
}

// SetFanSpeed sets the fan speed for the HVAC system.
func UpdateFanSpeed(speed int) {
	fmt.Printf("%s fan speed is set to %s%%\n", speed)
}

// SetStatus sets the status (e.g., "Cool", "Heat", "Fan", "Off") for the HVAC system.
func UpdateStatus(status bool) {
	fmt.Printf("%s status is set to %s\n", status)
}

func UpdateMode(mode string) {
	fmt.Printf("%s mode is set to %s\n", mode)
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
