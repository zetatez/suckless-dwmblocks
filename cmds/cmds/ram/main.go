package main

import (
	"fmt"

	"cmds/sugar"
)

func main() {
	fmt.Println(FormatRam())
}

func FormatRam() (str string) {
	ramPercent, err := sugar.GetRamPercent()
	if err != nil {
		return ""
	}
	str = fmt.Sprintf("r: %.2f%%", ramPercent)
	return str
}
