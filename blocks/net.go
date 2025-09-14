package blocks

import (
	"dwmblocks/utils"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

const NetPath = "/sys/class/net/wlan0"

func BlockNet() string {
	state, err := GetNetState("wlan0")
	if err != nil || state == "down" {
		return "󰞃"
	}
	_, signal := GetActiveWifi()
	signalIcon := map[string]string{
		"90": "󰤨",
		"75": "󰤥",
		"50": "󰤢",
		"25": "󰤟",
		"5":  "󰤯",
		"0":  "󰤮",
	}
	icon := "󰤮"
	switch {
	case signal >= 90:
		icon = signalIcon["90"]
	case signal >= 75:
		icon = signalIcon["75"]
	case signal >= 50:
		icon = signalIcon["50"]
	case signal >= 25:
		icon = signalIcon["25"]
	case signal >= 5:
		icon = signalIcon["5"]
	default:
		icon = signalIcon["0"]
	}
	return fmt.Sprintf("%s ", icon)
}

func GetNetState(iface string) (string, error) {
	filePath := path.Join("/sys/class/net", iface, "operstate")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func GetActiveWifi() (ssid string, signal float64) {
	stdout, _, err := utils.RunScript("bash", "nmcli -t -f ACTIVE,SSID,SIGNAL device wifi")
	if err != nil {
		return "", 0.0
	}
	lines := strings.Split(string(stdout), "\n")
	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) == 3 && (fields[0] == "yes" || fields[0] == "是") {
			ssid = fields[1]
			signalInt64, _ := strconv.Atoi(fields[2])
			signal := float64(signalInt64)
			return ssid, signal
		}
	}
	return "", 0.0
}
