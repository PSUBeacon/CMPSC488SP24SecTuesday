package lighting

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"encoding/json"
	"fmt"
	"time"
)

// Lighting represents the state of a smart light in the system.
type Lighting struct {
	UUID              string               `json:"uuid"`              // Unique identifier for the light
	Location          string               `json:"location"`          // Location of the light (e.g., "Living Room", "Kitchen").
	Brightness        string               `json:"brightness"`        // Brightness level as a string to include numeric value.
	Status            bool                 `json:"status"`            // true for on, false for off.
	EnergyConsumption int                  `json:"energyConsumption"` // Energy consumption in kilowatts.
	LastChanged       string               `json:"lastChanged"`       // Timestamp of the last change in ISODate format.
	Timers            []Timer              `json:"timers"`            // Timers for automatic on/off.
	Schedules         []Schedule           `json:"schedules"`         // Daily or weekly schedules.
	MoodScenes        map[string]MoodScene `json:"moodScenes"`        // Saved lighting scenes or moods.
}

// Timer defines a simple timer for turning the light on or off at a specific time.
type Timer struct {
	Time  string `json:"time"`  // Time when the light should change status.
	State bool   `json:"state"` // Desired state at the specified time: true for on, false for off.
}

// Schedule defines a lighting schedule, e.g., for daily or weekly routines.
type Schedule struct {
	DaysOfWeek []time.Weekday `json:"daysOfWeek"` // Days of the week the schedule applies to.
	Time       string         `json:"time"`       // Time of day for the action.
	State      bool           `json:"state"`      // Desired state: true for on, false for off.
}

// MoodScene represents a lighting scene or mood with specific settings.
type MoodScene struct {
	Brightness string `json:"brightness"` // Brightness setting for the mood.
	Color      string `json:"color"`      // Color setting for the mood, if applicable.
}

// NewLighting creates a new Lighting instance with the given parameters.
//func NewLighting(uuid, location string, status bool, brightness string, energyConsumption int) *Lighting {
//	return &Lighting{
//		UUID:              uuid,
//		Location:          location,
//		Status:            status,
//		Brightness:        brightness,
//		EnergyConsumption: energyConsumption,
//		LastChanged:       time.Now().Format(time.RFC3339),
//		MoodScenes:        make(map[string]MoodScene),
//	}
//}

// TurnOn turns the lighting on.
func UpdateStatus(newStatus bool) {
	fmt.Printf("%s is now turned \n", newStatus)
	//gocode.Initialize(7)
	gocode.BuzzerStatus(7, newStatus)
}

// SetBrightness sets the brightness of the lighting.
func SetBrightness(brightness int) {
	if brightness < 0 {
		brightness = 0
	} else if brightness > 100 {
		brightness = 100
	}
	fmt.Printf("%s brightness is set to %s\n", brightness)
}

// AddTimer adds a new timer to the lighting system.
func (l *Lighting) AddTimer(timer Timer) {
	l.Timers = append(l.Timers, timer)
}

// AddSchedule adds a new schedule for the lighting system.
func (l *Lighting) AddSchedule(schedule Schedule) {
	l.Schedules = append(l.Schedules, schedule)
}

// SetMoodScene creates or updates a mood scene.
func (l *Lighting) SetMoodScene(name string, scene MoodScene) {
	l.MoodScenes[name] = scene
}

// ActivateMoodScene applies a previously saved mood scene.
func (l *Lighting) ActivateMoodScene(name string) {
	scene, exists := l.MoodScenes[name]
	if !exists {
		fmt.Println("Mood scene does not exist:", name)
		return
	}
	l.Brightness = scene.Brightness
	// Assuming color functionality exists; simulate setting color here.
	fmt.Printf("Activated mood scene '%s' with brightness %s and color %s\n", name, scene.Brightness, scene.Color)
}

// serializeLight converts a Lighting object to a JSON string.
func serializeLight(light *Lighting) (string, error) {
	lightJSON, err := json.Marshal(light)
	if err != nil {
		return "", err
	}
	return string(lightJSON), nil
}

// deserializeLight converts a JSON string back into a Lighting object.
func deserializeLight(lightJSON string) (*Lighting, error) {
	var light Lighting
	err := json.Unmarshal([]byte(lightJSON), &light)
	if err != nil {
		return nil, err
	}
	return &light, nil
}

//
//func main() {
//	// Example setup and usage
//	lightUUID := "123e4567-e89b-12d3-a456-426614174000"
//	light := NewLighting(lightUUID, "Living Room", false, "50", 10)
//
//	// Adding a timer
//	light.AddTimer(Timer{Time: "2023-10-05T19:00:00Z", State: true})
//
//	// Adding a schedule
//	light.AddSchedule(Schedule{DaysOfWeek: []time.Weekday{time.Monday, time.Wednesday, time.Friday}, Time: "18:00", State: true})
//
//	// Setting and activating a mood scene
//	light.SetMoodScene("Movie Night", MoodScene{Brightness: "30", Color: "Warm White"})
//	light.ActivateMoodScene("Movie Night")
//
//	// Serialization example
//	lightJSON, err := serializeLight(light)
//	if err != nil {
//		fmt.Println("Error serializing light:", err)
//		return
//	}
//	fmt.Println("Serialized light:", lightJSON)
//
//	// Deserialization example
//	deserializedLight, err := deserializeLight(lightJSON)
//	if err != nil {
//		fmt.Println("Error deserializing light:", err)
//		return
//	}
//	fmt.Printf("Deserialized light: %+v\n", deserializedLight)
//}
