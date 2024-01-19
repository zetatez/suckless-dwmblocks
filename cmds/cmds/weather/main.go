package main

import (
	"fmt"

	"cmds/sugar"
)

func main() {
	fmt.Println(FormatWeather())
}

func FormatWeather() (str string) {
	temp, wind, err := sugar.GetWeather()
	if err != nil {
		return ""
	}
	str = fmt.Sprintf("[weather]: %s %s", temp, wind)
	return str
}
