package main

import (
	"log"
	"time"

	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

const (
	width  = 8
	height = 8
	leds   = width * height
)

func main() {
	opt := ws2811.DefaultOptions
	opt.Channels[0].Brightness = 255
	opt.Channels[0].LedCount = leds

	dev, err := ws2811.MakeWS2811(&opt)
	if err != nil {
		log.Fatal(err)
	}

	if err := dev.Init(); err != nil {
		log.Fatal(err)
	}
	defer dev.Fini()

	// Turn on all LEDs
	for i := 0; i < leds; i++ {
		dev.Leds(0)[i] = 0xFFFFFF // White color
	}

	if err := dev.Render(); err != nil {
		log.Fatal(err)
	}

	// Keep the LEDs on for 5 seconds
	time.Sleep(5 * time.Second)

	// Turn off all LEDs
	for i := 0; i < leds; i++ {
		dev.Leds(0)[i] = 0x000000 // Off
	}

	if err := dev.Render(); err != nil {
		log.Fatal(err)
	}
}
