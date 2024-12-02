package api

func (l Api) Router() {
	l.gin.POST("/v1/fans", l.SpeedControlLogic)
}
