package temp

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
)

func CpuTemperature() (temp []int, err error) {
	cpuInfo, err := host.SensorsTemperatures()
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	var cpuTemperature []int
	for _, cpu := range cpuInfo {
		cpuTemperature = append(cpuTemperature, int(cpu.Temperature))
	}
	return cpuTemperature, nil
}
