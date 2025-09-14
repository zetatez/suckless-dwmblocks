package blocks

import (
	"dwmblocks/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func BlockVol() string {
	status, pct, err := GetVolume()
	if err != nil {
		return "?"
	}
	statusIcons := map[string]string{
		"on":  "󰕾",
		"off": "󰖁",
	}
	return fmt.Sprintf("%s %02.0f%%", statusIcons[status], pct)
}

func GetVolume() (status string, percent float64, err error) {
	stdout, _, err := utils.RunScript("sh", "amixer get Master")
	if err != nil {
		return "", 0.0, err
	}
	status = "on"
	if strings.Contains(string(stdout), "[off]") {
		status = "off"
	}
	r := regexp.MustCompile(`\[(?P<percent>\d+)%\]`)
	xs := r.FindAllStringSubmatch(string(stdout), -1)
	if len(xs) == 0 {
		return status, 0.0, fmt.Errorf("get volume failed")
	}
	sum, cnt := 0.0, 0.0
	for _, x := range xs {
		p, _ := strconv.ParseFloat(x[1], 64)
		sum += p
		cnt++
	}
	percent = sum / cnt
	return status, percent, nil
}
