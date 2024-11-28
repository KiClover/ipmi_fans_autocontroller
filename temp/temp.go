package temp

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"serverTemperature/model"
)

func maxInt(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("arrays empty")
	}

	maxValue := arr[0]
	for _, value := range arr {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue, nil
}

func LevelCheck(temp []int, conf model.Config) (level int, speed int) {
	var levels []int
	for i := 0; i < len(temp); i++ {
		if temp[i] >= conf.TempLevel.Level1.Temp && temp[i] < conf.TempLevel.Level2.Temp {
			levels = append(levels, 1)
		}
		if temp[i] >= conf.TempLevel.Level2.Temp && temp[i] < conf.TempLevel.Level3.Temp {
			levels = append(levels, 2)
		}
		if temp[i] >= conf.TempLevel.Level3.Temp && temp[i] < conf.TempLevel.Level4.Temp {
			levels = append(levels, 3)
		}
		if temp[i] >= conf.TempLevel.Level4.Temp && temp[i] < conf.TempLevel.Level5.Temp {
			levels = append(levels, 4)
		}
		if temp[i] >= conf.TempLevel.Level5.Temp {
			levels = append(levels, 5)
		}
	}
	level, err := maxInt(levels)
	if err != nil {
		logrus.Warnf("comparison function error: %v", err)
		return 3, conf.TempLevel.Level3.Speed
	}
	switch level {
	case 1:
		speed = conf.TempLevel.Level1.Speed
	case 2:
		speed = conf.TempLevel.Level2.Speed
	case 3:
		speed = conf.TempLevel.Level3.Speed
	case 4:
		speed = conf.TempLevel.Level4.Speed
	case 5:
		speed = conf.TempLevel.Level5.Speed
	}
	return level, speed
}
