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

	str = fmt.Sprintf("d: %.2f%%", diskPercent)
	return str
}
