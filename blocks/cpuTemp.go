package blocks

import (
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/host"
)

func BlockCPUTemp() string {
	temp, err := GetCpuMaxTemp()
	if err != nil || temp == 0 {
		return "❄ --󰔄"
	}
	return fmt.Sprintf("❄ %02.0f󰔄", temp)
}

func GetCpuMaxTemp() (maxTemp float64, err error) {
	sensors, _ := host.SensorsTemperatures()
	if sensors == nil {
		return 0, nil
	}
	maxTemp = 0.0
	for _, sensor := range sensors {
		if strings.HasPrefix(sensor.SensorKey, "coretemp_core") && strings.HasSuffix(sensor.SensorKey, "_input") {
			if sensor.Temperature > maxTemp {
				maxTemp = sensor.Temperature
			}
		}
	}
	return maxTemp, nil
}
