package main

import (
	"fmt"

	"cmds/sugar"
)

func main() {
	fmt.Println(FormatScreenLight())
}

func FormatScreenLight() (str string) {
	screenLight, err := sugar.GetScreenLight()
	if err != nil {
		return ""
	}
	str = fmt.Sprintf("󱩒  %02.0f%%", screenLight)
	return str
}
