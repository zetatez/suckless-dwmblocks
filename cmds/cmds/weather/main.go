package main

import (
	"fmt"

	"cmds/sugar"
)

func main() {
	fmt.Println(FormatWeather())
}

func FormatWeather() (str string) {
	temp, _, err := sugar.GetWeather()
	if err != nil {
		return ""
	}
	str = fmt.Sprintf(" %s", temp)
	return str
}
