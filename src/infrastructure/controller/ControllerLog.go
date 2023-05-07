package controller

import (
	"github.com/gin-gonic/gin"
	"go-ascii/src/service"
)

type ControllerLog struct {
	Service service.ServiceLog
	RouterGroup gin.RouterGroup
}

func NewControllerLog(router *gin.Engine, service service.ServiceLog) (controller ControllerLog) {
	controller = ControllerLog{Service: service, RouterGroup: *router.Group("/api")}
	controller.RouterGroup.GET("/log", controller.findAll)
	//controller.RouterGroup.GET("/log/:category", controller.find)
	return
}

func (this ControllerLog) findAll(c *gin.Context, ) {
	body := gin.H{
		"message": this.Service.FindAll(),
	}
	c.JSON(200, body)
}

func (this ControllerLog) find(c *gin.Context) {
	code := c.Param("code")
	image:= this.Service.Find(code)
	c.JSON(200, &image)
}