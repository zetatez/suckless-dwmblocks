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
	ipAddr, _ := sugar.GetLocalIpv4ByInterfaceName("wlan0")
	return fmt.Sprintf("%s %s ", ipAddr, state)
}
