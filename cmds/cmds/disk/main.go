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
	fmt.Println(FormatDisk())
}

func FormatDisk() (str string) {
	diskPercent, err := sugar.GetDiskPercent()
	if err != nil {
		return ""
	}

	str = fmt.Sprintf("d: %.2f%%", diskPercent)
	return str
}
