package main

import (
	"fmt"

	"cmds/sugar"
)

func main() {
	fmt.Println(FormatRAM())
}

func FormatRAM() (str string) {
	memPercent, err := sugar.GetMemPercent()
	if err != nil {
		return ""
	}
	str = fmt.Sprintf("%02.0f", memPercent)
	return str
}

// func FormatRAM() (str string) {
// 	memPercent, err := sugar.GetMemPercent()
// 	if err != nil {
// 		return ""
// 	}
// 	icons := map[string]string{
// 		"10":  "",
// 		"25":  "󰖃",
// 		"50":  "󰜎",
// 		"75":  "󰑮",
// 		"100": "󱄟",
// 	}
// 	icon := ""
// 	switch {
// 	case memPercent < 10:
// 		icon = icons["10"]
// 	case memPercent < 25:
// 		icon = icons["25"]
// 	case memPercent < 50:
// 		icon = icons["50"]
// 	case memPercent < 75:
// 		icon = icons["75"]
// 	case memPercent < 100:
// 		icon = icons["100"]
// 	}
// 	// str = fmt.Sprintf("ram %s", icon)
// 	str = fmt.Sprintf("%s", icon)
// 	return str
// }
