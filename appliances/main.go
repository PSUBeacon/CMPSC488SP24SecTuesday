package appliances

import (
	"fmt"
)

// TurnOn turns the appliance on.
func UpdateStatus(status bool) {
	fmt.Printf("%s Status is now set to: \n", status)
}

func UpdateTemperature(temp int) {
	fmt.Printf("%s Temperature is now set to %d degrees farenheit\n", temp)
	// Add function to get the temp

}
 
// TurnOn turns the appliance on.
func UpdateStatus(status bool) {
	fmt.Printf("%s Status is now set to: \n", status)
}

func UpdateTemperature(temp int) {
	//Add microwave
	fmt.Printf("%s Temperature is now set to %d degrees farenheit\n", temp)
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
