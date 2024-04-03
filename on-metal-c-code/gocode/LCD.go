package gocode

import (
	"fmt"
	lcd "github.com/wjessop/lcm1602_lcd"
	"golang.org/x/exp/io/i2c"
	"log"
	"strings"
)

type LCD interface {
	WriteLCD()
	ClearLCD()
}

func WriteLCD(LCDMessage string) string {
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

	if len(LCDMessage) <= 16 {
		err = lcdDisplay.WriteString(LCDMessage, 0, 0)
		if err != nil {
			log.Fatal(err)
		}

	} else if len(LCDMessage) <= 32 {
		lines := strings.Split(LCDMessage, "")
		line1 := strings.Join(lines[:16], "")
		line2 := strings.Join(lines[16:], "")

		// Write to both lines
		err = lcdDisplay.WriteString(line1, 1, 0)
		if err != nil {
			log.Fatal(err)
		}
		err = lcdDisplay.WriteString(line2, 2, 0)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Message is longer than 32 characters
		errMsg := "Message too long"
		fmt.Println(errMsg)
		return errMsg
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
