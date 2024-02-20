package main

import (
	"CMPSC488SP24SecTuesday/appliances"
	"fmt"
)

// Git not working big annoying
func main() {
	// Create a new light switch appliance
	//lightSwitch := appliances.NewAppliance("Light Switch", false, 0)
	ReceivedVal := appliances.ReceiverController()
	fmt.Println("Value is: ", ReceivedVal)
	//if ReceivedVal == "lights off" {
	//	lightSwitch.TurnOff()
	//}
	//
	//if ReceivedVal == "lights on" {
	//	lightSwitch.TurnOn()
	//}

	// Use the appliance
	oven := appliances.NewAppliance("Oven", false, 0)

	oven.AdjustTemp(300)
	oven.TurnOn()
}
