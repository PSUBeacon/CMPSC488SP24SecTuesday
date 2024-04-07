package gocode

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
)

func FanStatus(fanPin uint8, status bool) {
	if err := rpio.Open(); err != nil {
		fmt.Println("Unable to open GPIO:", err)
		return
	}
	defer rpio.Close()
	fmt.Printf("%s status.\n", status)
	fmt.Printf("on pin %s", fanPin)
	if status == true {
		pin := rpio.Pin(fanPin)
		pin.Mode(rpio.Output)
		pin.Write(rpio.High)
		fmt.Printf("on pin %d is switched on.\n", pin)
	}
	if status == false {
		pin := rpio.Pin(fanPin)
		pin.Output()
		pin.Low()
		fmt.Printf("Fan on pin %d is switched off.\n", pin)
	}
}

func SetFanSpeed(fanPin uint8, speed int) {
	if err := rpio.Open(); err != nil {
		fmt.Println("Unable to open GPIO:", err)
		return
	}
	defer rpio.Close()
	pin := rpio.Pin(fanPin)
	pin.Mode(rpio.Pwm)
	pin.Freq(64000)
	pin.DutyCycle(uint32(speed), 100)
}

// Set the intensity (brightness) of the LED matrix
func setSpeed(intensity byte) {
	dinPin := rpio.Pin(dinPinNumber)
	csPin := rpio.Pin(csPinNumber)
	clkPin := rpio.Pin(clkPinNumber)

	initializeMatrix(dinPin, csPin, clkPin)

	if intensity > 0x0F {
		intensity = 0x0F // Maximum intensity value is 0x0F
	}
	sendData(csPin, dinPin, clkPin, 0x0A, intensity)
}
