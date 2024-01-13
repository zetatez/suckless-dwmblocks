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
		"up":   "󰖩",
		"down": "󰖪",
	}
	str = operstateIcons[operstate]
	return str
}
