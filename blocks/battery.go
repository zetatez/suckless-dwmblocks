package blocks

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	BatteryPath     = "/sys/class/power_supply/BAT0"
	LowBatteryWarn  = 25
	CriticalBattery = 10
)

var statusIcons = map[string]string{
	"Full":         "󰂅",
	"Discharging":  "",
	"Charging":     "",
	"Not charging": "󰢟",
	"Unknown":      "󰂑",
	"Warning":      "",
}

func BlockBattery() string {
	capacity, status, err := GetBattery(BatteryPath)
	if err != nil {
		return ""
	}

	return formatBattery(capacity, status)
}

func formatBattery(capacity float64, status string) string {
	switch {
	case status == "Charging":
		return fmt.Sprintf("%s %02.0f", statusIcons[status], capacity)
	case capacity <= CriticalBattery:
		return fmt.Sprintf("%s %02.0f", statusIcons["Warning"], capacity)
	case capacity < LowBatteryWarn:
		return fmt.Sprintf("%s %02.0f", statusIcons["Warning"], capacity)
	case capacity == 100:
		return fmt.Sprintf("%s %02.0f", statusIcons["Full"], capacity)
	default:
		icon, ok := statusIcons[status]
		if !ok {
			icon = statusIcons["Unknown"]
		}
		return fmt.Sprintf("%s %02.0f", icon, capacity)
	}
}

func GetBattery(batteryPath string) (capacity float64, status string, err error) {
	readFile := func(filename string) (string, error) {
		data, err := os.ReadFile(path.Join(batteryPath, filename))
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(data)), nil
	}

	capStr, err := readFile("capacity")
	if err != nil {
		return 0, "", err
	}
	capacity, err = strconv.ParseFloat(capStr, 64)
	if err != nil {
		return 0, "", err
	}

	status, err = readFile("status")
	if err != nil {
		return 0, "", err
	}

	return capacity, status, nil
}
