package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"serverTemperature/model"
	"serverTemperature/temp"
	"time"
)

func main() {
	// 读取配置文件路径
	var path = pflag.String("f", "./config.yaml", "config file")
	pflag.Parse()
	// 读取配置文件
	v := viper.New()
	v.SetConfigFile(*path)
	err := v.ReadInConfig()
	if err != nil {
		logrus.Fatal(fmt.Sprintf("config read fail: %v", err))
		return
	}
	// 配置文件解析
	var conf model.Config
	err = v.Unmarshal(&conf)
	if err != nil {
		logrus.Fatalf("config unmarshal fail: %v", err)
		return
	}
	logrus.SetLevel(logrus.Level(conf.LogLevel))
	e, err := temp.GpuTemperature()
	fmt.Println(e)

	for {
		cpuTemp, err := temp.CpuTemperature()
		if err != nil {
			logrus.Warnf("get cpu temperature err: %v", err)
		}
		logrus.Infof("cpus temperature: %v", cpuTemp)
		gpuTemp, err := temp.GpuTemperature()
		if err != nil {
			logrus.Warnf("get gpu temperature err: %v", err)
		}
		logrus.Infof("gpus temperature: %v", gpuTemp)
		time.Sleep(time.Duration(conf.Monitor.Duration) * time.Second)
	}
}
