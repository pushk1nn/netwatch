#!/bin/bash

# Dependencies for packet decoding
sudo apt install libpcap-dev gcc
export CGO_ENABLED=1

# Configure env variables
export IFACE="$1"
echo "[+] Set to listen on interface $1"
