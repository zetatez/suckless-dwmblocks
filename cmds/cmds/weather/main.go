package main

import (
	"fmt"

	"cmds/sugar"
)

func main() {
	fmt.Println(FormatWeather())
}

func FormatWeather() (str string) {
	weather, err := sugar.GetWeather()
	if err != nil {
		return ""
	}
	str = weather
	return str
}
