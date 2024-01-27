package main

import (
	"fmt"

	"cmds/sugar"
)

func main() {
	fmt.Println(FormatVolume())
}

func FormatVolume() (str string) {
	status, percent, err := sugar.GetVolume()
	if err != nil {
		return ""
	}
	statusIcons := map[string]string{
		"on":  "󰕾",
		"off": "󰖁",
	}
	str = fmt.Sprintf("%s %02.0f", statusIcons[status], percent)
	return str
}
