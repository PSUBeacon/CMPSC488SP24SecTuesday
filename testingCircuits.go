package main

import (
	"fmt"
	"log"

	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

func main() {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Open a handle to the SPI device:
	port, err := spireg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	// Communicate with the device:
	s, err := port.Connect(100, spi.Mode0, 8)
	if err != nil {
		log.Fatal(err)
	}

	// Send data to turn on all LEDs:
	for i := 0; i < 8; i++ {
		if err := s.Tx([]byte{0x01 << uint(i), 0xFF}, nil); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("All LEDs turned on.")
}
