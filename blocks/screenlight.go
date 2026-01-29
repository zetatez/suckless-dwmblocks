package blocks

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type BacklightInfo struct {
	brightnessPath string
	brightnessMax  float64
	err            error
}

var backlight = sync.OnceValue(func() BacklightInfo {
	const root = "/sys/class/backlight"
	ents, err := os.ReadDir(root)
	if err != nil {
		return BacklightInfo{err: err}
	}
	if len(ents) == 0 {
		return BacklightInfo{err: fmt.Errorf("no backlight device")}
	}
	dev := ents[0].Name()
	m, err := os.ReadFile(filepath.Join(root, dev, "max_brightness"))
	if err != nil {
		return BacklightInfo{err: err}
	}
	maxVal, err := strconv.ParseFloat(strings.TrimSpace(string(m)), 64)
	if err != nil {
		return BacklightInfo{err: err}
	}
	if maxVal <= 0 {
		return BacklightInfo{err: fmt.Errorf("invalid max brightness")}
	}
	return BacklightInfo{
		brightnessPath: filepath.Join(root, dev, "brightness"),
		brightnessMax:  maxVal,
	}
})

func BlockScreenLight() string {
	pct, err := GetScreenLight()
	if err != nil {
		return "󱩒--"
	}
	return fmt.Sprintf("󱩒%02.0f", pct)
}

func GetScreenLight() (float64, error) {
	info := backlight()
	if info.err != nil {
		return 0, info.err
	}
	b, err := os.ReadFile(info.brightnessPath)
	if err != nil {
		return 0, err
	}
	brightness, err := strconv.ParseFloat(strings.TrimSpace(string(b)), 64)
	if err != nil {
		return 0, err
	}
	return (brightness / info.brightnessMax) * 100.0, nil
}
