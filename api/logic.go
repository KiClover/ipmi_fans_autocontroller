package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"serverTemperature/ipmi"
	"strconv"
)

func (l *Api) SpeedControlLogic(context *gin.Context) {
	s := context.Query("speed")
	speed, _ := strconv.Atoi(s)
	err := ipmi.ControlFansByWeb(speed, *l.Conf, l.C)
	if err != nil {
		logrus.Warn("control fans speed api error: %v", err)
		context.JSON(500, gin.H{
			"message": "control fans speed error",
		})
	}
	logrus.Infof("control fans speed success: %d", speed)
	context.JSON(200, gin.H{
		"message": "control fans speed success",
	})
}

func (l *Api) AutoControlStatus(context *gin.Context) {
	s := context.Query("status")
	if s == "true" {
		logrus.Infof("Enable auto control")
	}
}
