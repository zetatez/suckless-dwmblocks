package main

import (
	"fmt"

	"cmds/sugar"
)

func main() {
	fmt.Println(FormatCpu())
}

func FormatCpu() (str string) {
	cpuPercent, err := sugar.GetCpuPercent()
	if err != nil {
		return ""
	}
	icons := map[string]string{
		"10":  "",
		"25":  "󰖃",
		"50":  "󰜎",
		"75":  "󰑮",
		"100": "󱄟",
	}
	icon := ""
	switch {
	case cpuPercent < 10:
		icon = icons["10"]
	case cpuPercent < 25:
		icon = icons["25"]
	case cpuPercent < 50:
		icon = icons["50"]
	case cpuPercent < 75:
		icon = icons["75"]
	case cpuPercent < 100:
		icon = icons["100"]
	}
	// str = fmt.Sprintf("cpu %s", icon)
	str = fmt.Sprintf("%s", icon)
	return str
}
