package gocode

import (
	"fmt"
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
	RowPins      []int
	ColumnPins   []int
	Keys         [][]Key
	DebounceTime time.Duration
	HoldTime     time.Duration
	Listener     func(rune)
	LastScanTime time.Time
}

// NewKeypad creates a new Keypad instance.
func NewKeypad(rowPins, colPins []int, keys [][]Key) *Keypad {
	return &Keypad{
		RowPins:      rowPins,
		ColumnPins:   colPins,
		Keys:         keys,
		DebounceTime: 10 * time.Millisecond,
		HoldTime:     500 * time.Millisecond,
		LastScanTime: time.Now(),
	}
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
	// Initialize the keypad, set pin modes, etc.
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
	// This is a placeholder for the actual hardware scanning logic.
	// You'll need to implement this based on your hardware setup.
}

// UpdateKeyState updates the state of a given key.
func (k *Keypad) UpdateKeyState(key *Key, newState KeyState) {
	key.State = newState
	key.StateChanged = true
	if k.Listener != nil {
		k.Listener(key.Char)
	}
}

func main() {
	// Example usage
	rowPins := []int{1, 2, 3, 4}
	columnPins := []int{5, 6, 7}
	keys := [][]Key{
		{{'1', 0, IDLE, false}, {'2', 1, IDLE, false}, {'3', 2, IDLE, false}},
		{{'4', 3, IDLE, false}, {'5', 4, IDLE, false}, {'6', 5, IDLE, false}},
		{{'7', 6, IDLE, false}, {'8', 7, IDLE, false}, {'9', 8, IDLE, false}},
		{{'*', 9, IDLE, false}, {'0', 10, IDLE, false}, {'#', 11, IDLE, false}},
	}
	keypad := NewKeypad(rowPins, columnPins, keys)
	keypad.Begin()

	// Set an event listener
	keypad.AddEventListener(func(char rune) {
		fmt.Printf("Key pressed: %c\n", char)
	})

	// Main loop
	for {
		keypad.GetKeys()
		time.Sleep(100 * time.Millisecond)
	}
}
