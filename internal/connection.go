package internal

import (
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type Connection struct {
	protocol gopacket.LayerType
	start    time.Time
	end      time.Time
}

func Process(packet gopacket.Packet) {
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		if tcp.SYN && tcp.ACK {
			fmt.Println("Connected")
		}
		if tcp.FIN && tcp.ACK {
			fmt.Println("Disconnected")
		}
	}
}
