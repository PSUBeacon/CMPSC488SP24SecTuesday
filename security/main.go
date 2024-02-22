package main

import (
	"fmt"
)

// Here is the Security System Structure
type SecuritySystem struct {
	Location          string //(mcq)
	SensorType        string //(mcq)
	Status            bool
	EnergyConsumption int    //(kilowatts)
	LastTriggered     string //(this will be the main notification)
}

// MotionSensor represents a motion sensor component.

type MotionSensor struct {
	Name           string
	MotionDetected bool
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
	Name    string
	Armed   bool
	Sounded bool
	Energy  int
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

// PadLock strucute
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
}

