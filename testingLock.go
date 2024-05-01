package main

//
//import (
//	"fmt"
//	"log"
//	"os"
//	"os/signal"
//	"syscall"
//	"time"
//
//	"github.com/stianeikeland/go-rpio/v4"
//)
////
////const (
////	servoPinNumber   = 23
////	pulseFrequency   = 20 * time.Millisecond
////	minPulseWidth    = 700 * time.Microsecond
////	maxPulseWidth    = 1500 * time.Microsecond
////	rotationDuration = 1 * time.Second
////)
//
//func main() {
//	if err := rpio.Open(); err != nil {
//		log.Fatal(err)
//	}
//	defer rpio.Close()
//
//	servoPin := rpio.Pin(servoPinNumber)
//	servoPin.Output()
//
//	fmt.Println("Servo control started. Type 'open' to open or 'close' to close. Press Ctrl+C to exit.")
//
//	sigChan := make(chan os.Signal, 1)
//	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
//	go func() {
//		<-sigChan
//		fmt.Println("\nExiting...")
//		os.Exit(0)
//	}()
//
//	for {
//		var command string
//		fmt.Print("Command (open/close): ")
//		fmt.Scanln(&command)
//
//		switch command {
//		case "open":
//			setServoAngle(servoPin, maxPulseWidth) // Move to one side
//		case "close":
//			setServoAngle(servoPin, minPulseWidth) // Move to the opposite side
//		default:
//			fmt.Println("Invalid command. Please type 'open' or 'close'.")
//		}
//	}
//}
//
//func setServoAngle(pin rpio.Pin, pulseWidth time.Duration) {
//	end := time.Now().Add(rotationDuration)
//	for time.Now().Before(end) {
//		// Send the pulse
//		pin.High()
//		time.Sleep(pulseWidth)
//		pin.Low()
//		time.Sleep(pulseFrequency - pulseWidth)
//	}
//}
