package blocks

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func BlockNet() string {
	cfg := getConfig()
	state, err := GetNetState(cfg.NetInterface)
	if err != nil || state == "down" {
		return "󰞃"
	}
	signal, _ := GetWirelessSignalPercent(cfg.NetInterface)
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
	file, err := os.Open("/proc/net/wireless")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, iface) {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue
		}
		link := fields[3]
		link = strings.TrimSuffix(link, ".")
		signal, err := strconv.ParseFloat(link, 64)
		if err != nil {
			return 0, err
		}
		return signal, nil
	}
	return 0, fmt.Errorf("wireless interface %s not found", iface)
}
