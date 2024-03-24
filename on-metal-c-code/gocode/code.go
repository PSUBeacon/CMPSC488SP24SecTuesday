package gocode

// #include Fan.h
//
//
import "C"
import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
)

func setPins() {
	// Open /dev/mem using the gpio driver
	err := rpio.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err := rpio.Close()
		if err != nil {

		}
	}()
	//
	//// Define pin numbers
	//DHT_PIN_DATA := 4
	//BUZZER_PIN_SIG := 4
	//KEYPADMEM3X4_PIN_ROW1 := 22
	//KEYPADMEM3X4_PIN_ROW2 := 23
	//KEYPADMEM3X4_PIN_ROW3 := 24
	//KEYPADMEM3X4_PIN_ROW4 := 25
	//KEYPADMEM3X4_PIN_COL1 := 17
	//KEYPADMEM3X4_PIN_COL2 := 18
	//KEYPADMEM3X4_PIN_COL3 := 27
	//LEDMATRIX_PIN_DIN := 9
	//LEDMATRIX_PIN_CLK := 10
	//LEDMATRIX_PIN_CS := 4
	//PIR_PIN_SIG := 11
	//PCFAN_PIN_COIL1 := 18

	// Your GPIO operations here
}
