package internal

import (
	"fmt"
	"time"

	"github.com/google/gopacket"
)

// ActiveConnections is a map of host IPs, mapped to UUIDs representing
// different connections. This is done to handle a case in which one remote
// host has several open connections
var ActiveConnections map[string]*Connection

var loc *time.Location

type Connection struct {
	protocol string
	ip       string
	start    time.Time
	end      time.Time
	count    int
}

func init() {
	ActiveConnections = make(map[string]*Connection)

	// Load the Eastern Time Zone location
	loc, _ = time.LoadLocation("America/New_York")
}

// Start is called by main when the host sends a SYN/ACK from port 22,
// signaling a new SSH connection
func Start(packet gopacket.Packet, ip string) {
	connection, ok := ActiveConnections[ip]
	if ok {
		connection.count += 1
	} else {
		connection = &Connection{
			protocol: "SSH",
			ip:       ip,
			start:    packet.Metadata().Timestamp.In(loc),
		}

		ActiveConnections[ip] = connection
	}
}

func (c *Connection) End(packet gopacket.Packet) {
	c.end = packet.Metadata().Timestamp.In(loc)
}

func (c *Connection) Print() {
	fmt.Printf("Connection from %s started at: %s, ended at %s. %d connection(s) from host are still active\n", c.ip, c.start.Format(time.DateTime), c.end.Format(time.DateTime), c.count)
}
