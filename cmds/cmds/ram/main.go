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
	str = fmt.Sprintf("%s ", sugar.GetIconByPct(memPercent))
	return str
}

// func FormatRAM() (str string) {
// 	memPercent, err := sugar.GetMemPercent()
// 	if err != nil {
// 		return ""
// 	}
// 	str = fmt.Sprintf("󰍛  %02.0f%%", memPercent)
// 	return str
// }
