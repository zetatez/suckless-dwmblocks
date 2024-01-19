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
	str = fmt.Sprintf("c: %.2f%%", cpuPercent)
	return str
}
