package blocks

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

func BlockNet() string {
	state, err := GetNetState("wlan0")
	if err != nil || state == "down" {
		return "󰞃"
	}
	signal, _ := GetWirelessSignalPercent("wlan0")
	icon := "󰤮"
	switch {
	case signal >= 90:
		icon = "󰤨"
	case signal >= 75:
		icon = "󰤥"
	case signal >= 50:
		icon = "󰤢"
	case signal >= 25:
		icon = "󰤟"
	case signal >= 5:
		icon = "󰤯"
	default:
		icon = "󰤮"
	}
	return fmt.Sprintf("%s", icon)
}

func GetNetState(iface string) (string, error) {
	filePath := path.Join("/sys/class/net", iface, "operstate")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func GetWirelessSignalPercent(iface string) (float64, error) {
	out, err := exec.Command("nmcli", "-t", "-f", "IN-USE,SIGNAL", "dev", "wifi", "list", "ifname", iface, "--rescan", "no").CombinedOutput()
	if err != nil {
		return 0, err
	}

	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// Expected format: "*:57" for the currently used AP.
		if !strings.HasPrefix(line, "*:") {
			continue
		}
		vStr := strings.TrimPrefix(line, "*:")
		v, err := strconv.ParseFloat(vStr, 64)
		if err != nil {
			return 0, err
		}
		if v < 0 {
			v = 0
		}
		if v > 100 {
			v = 100
		}
		return v, nil
	}

	return 0, fmt.Errorf("nmcli: active wifi not found")
}
