package temp

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

func GpuTemperature() (temp []int, err error) {
	cmd := exec.Command("nvidia-smi", "--query-gpu=temperature.gpu", "--format=csv,noheader,nounits")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return nil, err
	}
	output := out.String()
	lines := strings.Split(strings.TrimSpace(output), "\n")
	var temps []int
	for _, line := range lines {
		tempAll, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		temps = append(temps, tempAll)
	}
	return temps, nil
}
