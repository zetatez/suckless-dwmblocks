package main

import (
	"fmt"

	"cmds/sugar"
)

func main() {
	fmt.Println(FormatCpuTemperature())
}

func FormatCpuTemperature() (str string) {
	avgTemerature, err := sugar.GetCpuTemperature()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("󰔐 %02.0f󰔄", avgTemerature)
}
