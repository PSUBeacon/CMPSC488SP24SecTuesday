package cCode

/*
#cgo LDFLAGS: -lwiringPi
#include "MaxMatrix.h"
*/
import "C"

type MaxMatrix struct {
	dataPin  uint8
	loadPin  uint8
	clockPin uint8
	num      uint8
}

func NewMaxMatrix(dataPin, loadPin, clockPin, num uint8) *MaxMatrix {
	return &MaxMatrix{
		dataPin:  dataPin,
		loadPin:  loadPin,
		clockPin: clockPin,
		num:      num,
	}
}

func (m *MaxMatrix) Init(maxInUse uint8) {
	C.MaxMatrix_init(C.uint8_t(m.dataPin), C.uint8_t(m.loadPin), C.uint8_t(m.clockPin), C.uint8_t(m.num), C.uint8_t(maxInUse))
}

func (m *MaxMatrix) SetIntensity(intensity uint8) {
	C.MaxMatrix_setIntensity(C.uint8_t(intensity))
}

func (m *MaxMatrix) Clear() {
	C.MaxMatrix_clear()
}
