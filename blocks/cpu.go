package blocks

import (
	"github.com/shirou/gopsutil/cpu"
)

func BlockCPU() string {
	pct, err := GetCpuPercent()
	if err != nil {
		return "CPU: ?"
	}
	// return fmt.Sprintf("%s %02.0f", GetIconByPct(pct), pct)
	return GetIconByPct(pct)
}

func GetCpuPercent() (float64, error) {
	percents, err := cpu.Percent(0, false)
	if err != nil || len(percents) == 0 {
		return 0, err
	}
	return percents[0], nil
}
