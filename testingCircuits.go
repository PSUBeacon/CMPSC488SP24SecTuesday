package main

import (
	"CMPSC488SP24SecTuesday/on-metal-c-code/gocode"
	"time"
)

func main() {

	// Display a pattern on the matrix

	gocode.MatrixStatus(9, 4, 10, true, 5)
	time.Sleep(5)
	gocode.TurnOffMatrix(9, 4, 10)

	gocode.MatrixStatus(9, 4, 10, true, 5)
	time.Sleep(5)
	gocode.TurnOffMatrix(9, 4, 10)
}
