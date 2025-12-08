package internal

import (
	"log"
	"time"

	"github.com/google/gopacket"
)

var loc *time.Location

func init() {

	// Load the Eastern Time Zone location
	loc, _ = time.LoadLocation("America/New_York")
}

// Start is called by main when the host sends a SYN/ACK from port 22,
// signaling a new SSH connection
func Start(packet gopacket.Packet, ip string) {

	time := packet.Metadata().Timestamp.In(loc)

	conn, err := Client.Connections.
		Create().
		SetIP(ip).
		SetTime(time).
		SetUnixTime(time.Unix()).
		SetType("CONNECT").
		Save(Ctx)

	log.Println("connection was logged: ", conn)

	if err != nil {
		log.Printf("failed to log connection: %v", err)
	}

}

func End(packet gopacket.Packet, ip string) {
	time := packet.Metadata().Timestamp.In(loc)

	conn, err := Client.Connections.
		Create().
		SetIP(ip).
		SetTime(time).
		SetUnixTime(time.Unix()).
		SetType("DISCONNECT").
		Save(Ctx)

	log.Println("connection was logged: ", conn)

	if err != nil {
		log.Printf("failed to log connection: %v", err)
	}
}
