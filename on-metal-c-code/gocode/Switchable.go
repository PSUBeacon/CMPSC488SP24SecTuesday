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
	fmt.Printf("%s status.\n", status)
	fmt.Printf("on pin %s", pinNum)
	if status == true {
		pin := rpio.Pin(pinNum)
		pin.Mode(rpio.Output) // Alternative syntax
		pin.Write(rpio.High)
		fmt.Printf("on pin %d is switched on.\n", pinNum)
		rpio.Close()
	}
	if status == false {
		pin := rpio.Pin(pinNum)
		pin.Output()
		pin.Low()
		fmt.Printf("Buzzer on pin %d is switched off.\n", pinNum)
		rpio.Close()
	}
	pin := rpio.Pin(17)
	pin.Mode(rpio.Output) // Alternative syntax
	pin.Write(rpio.High)
	fmt.Printf("on pin %d is switched on.\n", pinNum)
	rpio.Close()
}
