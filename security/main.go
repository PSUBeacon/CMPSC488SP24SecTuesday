package security

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"encoding/json"
	"fmt"
	"os"
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

//
//func DetectMotion() {
//	var alarm dal.MessagingStruct
//	motion := gocode.CheckForMotion()
//	if motion {
//		alarm.UUID = "0"
//		alarm.Name = "security"
//		alarm.AppType = "security"
//		alarm.Function = "alarm"
//		alarm.Change = "alarm"
//
//		soundAlarm, err := json.MarshalIndent(alarm, "", "  ")
//		if err != nil {
//			panic(err)
//		}
//
//		messaging.BroadCastMessage(soundAlarm)
//		fmt.Println("Motion detected!")
//	}
//}

// get keypad input
func GetKeypadInput() {
	gocode.InitKeypad()

}

func UpdateAlarmStatus(armed bool) {

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

	if armed {
		fmt.Println("Alarm armed.")

		securitySystem.Status = "Armed"
		securitySystem.SensorStatus = "ON"
		DisplayLCDSecurity(securitySystem.Status, securitySystem.SensorStatus)

		//go DetectMotion()
	} else {
		//gocode.BuzzerStatus()
		fmt.Println("Alarm disarmed.")

		securitySystem.Status = "Disarmed"
		securitySystem.SensorStatus = "OFF"

		DisplayLCDSecurity(securitySystem.Status, securitySystem.SensorStatus)
	}

	securityJSON, err := json.MarshalIndent(securitySystem, "", "	")
	if err != nil {
		fmt.Println("Error marshalling security data:", err)
		return
	}

	err = os.WriteFile("security/Keys.json", securityJSON, 0644)
	if err != nil {
		panic(err)
	}

}

func LockOrUnlock(lock bool) {
	fmt.Println("Door is locked:", lock)
	// Add logic to control door lock here
}

//func HandleMotionDetection() {
//	ticker := time.NewTicker(5 * time.Second)
//	defer ticker.Stop()
//
//	for range ticker.C {
//		motion = !motion
//		if motion {
//			UpdateAlarmStatus(true)
//		} else {
//			UpdateAlarmStatus(false)
//		}
//	}
//}

type System struct {
	UUID         string `json:"UUID"`
	Status       string `json:"Status"`
	SensorStatus string `json:"SensorStatus"`
}

func DisplayLCDSecurity(status string, motionStatus string) {
	gocode.WriteLCD(fmt.Sprintf("%-16s", "Stat:"+status) + fmt.Sprintf("%-16s", "Motion:"+motionStatus))
}
