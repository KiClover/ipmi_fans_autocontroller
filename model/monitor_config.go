package model

type MonitorConfig struct {
	// 监控间隔，单位为秒
	Duration int
}

type IpmiConfig struct {
	Host string
	User string
	Pwd  string
}
