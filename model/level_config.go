package model

// TempLevel CPU温度等级
type TempLevel struct {
	Level1 Level1FansSpeedConfig
	Level2 Level2FansSpeedConfig
	Level3 Level3FansSpeedConfig
	Level4 Level4FansSpeedConfig
	Level5 Level4FansSpeedConfig
}

type Level1FansSpeedConfig struct {
	Temp  float64
	Speed int
}

type Level2FansSpeedConfig struct {
	Temp  float64
	Speed int
}

type Level3FansSpeedConfig struct {
	Temp  float64
	Speed int
}

type Level4FansSpeedConfig struct {
	Temp  float64
	Speed int
}

type Level5FansSpeedConfig struct {
	Temp  float64
	Speed int
}
