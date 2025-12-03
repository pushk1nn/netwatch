package main

import (
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/pushk1nn/netwatch/internal"
)

// listen will initialize the interface to be used for packet captures,
// then start the loop to handle packets.
func listen(dev string) {
	if handle, err := pcap.OpenLive(dev, 1600, true, pcap.BlockForever); err != nil {
		panic(err)
	} else if err := handle.SetBPFFilter("tcp src port 22"); err != nil { // For early development, listen for ssh
		// TODO: make the filter string assembled by stringing together filters by "and"
		// to make it modular.
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

		for packet := range packetSource.Packets() {
			// get TCP layer of packet
			if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
				tcp, _ := tcpLayer.(*layers.TCP)
				// get IPv4 layer of packet
				if ipv4Layer := packet.Layer(layers.LayerTypeIPv4); ipv4Layer != nil {
					// get IP address of host receiving the connection from server
					ipv4, _ := ipv4Layer.(*layers.IPv4)
					ip := ipv4.DstIP.String()

					if tcp.SYN && tcp.ACK { // Indicates start of SSH session
						internal.Start(packet, ip)
					} else if tcp.FIN && tcp.ACK { // Indicates end of SSH session
						internal.End(packet, ip)
					}
				}
			}
		}
	}
}

func main() {
	// host device to listen on
	dev := os.Args[1]

	listen(dev)
}
