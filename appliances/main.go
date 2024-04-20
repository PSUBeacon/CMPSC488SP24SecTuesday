package appliances

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"fmt"
)

// TurnOn turns the appliance on.
func UpdateStatus(appliance string, status bool) {

	kwHoursMap := map[string]int{
		"Oven":       4,
		"Microwave":  2,
		"Fridge":     2,
		"Toaster":    2,
		"Dishwasher": 4,
		"Solar":      15,
	}

	outStatus := ""
	if status == true {
		outStatus = "ON"
	}
	if status == false {
		outStatus = "OFF"
	}

	// Format the first and second line information
	firstLine := fmt.Sprintf("%-16s", appliance+":"+outStatus) // First line: Appliance and status

	// Ensure the appliance name and status do not exceed the first line limit
	if len(firstLine) > 16 {
		// Truncate and add ellipsis or reformat to fit within 16 characters
		firstLine = firstLine[:13] + "..."
	}

	// Format the second line for kWh
	secondLine := fmt.Sprintf("KWH: %d", kwHoursMap[appliance]) // Second line: kWh information
	if len(secondLine) > 16 {
		// Adjust second line to fit within 16 characters
		secondLine = fmt.Sprintf("KW:%d", kwHoursMap[appliance])
	}

	// Ensure second line does not exceed the limit
	if len(secondLine) > 16 {
		secondLine = secondLine[:16]
	}

	// Combine both lines into one string with a newline character to switch to the next line of the LCD
	fullLCDMessage := fmt.Sprintf("%-16s\n%-16s", firstLine, secondLine)

	// Write the combined string to the LCD in one call
	gocode.WriteLCD(fullLCDMessage)
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
