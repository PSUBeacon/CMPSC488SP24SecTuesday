package main

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"github.com/d2r2/go-max7219"
	"log"
	"time"
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
	gocode.ClearLCD()
	gocode.WriteLCD("This worked")

	// Create new LED matrix with number of cascaded devices is equal to 1
	mtx := max7219.NewMatrix(1)
	// Open SPI device with spibus and spidev parameters equal to 0 and 0.
	// Set LED matrix brightness is equal to 7
	err := mtx.Open(9, 10, 7)
	if err != nil {
		log.Fatal(err)
	}
	defer mtx.Close()
	// Output text message to LED matrix
	// Output a sequence of ascii codes in a loop
	font := max7219.FontCP437
	for i := 0; i <= len(font.GetLetterPatterns()); i++ {
		mtx.OutputAsciiCode(72, font, i, true)
		time.Sleep(500 * time.Millisecond)
	}

}
