package lighting

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"fmt"
)

const (
	dinPinNumber = 9  // GPIO pin for DIN (MOSI)
	csPinNumber  = 4  // GPIO pin for CS
	clkPinNumber = 10 // GPIO pin for CLK
)

// TurnOn turns the lighting on.
func UpdateStatus(newStatus bool) {
	fmt.Printf("%s is now turned \n", newStatus)
	gocode.MatrixStatus(9, 4, 10, newStatus)

}

// SetBrightness sets the brightness of the lighting.
func SetBrightness(brightness int) {
	//if brightness < 0 {
	//	brightness = 0
	//} else if brightness > 100 {
	//	brightness = 100
	//}
	//gocode.DrawLightbulb(9, 4, 10)
	//time.Sleep(3 * time.Second)
	//gocode.ClearMatrix(9, 4, 10)
	gocode.SetIntensity(9, 4, 10, brightness)
	fmt.Printf("%s brightness is set to %s\n", brightness)
}

//func drawBulblTimer(dinPin, csPin, clkPin rpio.Pin) {
//	gocode.DrawLightbulb(dinPin, csPin, clkPin)
//	time.Sleep(3 * time.Second)
//	gocode.ClearMatrix(dinPin, csPin, clkPin)
//
//}
