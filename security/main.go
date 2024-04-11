package security

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	pirPin    = 11
	buzzerPin = 4
)

var (
	motion       bool
	keypadRowPin []int
	keypadColPin []int
)

type DefaultSecurity struct {
	Status       string
	MotionStatus string
}

func DetectMotion() {
	motion, err := gocode.CheckForMotion(pirPin)
	if err != nil {
		fmt.Println("Error detecting motion:", err)
		return
	}
	if motion {
		gocode.BuzzerStatus(buzzerPin, true)
		fmt.Println("Motion detected!")
	}
}

// get keypad input
func GetKeypadInput() {
	gocode.InitKeypad()

}

func UpdateAlarmStatus(armed bool) {
	if armed {
		fmt.Println("Alarm armed.")
		DetectMotion()
	} else {
		gocode.BuzzerStatus(buzzerPin, false)
		fmt.Println("Alarm disarmed.")
	}
}

func LockOrUnlock(lock bool) {
	fmt.Println("Door is locked:", lock)
	// Add logic to control door lock here
}

func HandleMotionDetection() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		motion = !motion
		if motion {
			UpdateAlarmStatus(true)
		} else {
			UpdateAlarmStatus(false)
		}
	}
}

type System struct {
	UUID         string `json:"UUID"`
	Status       string `json:"Status"`
	SensorStatus string `json:"SensorStatus"`
}

func DisplayLCDSecurity(status string, motionStatus string) {
	jsonSecurityData, err := os.ReadFile("security/Keys.json")
	if err != nil {
		fmt.Println("Error reading key data:", err)
		return
	}
	//fmt.Println("Thermostat data, ", jsonThermData)
	// Unmarshal the JSON data into a Thermostat struct
	var securitySystem System
	if err := json.Unmarshal(jsonSecurityData, &securitySystem); err != nil {
		fmt.Println("Error unmarshalling security data:", err)
		return
	}

	gocode.WriteLCD("Stat:" + securitySystem.Status + " Motion:" + securitySystem.SensorStatus)
}
