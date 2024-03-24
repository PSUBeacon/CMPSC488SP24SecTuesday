package security

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

// creates new security system. I gotta figure out how to connect sensor type with sensor or location with the door strucute
func NewSecuirtySystem(location string, sensorType string, status bool, energyConsuption int, lastTriggered string) *SecuritySystem {
	return &SecuritySystem{
		Location:          location,
		SensorType:        sensorType,
		Status:            status,
		EnergyConsumption: energyConsuption,
		LastTriggered:     lastTriggered,
	}
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
func UpdateAlarmStatus(status bool) {
	fmt.Printf("%s Alarm status is set to: \n", status)
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

//func main() {
//	// Create a security system, motion sensor, alarm, padlock, frontdoor, backdoor
//	securitySystem := NewSecuirtySystem("House1", "imaginary sensor", true, 3, "02/21 10:32:05PM '24 -0700")
//	motionSensor := NewMotionSensor("Motion Sensor")
//	securityAlarm := NewAlarm("Security Alarm")
//	padLock1 := NewPadlock("PadLock1", "1234")
//	frontdoor := NewDoorLock("FrontDoor", true)
//	backdoor := NewDoorLock("BackDoor", true)
//
//	//Serializations---------------------------------
//
//	//security
//	securityJSON, err := json.Marshal(securitySystem)
//	if err != nil {
//		fmt.Println("Error:", err)
//	}
//
//	//motionSensor
//	motionSensorJSON, err := json.Marshal(motionSensor)
//	if err != nil {
//		fmt.Println("Error:", err)
//	}
//
//	//security alarm
//	securityAlarmJSON, err := json.Marshal(securityAlarm)
//	if err != nil {
//		fmt.Println("Error:", err)
//	}
//
//	//padlock
//	padLockJSON, err := json.Marshal(padLock1)
//	if err != nil {
//		fmt.Println("Error:", err)
//	}
//
//	frontDoorJSON, err := json.Marshal(frontdoor)
//	if err != nil {
//		fmt.Println("Error:", err)
//	}
//
//	backDoorJSON, err := json.Marshal(backdoor)
//	if err != nil {
//		fmt.Println("Error:", err)
//	}
//
//	//end of Serializations------------------------------
//	// Arm the security system
//	securityAlarm.Arm()
//
//	// Simulate motion detection
//	motionSensor.DetectMotion()
//
//	// Check if the alarm is triggered
//	securityAlarm.Trigger()
//
//	// Disarm the security system
//	securityAlarm.Disarm()
//
//	// Simulate motion detection again
//	motionSensor.DetectMotion()
//
//	// Check if the alarm is triggered (it should not be triggered this time)
//	securityAlarm.Trigger()
//
//	// checks if the pin worked in this case it doesn't because we set pin to 1234
//	padLock1.Verify("4321")
//
//	// pin works and should unlock
//	padLock1.Verify("1234")
//
//	// should return already locked
//	frontdoor.LockOrUnlock(true)
//
//	// should unlock it
//	frontdoor.LockOrUnlock(false)
//
//	// should return already unlocked
//	frontdoor.LockOrUnlock(false)
//
//	// should lock it
//	frontdoor.LockOrUnlock(true)
//
//	// should return already locked
//	backdoor.LockOrUnlock(true)
//
//	// should unlock it
//	backdoor.LockOrUnlock(false)
//
//	// should return already unlocked
//	backdoor.LockOrUnlock(false)
//
//	// should lock it
//	frontdoor.LockOrUnlock(true)
//
//	// Deserializations -----------------------------
//
//	// security
//	var loadedSecurity SecuritySystem
//	err = json.Unmarshal(securityJSON, &loadedSecurity)
//	if err != nil {
//		fmt.Println("Error", err)
//		return
//	}
//
//	// Motion Sensor
//	var loadedMotionSensor MotionSensor
//	err = json.Unmarshal(motionSensorJSON, &loadedMotionSensor)
//	if err != nil {
//		fmt.Println("Error", err)
//		return
//	}
//
//	// Security Alarm
//	var loadedSecurityAlarm Alarm
//	err = json.Unmarshal(securityAlarmJSON, &loadedSecurityAlarm)
//	if err != nil {
//		fmt.Println("Error", err)
//		return
//	}
//
//	// PadLock
//	var loadedPadLock PadLock
//	err = json.Unmarshal(padLockJSON, &loadedPadLock)
//	if err != nil {
//		fmt.Println("Error", err)
//		return
//	}
//
//	// Front Door
//	var loadedFrontDoor DoorLock
//	err = json.Unmarshal(frontDoorJSON, &loadedFrontDoor)
//	if err != nil {
//		fmt.Println("Error", err)
//		return
//	}
//
//	// Front Door
//	var loadedBackDoor DoorLock
//	err = json.Unmarshal(backDoorJSON, &loadedFrontDoor)
//	if err != nil {
//		fmt.Println("Error", err)
//		return
//	}
//
//	//end -------------------------------------------
//
//	// serialization prints-------------------------------
//
//	fmt.Println("\n\nHere are the Serializations:\n\nSecurity Object Serialization")
//	fmt.Println(string(securityJSON))
//	fmt.Println()
//
//	fmt.Println("Motion Sensor Object Serialization")
//	fmt.Println(string(motionSensorJSON))
//	fmt.Println()
//
//	fmt.Println("Security Alarm Object Serialization")
//	fmt.Println(string(securityAlarmJSON))
//	fmt.Println()
//
//	fmt.Println("PadLock Object Serialization")
//	fmt.Println(string(padLockJSON))
//	fmt.Println()
//
//	fmt.Println("FrontDoor Object Serialization")
//	fmt.Println(string(frontDoorJSON))
//	fmt.Println()
//
//	fmt.Println("BackDoor Object Serialization")
//	fmt.Println(string(backDoorJSON))
//	fmt.Println()
//	fmt.Println()
//	// ends------------------------------------------
//
//	//Deserialization prints---------------------------
//	fmt.Println("Deserialization Prints:\n\n")
//
//	//Security
//	fmt.Println("Deserialization Security System")
//	fmt.Println("Location: ", loadedSecurity.Location)
//	fmt.Println("SensorType: ", loadedSecurity.SensorType)
//	fmt.Println("Status: ", loadedSecurity.Status)
//	fmt.Println("EnergyConsumption: ", loadedSecurity.EnergyConsumption)
//	fmt.Println("LastTriggred: ", loadedSecurity.LastTriggered)
//	fmt.Println()
//
//	//Motion Sensor
//	fmt.Println("Deserialization Motion Sensor")
//	fmt.Println("Name: ", loadedMotionSensor.Name)
//	fmt.Println("MotionDetected: ", loadedMotionSensor.MotionDetected)
//	fmt.Println()
//
//	//Alarm
//	fmt.Println("Deserialization Alarm")
//	fmt.Println("Name: ", loadedSecurityAlarm.Name)
//	fmt.Println("Armed: ", loadedSecurityAlarm.Armed)
//	fmt.Println("Sounded: ", loadedSecurityAlarm.Sounded)
//	fmt.Println("Energy: ", loadedSecurityAlarm.Energy)
//	fmt.Println()
//
//	//FrontDoor
//	fmt.Println("Deserialization Front Door")
//	fmt.Println("Location: ", loadedFrontDoor.Location)
//	fmt.Println("Lock: ", loadedFrontDoor.Lock)
//	fmt.Println()
//
//	//BackDoor
//	fmt.Println("Deserialization Back Door")
//	fmt.Println("Location: ", loadedBackDoor.Location)
//	fmt.Println("Lock: ", loadedBackDoor.Lock)
//	fmt.Println()
//
//	// ends------------------------------------------------
//
//}
