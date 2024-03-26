package gocode

func FanStatus(pin uint8, status bool) {
	Switchable(pin, status)
}
