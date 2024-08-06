package main

import (
	"fmt"

	"cmds/sugar"
)

const (
	NetPath = "/sys/class/net/wlan0"
)

func main() {
	fmt.Println(FormatNet())
}

func FormatNet() (str string) {
	operstate, err := sugar.GetNet(NetPath)
	if err != nil {
		return ""
	}
	operstateIcons := map[string]string{
		"up":   "󰒢",
		"down": "󰞃",
	}
	state := operstateIcons[operstate]
	ipAddr, err := sugar.GetLocalIpv4ByInterfaceName("wlan0")
	if err != nil {
		ipAddr = "127.0.0.1"
	}
	_, signal := sugar.GetActiveWifi()
	return fmt.Sprintf("%s %s  %02.0d%%", ipAddr, state, signal)
}
