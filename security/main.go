package security

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"fmt"
)

var pirPin uint8 = 11
var keypadRowPin []int
var keypadColPin []int
var buzzerPin uint8 = 4

// MotionSensor represents a motion sensor component.

// DetectMotion simulates detecting motion.
func DetectMotion() {
	motion, err := gocode.CheckForMotion(pirPin)
	if err != nil {
		return
	}
	if motion == true {
		// If motion detected, turn buzzer on
		gocode.BuzzerStatus(buzzerPin, true)
		fmt.Printf("%s detected motion!\n")
	}

}

// Arm sets the alarm to the armed state.
func UpdateAlarmStatus(status bool) {
	if status == true {
		// Detect motion only when alarm is armed
		fmt.Printf("%s Alarm status is set to: \n", status)
		DetectMotion()
	} else {
		// A way to turn the buzzer off
		gocode.BuzzerStatus(buzzerPin, false)
	}

}

// Creates a lock or unlock feature
func LockOrUnlock(lock bool) {
	fmt.Println("Door is ", lock)
	gocode.TurnServo()
}

// creates Padlock
func NewPadlock(name string, pin string) {

}

// Padlock verify
func Verify(pin string) {

}
