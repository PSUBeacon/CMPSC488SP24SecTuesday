package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"go.bug.st/serial"
	"log"
	"time"
)

type PacketDetails struct {
	NetworkType     string `json:"network_type"`
	SourceMAC       string `json:"source_mac"`
	DestinationMAC  string `json:"destination_mac"`
	SourceIP        string `json:"source_ip"`
	DestinationIP   string `json:"destination_ip"`
	SourcePort      string `json:"source_port"`
	DestinationPort string `json:"destination_port"`
}

func processZigbeeFrames(data []byte) {
	// Ensure the minimum length for a Zigbee frame (adjust as needed)
	if len(data) < 15 {
		fmt.Println("Invalid Zigbee frame length")
		return
	}

	// Extract MAC address (64-bit) from the Zigbee frame
	mac := data[1:9]

	// Extract PAN ID (16-bit) from the Zigbee frame
	panID := data[11:13]

	// Print the parsed MAC address and PAN ID
	fmt.Printf("MAC Address: %X\n", mac)
	fmt.Printf("PAN ID: %X\n", panID)
}

func getNetEvents() {
	// list all adapters
	networkAdapters()

	// Define the network interface you want to capture packets from
	device := "any"

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

		// Extract link layer information (Ethernet header)
		ethernetLayer := packet.LinkLayer()
		if ethernetLayer == nil {
			// Extract source and destination MAC addresses
			continue

			//// Store MAC addresses in the map
			//macSet[srcMAC] = true
			//macSet[dstMAC] = true
		}

		srcMAC := ethernetLayer.LinkFlow().Src().String()
		dstMAC := ethernetLayer.LinkFlow().Dst().String()

		// Extract network layer information (IP header)
		networkLayer := packet.NetworkLayer()
		if networkLayer == nil {
			// ignore if not present
			continue

			//// Store IP addresses in the map
			//ipSet[srcIP] = true
			//ipSet[dstIP] = true
		}

		srcIP := networkLayer.NetworkFlow().Src().String()
		dstIP := networkLayer.NetworkFlow().Dst().String()

		// Get the transport layer
		transportLayer := packet.TransportLayer()
		if transportLayer == nil {
			// Get source and destination ports
			continue
		}

		srcPort := transportLayer.TransportFlow().Src().String()
		dstPort := transportLayer.TransportFlow().Dst().String()

		packetDetails := PacketDetails{
			NetworkType:     "tcp/ip",
			SourceMAC:       srcMAC,
			DestinationMAC:  dstMAC,
			SourceIP:        srcIP,
			DestinationIP:   dstIP,
			SourcePort:      srcPort,
			DestinationPort: dstPort,
		}

		jsonNetworkData, err := json.Marshal(packetDetails)
		if err != nil {
			log.Printf("Error marshalling JSON: %v", err)
			continue
		}

		// print net event data
		fmt.Println(string(jsonNetworkData))

		// Check if the capture duration has elapsed
		if time.Now().After(endTime) {
			fmt.Println("Capture duration reached. Exiting...")
			break
		}
	}

	print("Checking zigbee network traffic ...")

	// Open the XBee module for communication
	port, err := serial.Open("/dev/ttyUSB0", &serial.Mode{BaudRate: 9600})
	if err != nil {
		log.Fatal("Error opening XBee module:", err)
	}

	// Get the current time
	startTime := time.Now()

	for {

		// Read data from the serial port
		buf := make([]byte, 128) // Adjust buffer size as needed
		n, err := port.Read(buf)
		if err != nil {
			log.Println("Error reading from serial port:", err)
			continue
		}

		// Process the received data (Zigbee frames)
		fmt.Println("\nZigbee PAN and MAC on the network:")
		processZigbeeFrames(buf[:n])

		// Check if 30 seconds have passed
		if time.Since(startTime) >= 30*time.Second {
			break
		}

		// Optionally, you can add a small delay to reduce CPU usage
		time.Sleep(100 * time.Millisecond)
	}

}
