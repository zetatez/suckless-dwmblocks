package blocks

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

func BlockDisk() string {
	diskPercent, err := GetDiskPercent()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("💽%02.0f", diskPercent)
}

func GetDiskPercent() (percent float64, err error) {
	parts, err := disk.Partitions(false)
	if err != nil {
		return 0, err
	}
	diskInfo, err := disk.Usage(parts[0].Mountpoint)
	if err != nil {
		return 0, err
	}
	percent = diskInfo.UsedPercent
	return percent, nil
}
