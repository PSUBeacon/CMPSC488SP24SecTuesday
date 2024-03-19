package main

import (
	"CMPSC488SP24SecTuesday/appliances"
	"fmt"
)

// Git not working big annoying
func main() {
	// Create a new light switch appliance
	//lightSwitch :=
	appliances.NewAppliance("Light Switch", false, 0)
	oven := appliances.NewAppliance("Oven", false, 0)
	ReceivedVal := appliances.ReceiverController()
	fmt.Println("Value is: ", ReceivedVal)
	if ReceivedVal == "oven off" {
		oven.TurnOff()
	}
	if ReceivedVal == "oven on" {
		oven.TurnOn()
	}

	if ReceivedVal == "temp change" {
		oven.AdjustTemp(300)
	}

	// Use the appliance
}
