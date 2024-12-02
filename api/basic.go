package api

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"serverTemperature/model"
)

type Api struct {
	gin  *gin.Engine
	Conf model.Config
	C    *cache.Cache
}

func (l Api) New() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	l.gin = r
	l.Router()
	l.Run(":8080")
}
func (l Api) Run(port string) {
	err := l.gin.Run(port)
	if err != nil {
		logrus.Warn("Gin server start error: %v", err)
		return
	}
	logrus.Infof("Gin server start in port: %s", port)
}
