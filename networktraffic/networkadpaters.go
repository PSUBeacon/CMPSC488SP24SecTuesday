package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	// Get a list of available network interfaces
	interfaces, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// Print information about each network interface
	fmt.Println("List of available network interfaces:")
	for _, iface := range interfaces {
		fmt.Printf("Name: %s\n", iface.Name)
		fmt.Printf("Description: %s\n", iface.Description)
		fmt.Printf("Addresses:\n")
		for _, addr := range iface.Addresses {
			fmt.Printf("  IP: %s\n", addr.IP)
			fmt.Printf("  Netmask: %s\n", addr.Netmask)
		}
		fmt.Println("-----------------------------------")
	}
}
