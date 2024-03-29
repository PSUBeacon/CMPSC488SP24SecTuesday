package gocode

func BuzzerStatus(pin uint8, status bool) {
	Switchable(pin, status)
}
