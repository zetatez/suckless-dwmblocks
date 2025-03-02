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
	// ipAddr, err := sugar.GetLocalIpv4ByInterfaceName("wlan0")
	// if err != nil {
	// 	ipAddr = "127.0.0.1"
	// }
	_, signal := sugar.GetActiveWifi()
	sigIntens := map[string]string{
		"90": "󰤨",
		"75": "󰤥",
		"50": "󰤢",
		"25": "󰤟",
		"5":  "󰤯",
		"0":  "󰤮",
	}
	intens := "x"
	switch {
	case signal >= 90:
		intens = sigIntens["90"]
	case signal >= 75:
		intens = sigIntens["75"]
	case signal >= 50:
		intens = sigIntens["50"]
	case signal >= 25:
		intens = sigIntens["25"]
	case signal >= 5:
		intens = sigIntens["5"]
	default:
		intens = sigIntens["0"]
	}
	// return fmt.Sprintf("%s %s %s ", ipAddr, intens, state)
	return fmt.Sprintf("%s %s ", intens, state)
}
