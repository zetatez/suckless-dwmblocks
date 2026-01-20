package blocks

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var (
	microIcons = map[string]string{
		"on":  "",
		"off": "",
	}
	microPercentRe = regexp.MustCompile(`\[(\d+)%\]`)
)

func BlockMicro() string {
	status, pct, err := GetMicro()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s %02.0f", microIcons[status], pct)
}

func GetMicro() (status string, percent float64, err error) {
	stdout, err := exec.Command("amixer", "get", "Capture").CombinedOutput()
	if err != nil {
		return "", 0, err
	}
	out := string(stdout)
	status = "on"
	if strings.Contains(out, "[off]") {
		status = "off"
	}
	xs := microPercentRe.FindAllStringSubmatch(out, -1)
	if len(xs) == 0 {
		return status, 0, fmt.Errorf("get micro failed")
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
		return status, 0, fmt.Errorf("get micro failed")
	}
	percent = sum / cnt
	return status, percent, nil
}
