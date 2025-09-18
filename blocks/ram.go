package blocks

import (
	"github.com/shirou/gopsutil/mem"
)

func BlockRAM() string {
	pct, err := GetMemPercent()
	if err != nil {
		return "󰍛 --"
	}
	// return fmt.Sprintf("%s %02.0f", GetIconByPct(pct), pct)
	return GetIconByPct(pct)
}

func GetMemPercent() (float64, error) {
	stat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return stat.UsedPercent, nil
}
