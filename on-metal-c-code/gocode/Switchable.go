package gocode

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
)

// SwitchOn turns the buzzer on.
func Switchable(pinNum uint8, status bool) {
	if err := rpio.Open(); err != nil {
		fmt.Println("Unable to open GPIO:", err)
		return
	}
	defer func() {
		err := rpio.Close()
		if err != nil {

		}
	}()
	if status == true {
		pin := rpio.Pin(pinNum)
		pin.Toggle()
		fmt.Printf("on pin %d is switched on.\n", pinNum)
	}
	if status == false {
		pin := rpio.Pin(pinNum)
		pin.Toggle()
		fmt.Printf("Buzzer on pin %d is switched off.\n", pinNum)
	}
}
