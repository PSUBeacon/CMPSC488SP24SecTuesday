package security

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"fmt"
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

func DisplayLCDSecurity(status string, motionStatus string) {
	defaults := DefaultSecurity{
		Status:       "Armed",
		MotionStatus: "off",
	}

	if status == "" {
		status = defaults.Status
	}
	if motionStatus == "" {
		motionStatus = defaults.MotionStatus
	}
	key := gocode.InitKeypad()

	gocode.WriteLCD("Stat:" + status + " Motion:" + motionStatus + "Key: " + string(key))
}
