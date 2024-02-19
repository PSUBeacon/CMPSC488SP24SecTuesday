package main

import (
	"CMPSC488SP24SecTuesday/appliances"
	"fmt"
)

func main() {
	// Create a new light switch appliance
	lightSwitch := appliances.NewAppliance("Light Switch", false, 0)
	fmt.Println(appliances.ReceiverController())
	// Use the appliance
	lightSwitch.TurnOn()
	lightSwitch.TurnOff()

	oven := appliances.NewAppliance("Oven", false, 0)

	oven.AdjustTemp(300)
	oven.TurnOn()
}
