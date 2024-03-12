package main

import (
	"fmt"
	"CMPSC488SP24SecTuesday/security"
)

// Git not working big annoying
// I agree @Cameron
func main() {
	// Creates alarm, padlock, front door, back door

	securityAlarm := NewAlarm("Security Alarm")
	padLock1 := NewPadlock("PadLock1", "1234")
	frontdoor := NewDoorLock("FrontDoor", true)
	backdoor := NewDoorLock("BackDoor", true)
	ReceivedVal := appliances.ReceiverController()
	fmt.Println("Value is: ", ReceivedVal)

	if ReceivedVal == "lock front door" {
		frontdoor.LockOrUnlock(true)
	}

	if ReceivedVal == "unlock front door" {
		frontdoor.LockOrUnlock(false)
	}

	if ReceivedVal == "unlock back door" {
		backdoor.LockOrUnlock(false)
	}

	if ReceivedVal == "lock back door" {
		backdoor.LockOrUnlock(true)
	}

	if ReceivedVal == "lock back door" {
		backdoor.LockOrUnlock(true)
	}

	if ReceivedVal == "4321" {

		padLock1.Verify("4321")
	}

	if ReceivedVal == "arm" {

		securityAlarm.Arm()
	}
	if ReceivedVal == "disarm" {

		securityAlarm.Disarm()
	}

	if ReceivedVal == "trigger" {

		securityAlarm.Trigger()
	}

	if ReceivedVal == "detect" {

		motionSensor.DetectMotion()
	}

	if ReceivedVal == "detect" {

		motionSensor.DetectMotion()
	}

	// Use security modules
}
