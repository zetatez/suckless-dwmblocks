package main

import (
	"fmt"

	"cmds/sugar"
)

const (
	BatteryPath = "/sys/class/power_supply/BAT0"
	NetPath     = "/sys/class/net/wlan0"
	EmailPath   = ".mail/inbox"
	MsgPath     = ".msg"
)

func main() {
	fmt.Println(FormatCpu())
}

func FormatCpu() (str string) {
	cpuPercent, err := sugar.GetCpuPercent()
	if err != nil {
		return ""
	}
	str = fmt.Sprintf("c: %.2f%%", cpuPercent)
	return str
}