package security

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"fmt"
)

// MotionSensor represents a motion sensor component.

// DetectMotion simulates detecting motion.
func DetectMotion() {

	fmt.Printf("%s detected motion!\n")
}

// Arm sets the alarm to the armed state.
func UpdateAlarmStatus(status bool) {
	fmt.Printf("%s Alarm status is set to: \n", status)
}

// Creates a lock or unlock feature
func LockOrUnlock(lock bool) {
	fmt.Println("Door is ", lock)
	gocode.TurnServoTo90Degrees()
}

// creates Padlock
func NewPadlock(name string, pin string) {

}

// Padlock verify
func Verify(pin string) {

}
