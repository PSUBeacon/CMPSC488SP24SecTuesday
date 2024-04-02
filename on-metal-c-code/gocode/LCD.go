package gocode

import (
	lcd "github.com/wjessop/lcm1602_lcd"
	"golang.org/x/exp/io/i2c"
	"log"
)

type LCD interface {
	WriteLCD()
	ClearLCD()
}

func WriteLCD(LCDMessage string) string {
	rownum := 1
	if len(LCDMessage) > 16 {

		rownum = 2
		return "LCD message longer than 16"

	}

	// Configure this line with the device location and address of your device
	lcdDevice, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, 0x27)
	if err != nil {
		log.Fatal(err)
	}
	defer lcdDevice.Close()

	lcdDisplay, err := lcd.NewLCM1602LCD(lcdDevice)
	if err != nil {
		log.Fatal(err)
	}

	// Write a string to row 1, position 0 (ie, the start of the line)
	err = lcdDisplay.WriteString(LCDMessage, rownum, 0)
	if err != nil {
		log.Fatal(err)
	}
	return ""
}
func ClearLCD() {
	// Configure this line with the device location and address of your device
	lcdDevice, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, 0x27)
	if err != nil {
		log.Fatal(err)
	}
	defer lcdDevice.Close()

	lcdDisplay, err := lcd.NewLCM1602LCD(lcdDevice)
	if err != nil {
		log.Fatal(err)
	}
	if err := lcdDisplay.Clear(); err != nil {
		log.Fatal(err)

	}
}
