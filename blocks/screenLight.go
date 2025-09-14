package blocks

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func BlockScreenLight() string {
	pct, err := GetScreenLight()
	if err != nil {
		return "󱩒 --%"
	}
	return fmt.Sprintf("󱩒 %02.0f%%", pct)
}

func GetScreenLight() (float64, error) {
	stdout, err := exec.Command("light").Output()
	if err != nil {
		return 0, err
	}
	pct, err := strconv.ParseFloat(strings.TrimSpace(string(stdout)), 64)
	if err != nil {
		return 0, err
	}
	return pct, nil
}
