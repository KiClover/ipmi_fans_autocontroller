package api

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"serverTemperature/model"
)

type Api struct {
	gin  *gin.Engine
	Conf *model.Config
	C    *cache.Cache
}

func New(conf *model.Config, c *cache.Cache) *Api {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	return &Api{
		gin:  r,
		Conf: conf,
		C:    c,
	}
}
func (l *Api) Run(port string) {
	l.Router()
	err := l.gin.Run(port)
	if err != nil {
		logrus.Warn("Gin server start error: %v", err)
		return
	}
	logrus.Infof("Gin server start in port: %s", port)
}
