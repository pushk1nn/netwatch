package main

import (
	"os"

	"github.com/pushk1nn/netwatch/logging"
)

// device gets the environment variable to listen on
var device = os.Getenv("IFACE")

func main() {
	// Start listening for new packets
	logging.Listen(device)
}
