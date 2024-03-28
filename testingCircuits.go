//package main
//
//import (
//	"github.com/stianeikeland/go-rpio/v4"
//	"log"
//	"time"
//)
//
////
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

package main

import (
	"log"
	"time"

	"github.com/warthog618/go-gpiocdev"
)

func main() {
	chip, err := gpiocdev.NewChip("gpiochip0", gpiocdev.WithConsumer("8x8matrix"))
	if err != nil {
		log.Fatalf("Failed to open chip: %v", err)
	}
	defer chip.Close()

	din, err := chip.RequestLine(9, gpiocdev.AsOutput(0))
	if err != nil {
		log.Fatalf("Failed to request DIN line: %v", err)
	}
	defer din.Close()

	clk, err := chip.RequestLine(10, gpiocdev.AsOutput(0))
	if err != nil {
		log.Fatalf("Failed to request CLK line: %v", err)
	}
	defer clk.Close()

	cs, err := chip.RequestLine(4, gpiocdev.AsOutput(1))
	if err != nil {
		log.Fatalf("Failed to request CS line: %v", err)
	}
	defer cs.Close()

	// Initialize the matrix
	cs.SetValue(0)
	for i := 0; i < 8; i++ {
		sendByte(din, clk, byte(i+1)) // Address
		sendByte(din, clk, 0xFF)      // All LEDs on
	}
	cs.SetValue(1)

	// Keep the matrix on for a while
	time.Sleep(5 * time.Second)
}

func sendByte(din, clk *gpiocdev.Line, b byte) {
	for i := 0; i < 8; i++ {
		din.SetValue(int((b >> (7 - i)) & 0x01))
		clk.SetValue(1)
		time.Sleep(1 * time.Microsecond) // Small delay
		clk.SetValue(0)
	}
}
