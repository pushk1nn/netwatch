all:
	go build -o netwatch cmd/main.go && sudo setcap cap_net_raw,cap_net_admin=eip ./netwatch
