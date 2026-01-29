package blocks

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var (
	volIcons = map[string]string{
		"on":  "󰋋",
		"off": "󰟎",
	}
	volPercentRe = regexp.MustCompile(`\[(\d+)%\]`)
)

func BlockVol() string {
	status, pct, err := GetVolume()
	if err != nil {
		return "?"
	}
	return fmt.Sprintf("%s %02.0f", volIcons[status], pct)
}

func GetVolume() (status string, percent float64, err error) {
	stdout, err := exec.Command("amixer", "get", "Master").CombinedOutput()
	if err != nil {
		return "", 0.0, err
	}
	out := string(stdout)
	status = "on"
	if strings.Contains(out, "[off]") {
		status = "off"
	}
	xs := volPercentRe.FindAllStringSubmatch(out, -1)
	if len(xs) == 0 {
		return status, 0.0, fmt.Errorf("get volume failed")
	}
	sum, cnt := 0.0, 0.0
	for _, x := range xs {
		p, err := strconv.ParseFloat(x[1], 64)
		if err != nil {
			continue
		}
		sum += p
		cnt++
	}
	if cnt == 0 {
		return status, 0.0, fmt.Errorf("get volume failed")
	}
	percent = sum / cnt
	return status, percent, nil
}
