package temp

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
)

func CpuTemperature() ([]float64, error) {
	cpuInfo, err := host.SensorsTemperatures()
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	var cpuTemperature []float64
	for _, cpu := range cpuInfo {
		cpuTemperature = append(cpuTemperature, cpu.Temperature)
	}
	return cpuTemperature, nil
}

func CheckTempLevel() (int, error) {
	return 0, nil
}
