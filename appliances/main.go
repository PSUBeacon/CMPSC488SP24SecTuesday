package appliances

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"fmt"
)

// TurnOn turns the appliance on.
func UpdateStatus(appliance string, status bool) {
	outStatus := ""
	if status == true {
		outStatus = "ON"
	}
	if status == false {
		outStatus = "OFF"
	}
	toLCD := appliance + ":" + outStatus
	gocode.WriteLCD(toLCD)
	fmt.Printf("%s Status is now set to: \n", status)
}

func UpdateTemperature(temp int) {
	fmt.Printf("%s Temperature is now set to %d degrees farenheit\n", temp)
	// Add function to get the temp

}

func UpdateTimeStopTime(timerstop int) {
	fmt.Printf("%s Timer end time is now set to\n", timerstop)
}

func UpdatePower(power int) {
	fmt.Printf("%s Power is now set to\n", power)
}

func UpdateEnergySavingMode(status bool) {
	fmt.Printf("%s Energy saving mode is now set to \n", status)
}

func UpdateWashTime(washtime int) {
	fmt.Printf("%s Wash time is now set to \n", washtime)
}
