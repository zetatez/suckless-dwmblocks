package main

import (
	"fmt"

	"cmds/sugar"
)

func main() {
	fmt.Println(FormatDisk())
}

func FormatDisk() (str string) {
	diskPercent, err := sugar.GetDiskPercent()
	if err != nil {
		return ""
	}

	str = fmt.Sprintf("💽 %.2f%%", diskPercent)
	return str
}
