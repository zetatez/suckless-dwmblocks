package blocks

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func BlockScreenLight() string {
	pct, err := GetScreenLight()
	if err != nil {
		return "󱩒 --"
	}
	return fmt.Sprintf("󱩒 %02.0f", pct)
}

func GetScreenLight() (float64, error) {
	ents, err := os.ReadDir("/sys/class/backlight")
	if err != nil {
		return 0, err
	}
	if len(ents) == 0 {
		return 0, fmt.Errorf("no backlight device")
	}
	dev := ents[0].Name()
	brightnessPath := filepath.Join("/sys/class/backlight", dev, "brightness")
	maxPath := filepath.Join("/sys/class/backlight", dev, "max_brightness")

	b, err := os.ReadFile(brightnessPath)
	if err != nil {
		return 0, err
	}
	m, err := os.ReadFile(maxPath)
	if err != nil {
		return 0, err
	}
	brightness, err := strconv.ParseFloat(strings.TrimSpace(string(b)), 64)
	if err != nil {
		return 0, err
	}
	maxBrightness, err := strconv.ParseFloat(strings.TrimSpace(string(m)), 64)
	if err != nil {
		return 0, err
	}
	if maxBrightness <= 0 {
		return 0, fmt.Errorf("invalid max brightness")
	}
	return (brightness / maxBrightness) * 100.0, nil
}
