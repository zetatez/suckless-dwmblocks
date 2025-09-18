package blocks

import (
	"dwmblocks/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func BlockMicro() string {
	status, pct, err := GetMicro()
	if err != nil {
		return ""
	}
	statusIcons := map[string]string{
		"on":  "󰍬",
		"off": "󰍭",
	}
	return fmt.Sprintf("%s %02.0f", statusIcons[status], pct)
}

func GetMicro() (status string, percent float64, err error) {
	stdout, _, err := utils.RunScript("sh", "amixer get Capture")
	if err != nil {
		return "", 0, err
	}
	status = "on"
	if strings.Contains(string(stdout), "[off]") {
		status = "off"
	}
	r := regexp.MustCompile(`\[(?P<percent>\d+)%\]`)
	xs := r.FindAllStringSubmatch(string(stdout), -1)
	if len(xs) == 0 {
		return status, 0, fmt.Errorf("get micro failed")
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
