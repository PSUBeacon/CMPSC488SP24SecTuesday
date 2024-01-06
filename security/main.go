package main

import (
	"fmt"
)

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

func main() {
	// Create a motion sensor and an alarm
	motionSensor := NewMotionSensor("Motion Sensor")
	securityAlarm := NewAlarm("Security Alarm")

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
}
