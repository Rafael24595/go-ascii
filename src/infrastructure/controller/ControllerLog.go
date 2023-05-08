package controller

import (
	"github.com/gin-gonic/gin"
	"go-ascii/src/service"
	"go-ascii/src/commons/dto"
	"go-ascii/src/commons/configurator/configuration"
)

type ControllerLog struct {
	service service.ServiceLog
	routerGroup gin.RouterGroup
}

func NewControllerLog(router *gin.Engine, service service.ServiceLog) (controller ControllerLog) {
	configuration := configuration.GetInstance()

	controller = ControllerLog{service: service, routerGroup: *router.Group("/api")}
	if configuration.IsDebugSession() {
		controller.routerGroup.GET("/log", controller.filterLog)
	}
	return
}

func (this ControllerLog) filterLog(c *gin.Context) {
	dto := dto.NewLogParamsRequest(c.Query("category"), c.Query("family"), c.Query("from"), c.Query("to"))
	body := gin.H{
		"message": this.service.FilterLog(dto),
	}
	c.JSON(200, body)
}