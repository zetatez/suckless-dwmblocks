package main

import (
	"fmt"

	"cmds/sugar"
)

func main() {
	fmt.Println(FormatWifiIntensity())
}

func FormatWifiIntensity() (str string) {
	_, signal := sugar.GetActiveWifi()
	return fmt.Sprintf("%02.0d%%", signal)
}
