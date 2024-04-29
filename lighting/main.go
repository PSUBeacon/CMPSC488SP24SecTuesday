package lighting

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"fmt"
	"time"
)

const (
	dinPinNumber = 9  // GPIO pin for DIN (MOSI)
	csPinNumber  = 4  // GPIO pin for CS
	clkPinNumber = 10 // GPIO pin for CLK
)

var GlobBrightness = 15

// TurnOn turns the lighting on.
func UpdateStatus(newStatus bool) {
	fmt.Printf("%s is now turned \n", newStatus)
	gocode.MatrixStatus(9, 4, 10, newStatus, GlobBrightness)

}

// SetBrightness sets the brightness of the lighting.
func SetBrightness(brightness int) {
	if brightness < 0 {
		brightness = 0
	} else if brightness > 100 {
		brightness = 100
	}
	gocode.DrawLightbulb(9, 4, 10, brightness)
	time.Sleep(3 * time.Second)
	gocode.TurnOffMatrix(9, 4, 10)
	gocode.TurnOnMatrix(9, 4, 10)
	GlobBrightness = brightness
	gocode.SetIntensity(9, 4, 10, brightness)
	fmt.Printf("%s brightness is set to %s\n", brightness)
}

//func drawBulblTimer(dinPin, csPin, clkPin rpio.Pin) {
//	gocode.DrawLightbulb(dinPin, csPin, clkPin)
//	time.Sleep(3 * time.Second)
//	gocode.ClearMatrix(dinPin, csPin, clkPin)
//
//}
