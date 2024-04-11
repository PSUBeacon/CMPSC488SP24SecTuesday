package security

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"fmt"
	"time"
)

var (
	pirPin       uint8 = 11
	motion       bool  // Motion detection status
	buzzerPin    uint8 = 4
	keypadRowPin []int
	keypadColPin []int
)

// DetectMotion simulates detecting motion.
func DetectMotion() {
	motion, err := gocode.CheckForMotion(pirPin)
	if err != nil {
		fmt.Println("Error detecting motion:", err)
		return
	}
	if motion {
		// If motion detected, turn buzzer on
		gocode.BuzzerStatus(buzzerPin, true)
		fmt.Println("Motion detected!")
	}
}

// UpdateAlarmStatus sets the alarm status and triggers motion detection if armed.
func UpdateAlarmStatus(armed bool) {
	if armed {
		fmt.Println("Alarm armed.")
		// Detect motion only when alarm is armed
		DetectMotion()
	} else {
		// Turn off buzzer if alarm is disarmed
		gocode.BuzzerStatus(buzzerPin, false)
		fmt.Println("Alarm disarmed.")
	}
}

// LockOrUnlock simulates locking or unlocking a door.
func LockOrUnlock(lock bool) {
	fmt.Println("Door is locked:", lock)
	// Add logic to control door lock here
}

// HandleMotionDetection simulates periodic motion detection.
func HandleMotionDetection() {
	// Simulate motion detection every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Simulate motion detection by toggling motion status
		motion = !motion

		// Update security system based on motion status
		if motion {
			UpdateAlarmStatus(true)
		} else {
			UpdateAlarmStatus(false)
		}
	}
}

// DisplayLCDSecurity displays security system information on an LCD
func DisplayLCDSecurity(status string, motionStatus string) {
	// Update LCD display with security system information
	gocode.WriteLCD("Status: " + status + " Motion: " + motionStatus)
}

// Example function to demonstrate usage
func main() {
	// Start motion detection routine
	go HandleMotionDetection()

	// Simulate changing status and motion detection
	time.Sleep(10 * time.Second)
	DisplayLCDSecurity("Armed", "Motion detected")

	// Keep the program running
	select {}
}
