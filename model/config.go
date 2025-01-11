package model

type Config struct {
	LogLevel  int
	TempLevel TempLevel
	Monitor   MonitorConfig
	Ipmi      IpmiConfig
	WebEnable bool `json:"webEnable"`
}
