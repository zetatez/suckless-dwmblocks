package main

import (
	"fmt"

	"cmds/sugar"
)

func main() {
	fmt.Println(FormatClock())
}

func FormatClock() (str string) {
	return sugar.GetClock()
}
