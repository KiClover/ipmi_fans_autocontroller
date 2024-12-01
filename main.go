package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"serverTemperature/ipmi"
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
	// 初始化KV缓存数据库
	c := cache.New(600*time.Second, 60*time.Second)
	_ = ipmi.WebLogin(conf, c)
	go ipmi.RefreshClock(conf, c)
	for {
		// 获取温度
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
		// 获取需求风扇转速与自动控制转速等级
		level, speed := temp.LevelCheck(append(cpuTemp, gpuTemp...), conf)
		logrus.Infof("controller level: %d , fans speed: %d", level, speed)
		lv, isLevel := c.Get("level")
		if isLevel {
			if lv != level {
				err := ipmi.ControlFansByWeb(speed, conf, c)
				if err != nil {
					logrus.Warnf("control fans speed error: %v", err)
				}
			}
			c.Set("level", level, 0)
		} else {
			err = c.Add("level", level, 0)
			if err != nil {
				logrus.Warnf("add cache error: %v", err)
			}
			err := ipmi.ControlFansByWeb(speed, conf, c)
			if err != nil {
				logrus.Warnf("control fans speed error: %v", err)
			}
		}
		// 监测间隔控制
		time.Sleep(time.Duration(conf.Monitor.Duration) * time.Second)
	}
}
