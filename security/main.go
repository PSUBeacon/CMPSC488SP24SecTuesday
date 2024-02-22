package main

import (
	"fmt"
)

// Here is the Security System Structure
type SecuritySystem struct {
	Location          string `json:"Location"`   //(mcq)
	SensorType        string `json:"SensorType"` //(mcq)
	Status            bool   `json:"Status"`
	EnergyConsumption int    `json:"EnergyConsumption"` //(kilowatts)
	LastTriggered     string `json:"LastTriggered"`     //(this will be the main notification)
}

// MotionSensor represents a motion sensor component.

type MotionSensor struct {
	Name           string `json:"Name"`
	MotionDetected bool   `json:"MotionDetected"`
}

// NewMotionSensor creates a new MotionSensor instance with the given name.
func NewMotionSensor(name string) *MotionSensor {
	return &MotionSensor{
		Name: name,
	}
}

// DetectMotion simulates detecting motion.
func (m *MotionSensor) DetectMotion() {
	m.MotionDetected = true
	fmt.Printf("%s detected motion!\n", m.Name)
}

// Alarm represents a simple alarm component.
type Alarm struct {
	Name    string `json:"Name"`
	Armed   bool   `json:"Armed"`
	Sounded bool   `json:"Sounded"`
	Energy  int    `json:"Energy"`
}

// NewAlarm creates a new Alarm instance with the given name.
func NewAlarm(name string) *Alarm {
	return &Alarm{
		Name:  name,
		Armed: false,
	}
}

// Arm sets the alarm to the armed state.
func (a *Alarm) Arm() {
	a.Armed = true
	fmt.Printf("%s is armed and ready.\n", a.Name)
}

// Disarm disarms the alarm.
func (a *Alarm) Disarm() {
	a.Armed = false
	fmt.Printf("%s is disarmed.\n", a.Name)
}

// Trigger activates the alarm if it's armed.
func (a *Alarm) Trigger() {
	if a.Armed {
		a.Sounded = true
		fmt.Printf("%s is triggered! Alarm is sounding.\n", a.Name)
	}
}

// DoorLock structure
type DoorLock struct {
	Location string `json:"Location"`
	Lock     bool   `json:"Lock"`
}

// creates door lock at location like lock kitchen or lock front door
func NewDoorLock(location string, lock bool) *DoorLock {
	return &DoorLock{
		Location: location,
		Lock:     lock,
	}
}

// Creates a lock or unlock feature
func (a *DoorLock) LockOrUnlock(lock bool) {

	if lock == true {
		// if the location is already locked then tell them it is already locked
		if lock == a.Lock {
			fmt.Printf("%s is already locked!\n", a.Location)

			// else lock door at that location
		} else {
			a.Lock = true
			fmt.Printf("%s is now locked!\n", a.Location)
		}

		// if lock is false
	} else {

		// if lock is already false meaning not locked then tell user it is already unlocked
		if lock == a.Lock {
			fmt.Printf("%s is already unlocked!\n", a.Location)
			// else unlock it at that location
		} else {
			a.Lock = false
			fmt.Printf("%s is now unlocked.\n", a.Location)
		}

	}
}

// PadLock structure
type PadLock struct {
	Name string
	Pin  string
	Lock bool
}

// creates Padlock
func NewPadlock(name string, pin string) *PadLock {
	return &PadLock{
		Name: name,
		Pin:  pin,
		Lock: true,
	}
}

// Padlock verify
func (a *PadLock) Verify(pin string) {
	if pin == a.Pin {
		a.Lock = false
		fmt.Printf("%s is unlocked! Welcome!\n", a.Name)
	} else {
		fmt.Printf("Pin is incorrect. Try again.\n")
	}
}

func main() {
	// Create a motion sensor and an alarm
	motionSensor := NewMotionSensor("Motion Sensor")
	securityAlarm := NewAlarm("Security Alarm")
	padLock1 := NewPadlock("PadLock1", "1234")
	frontdoor := NewDoorLock("FrontDoor", true)

	// Arm the security system
	securityAlarm.Arm()

	// Simulate motion detection
	motionSensor.DetectMotion()

	// Check if the alarm is triggered
	securityAlarm.Trigger()

	// Disarm the security system
	securityAlarm.Disarm()

	// Simulate motion detection again
	motionSensor.DetectMotion()

	// Check if the alarm is triggered (it should not be triggered this time)
	securityAlarm.Trigger()

	// checks if the pin worked in this case it doesn't because we set pin to 1234
	padLock1.Verify("4321")

	// pin works and should unlock
	padLock1.Verify("1234")

	// should return already locked
	frontdoor.LockOrUnlock(true)

	// should unlock it
	frontdoor.LockOrUnlock(false)

	// should return already unlocked
	frontdoor.LockOrUnlock(false)

	// should lock it
	frontdoor.LockOrUnlock(true)
}
