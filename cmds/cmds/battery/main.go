package main

import (
	"fmt"

	"cmds/sugar"
)

const (
	BatteryPath = "/sys/class/power_supply/BAT0"
)

func main() {
	fmt.Println(FormatBattery())
}

func FormatBattery() (str string) {
	capacity, status, err := sugar.GetBattery(BatteryPath)
	if err != nil {
		return ""
	}
	warn := "🔥"
	statusIcons := map[string]string{
		"Full":         "󱊣",
		"Discharging":  "",
		"Charging":     "",
		"Not charging": "",
		"Unknown":      "󰂑",
	}
	switch {
	case capacity < 25:
		str = fmt.Sprintf("%s %02.0f%%", warn, capacity)
	default:
		str = fmt.Sprintf("%s %02.0f%%", statusIcons[status], capacity)
	}
	return str
}
