package api

func (l *Api) Router() {
	l.gin.POST("/v1/fans", l.SpeedControlLogic)
	l.gin.GET("/v1/status", l.AutoControlStatus)
}
