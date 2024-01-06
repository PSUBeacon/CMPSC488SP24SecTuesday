package main

//go get github.com/google/gopacket
//go get github.com/google/gopacket/pcap

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Define the network interface you want to capture packets from
	device := "\\Device\\NPF_{285EEC32-9286-4AB7-83F9-0C2E8BDE60FC}"

	// Open the network device for packet capture
	handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Create a packet source to decode packets from the network interface
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Capture packets for a specified duration (e.g., 30 seconds)
	duration := 30 * time.Second
	endTime := time.Now().Add(duration)

	fmt.Printf("Capturing network traffic on interface %s for %s...\n", device, duration)

	// Loop to capture and analyze packets
	for packet := range packetSource.Packets() {
		// Print basic packet information
		fmt.Printf("Packet captured at %s\n", packet.Metadata().Timestamp)
		fmt.Printf("Packet length: %d bytes\n", packet.Metadata().Length)
		fmt.Printf("Packet data:\n%s\n", packet.Dump())

		// You can perform more advanced analysis on the packet data here

		// Check if the capture duration has elapsed
		if time.Now().After(endTime) {
			fmt.Println("Capture duration reached. Exiting...")
			break
		}
	}
}
