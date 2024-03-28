package main

import (
	"github.com/d2r2/go-max7219"
	"log"
)

//
//import (
//	"flag"
//	"fmt"
//	"golang.org/x/exp/io/i2c"
//	"os"
//	"periph.io/x/periph/conn/i2c/i2creg"
//	"periph.io/x/periph/experimental/devices/hd44780"
//	"periph.io/x/periph/host"
//	"log"
//	"strings"
//	"time"
//	"periph.io/x/conn/v3/gpio"
//	"periph.io/x/conn/v3/gpio/gpioreg"
//	"periph.io/x/devices/v3/hd44780"
//	"periph.io/x/host/v3"
//)

//func main() {
//	gocode.Switchable(17, true)
//
//}

func main() {
	// Create a new device with 1 cascaded MAX7219
	device := max7219.NewDevice(1)

	// Open the device (using SPI bus 0, device 0)
	if err := device.Open(0, 0, 15); err != nil {
		log.Fatal(err)
	}
	defer device.Close()

	// Set all LEDs to on
	for i := 0; i < 8; i++ {
		if err := device.SetBufferLine(0, i, 0xFF, true); err != nil {
			log.Fatal(err)
		}
	}

	// Update the display
	if err := device.Flush(); err != nil {
		log.Fatal(err)
	}
}
