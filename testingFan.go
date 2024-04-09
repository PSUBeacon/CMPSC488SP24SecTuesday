//package main
//
//import (
//	"fmt"
//	"os"
//	"time"
//
//	"github.com/stianeikeland/go-rpio/v4"
//)
//
//func main() {
//	// Open and map memory to access GPIO, check for errors
//	if err := rpio.Open(); err != nil {
//		fmt.Println("Unable to open GPIO:", err)
//		os.Exit(1)
//	}
//	defer rpio.Close()
//
//	// Set pin to output mode
//	pin := rpio.Pin(14)
//	pin.Output()
//
//	for i := 0; i < 5; i++ {
//		// Turn the fan on
//		pin.High()
//		fmt.Println("Fan ON")
//		time.Sleep(2 * time.Second)
//
//		// Turn the fan off
//		pin.Low()
//		fmt.Println("Fan OFF")
//		time.Sleep(2 * time.Second)
//	}
//}

package main

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
)

const (
	LowSpeed    = 20
	MediumSpeed = 50
	HighSpeed   = 90
)

func main() {
	gocode.SetFanSpeed(12, LowSpeed)
}
