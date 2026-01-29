package blocks

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

func BlockMem() string {
	pct, err := GetMemPercent()
	if err != nil {
		return "󰍛 --"
	}
	return fmt.Sprintf("%s", GetIconByPct(pct))
}

func GetMemPercent() (float64, error) {
	stat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return stat.UsedPercent, nil
}
