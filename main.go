package main

import (
	"fmt"
	"serverTemperature/temp"
	"time"
)

func main() {
	for {
		cpuTemp, err := temp.CpuTemperature()
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		fmt.Printf("%v\n", cpuTemp)

		time.Sleep(2 * time.Second)
	}
}
