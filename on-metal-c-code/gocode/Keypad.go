package gocode

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"os"
	"time"
)

// KeyState represents the state of a key.
type KeyState int

// Constants for different key states.
const (
	IDLE KeyState = iota
	PRESSED
	HOLD
	RELEASED
)

// Key represents a single key.
type Key struct {
	Char         rune
	Code         int
	State        KeyState
	StateChanged bool
}

// Keypad represents a keypad.
type Keypad struct {
	RowPins      []rpio.Pin
	ColumnPins   []rpio.Pin
	Keys         [][]Key
	DebounceTime time.Duration
	HoldTime     time.Duration
	Listener     func(rune)
	LastScanTime time.Time
}

// NewKeypad creates a new Keypad instance.
func NewKeypad(rowPins, colPins []int, keys [][]Key) *Keypad {
	kp := &Keypad{
		RowPins:      make([]rpio.Pin, len(rowPins)),
		ColumnPins:   make([]rpio.Pin, len(colPins)),
		Keys:         keys,
		DebounceTime: 10 * time.Millisecond,
		HoldTime:     500 * time.Millisecond,
		LastScanTime: time.Now(),
	}
	for i, pin := range rowPins {
		kp.RowPins[i] = rpio.Pin(pin)
		kp.RowPins[i].Output()
	}
	for i, pin := range colPins {
		kp.ColumnPins[i] = rpio.Pin(pin)
		kp.ColumnPins[i].Input()
		kp.ColumnPins[i].PullDown()
	}
	return kp
}

// SetDebounceTime sets the debounce time for the keypad.
func (k *Keypad) SetDebounceTime(duration time.Duration) {
	k.DebounceTime = duration
}

// SetHoldTime sets the hold time for the keypad.
func (k *Keypad) SetHoldTime(duration time.Duration) {
	k.HoldTime = duration
}

// AddEventListener sets the event listener for the keypad.
func (k *Keypad) AddEventListener(listener func(rune)) {
	k.Listener = listener
}

// Begin initializes the keypad.
func (k *Keypad) Begin() {
	// Nothing to do here for now
}

// GetKey returns a single key press.
func (k *Keypad) GetKey() rune {
	k.ScanKeys()
	for _, row := range k.Keys {
		for _, key := range row {
			if key.State == PRESSED && key.StateChanged {
				return key.Char
			}
		}
	}
	return ' '
}

// GetKeys scans and updates the state of all keys.
func (k *Keypad) GetKeys() {
	currentTime := time.Now()
	if currentTime.Sub(k.LastScanTime) > k.DebounceTime {
		k.ScanKeys()
		k.LastScanTime = currentTime
	}
}

// ScanKeys scans the keys and updates their states.
func (k *Keypad) ScanKeys() {
	for i, rowPin := range k.RowPins {
		rowPin.High()
		for j, colPin := range k.ColumnPins {
			if colPin.Read() == rpio.High {
				k.UpdateKeyState(&k.Keys[i][j], PRESSED)
			} else {
				k.UpdateKeyState(&k.Keys[i][j], IDLE)
			}
		}
		rowPin.Low()
	}
}

// UpdateKeyState updates the state of a given key.
func (k *Keypad) UpdateKeyState(key *Key, newState KeyState) {
	if key.State != newState {
		key.State = newState
		key.StateChanged = true
		if newState == PRESSED && k.Listener != nil {
			k.Listener(key.Char)
		}
	} else {
		key.StateChanged = false
	}
}

//const securityCode = "123456" // Exam

var enteredCode string

type System struct {
	UUID         string `json:"UUID"`
	Status       string `json:"Status"`
	SensorStatus string `json:"SensorStatus"`
}

func onKeyPress(char rune) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	securityCode := os.Getenv("SECURITY_KEY")
	if len(enteredCode) < 6 {
		// Append pressed key to the entered code sequence
		enteredCode += string(char)

		WriteLCD("Code: " + enteredCode)

		//Check if the entered code matches the security code
		if enteredCode == securityCode {
			// Trigger the appropriate action when the correct code is entered
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

			if securitySystem.Status == "Armed" {
				fmt.Println("Security code entered correctly. disarming alarm...")
				securitySystem.Status = "Disarmed"
				securitySystem.SensorStatus = "OFF"
				WriteLCD("Stat:" + securitySystem.Status + " Motion:" + securitySystem.SensorStatus)
				enteredCode = ""
				securityJSON, err := json.MarshalIndent(securitySystem, "", "	")
				if err != nil {
					fmt.Println("Error marshalling security data:", err)
					return
				}

				err = os.WriteFile("security/Keys.json", securityJSON, 0644)
				if err != nil {
					panic(err)
				}
				return

			}
			if securitySystem.Status == "Disarmed" {
				fmt.Println("Security code entered correctly. arming alarm...")
				securitySystem.Status = "Armed"
				securitySystem.SensorStatus = "ON"
				WriteLCD("Stat:" + securitySystem.Status + " Motion:" + securitySystem.SensorStatus)
				enteredCode = ""
				securityJSON, err := json.MarshalIndent(securitySystem, "", "	")
				if err != nil {
					fmt.Println("Error marshalling security data:", err)
					return
				}

				err = os.WriteFile("security/Keys.json", securityJSON, 0644)
				if err != nil {
					panic(err)
				}
				return
			}
		}
		if len(enteredCode) == 6 {
			if enteredCode != securityCode {
				// Trigger the appropriate action when the correct code is entered
				fmt.Println("Security code entered incorrectly")
				// Call a function to disarm the alarm, for example
				WriteLCD("Code: ")
				enteredCode = ""
			}
		}
	}
}

// Initialize the keypad and start listening for key presses
func InitKeypad() rune {
	if err := rpio.Open(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open GPIO: %v\n", err)
		os.Exit(1)
	}

	// GPIO pin numbers for rows and columns
	rowPins := []int{22, 23, 24, 25}
	columnPins := []int{17, 16, 27}

	// Define keys
	keys := [][]Key{
		{{'1', 0, IDLE, false}, {'2', 1, IDLE, false}, {'3', 2, IDLE, false}},
		{{'4', 3, IDLE, false}, {'5', 4, IDLE, false}, {'6', 5, IDLE, false}},
		{{'7', 6, IDLE, false}, {'8', 7, IDLE, false}, {'9', 8, IDLE, false}},
		{{'*', 9, IDLE, false}, {'0', 10, IDLE, false}, {'#', 11, IDLE, false}},
	}

	// Create and initialize the keypad
	keypad := NewKeypad(rowPins, columnPins, keys)
	keypad.Begin()

	// Set the event listener
	keypad.AddEventListener(onKeyPress)

	//	Main loop
	for {
		// Scan keys periodically
		keypad.GetKeys()
		time.Sleep(100 * time.Millisecond)
	}

	//for {
	//	// Scan keys periodically
	//	key := keypad.GetKey()
	//	time.Sleep(100 * time.Millisecond)
	//	if key != ' ' {
	//		return key // Return the pressed key
	//	}
	//}
}

//func main() {
//	if err := rpio.Open(); err != nil {
//		fmt.Printf("Error opening GPIO: %v\n", err)
//		return
//	}
//	defer rpio.Close()
//
//	// Example usage
//	rowPins := []int{22, 23, 24, 25} // GPIO pin numbers for rows
//	columnPins := []int{17, 18, 27}  // GPIO pin numbers for columns
//	keys := [][]Key{
//		{{'1', 0, IDLE, false}, {'2', 1, IDLE, false}, {'3', 2, IDLE, false}},
//		{{'4', 3, IDLE, false}, {'5', 4, IDLE, false}, {'6', 5, IDLE, false}},
//		{{'7', 6, IDLE, false}, {'8', 7, IDLE, false}, {'9', 8, IDLE, false}},
//		{{'*', 9, IDLE, false}, {'0', 10, IDLE, false}, {'#', 11, IDLE, false}},
//	}
//	keypad := NewKeypad(rowPins, columnPins, keys)
//	keypad.Begin()
//
//	// Set an event listener
//	keypad.AddEventListener(func(char rune) {
//		fmt.Printf("Key pressed: %c\n", char)
//	})
//
//	// Main loop
//	for {
//		keypad.GetKeys()
//		time.Sleep(100 * time.Millisecond)
//	}
//}
