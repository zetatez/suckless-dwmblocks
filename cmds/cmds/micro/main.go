package main

import (
	"fmt"

	"cmds/sugar"
)

func main() {
	fmt.Println(FormatMicro())
}

func FormatMicro() (str string) {
	status, percent, err := sugar.GetMicro()
	if err != nil {
		return ""
	}
	statusIcons := map[string]string{
		"on":  "󰍬",
		"off": "󰍭",
	}
	str = fmt.Sprintf("%s %02.0f", statusIcons[status], percent)
	return str
}
