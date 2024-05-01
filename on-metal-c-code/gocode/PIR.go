package gocode

import (
	"github.com/stianeikeland/go-rpio/v4"
)

// CheckForMotion initializes the GPIO then checks for motion
func CheckForMotion(pinNum uint8) (bool, error) {
	//if err := rpio.Open(); err != nil {
	//	return false, err
	//}
	//defer func() {
	//	err := rpio.Close()
	//	if err != nil {
	//
	//	}
	//}()

	pin := rpio.Pin(pinNum)
	pin.Input()
	pin.PullDown()

	motionDetected := pin.Read() == rpio.High
	return motionDetected, nil
}
