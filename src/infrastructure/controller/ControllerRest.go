package controller

import (
	"github.com/gin-gonic/gin"
	"go-ascii/src/service"
	"go-ascii/src/commons/dto"
)

type ControllerRest struct {
	service service.ServiceAscii
	routerGroup gin.RouterGroup
}

func NewControllerRest(router *gin.Engine, service service.ServiceAscii) (controller ControllerRest) {
	controller = ControllerRest{service: service, routerGroup: *router.Group("/api")}
	controller.routerGroup.GET("/ascii", controller.findAll)
	controller.routerGroup.GET("/ascii/:code", controller.find)
	controller.routerGroup.POST("/ascii", controller.insert)
	controller.routerGroup.PUT("/ascii/:code", controller.modify)
	controller.routerGroup.DELETE("/ascii/:code", controller.delete)
	return
}

func (this ControllerRest) findAll(c *gin.Context, ) {
	body := gin.H{
		"message": this.service.FindAll(),
	}
	c.JSON(200, body)
}

func (this ControllerRest) find(c *gin.Context) {
	code := c.Param("code")
	image:= this.service.Find(code)
	c.JSON(200, &image)
}

func (this ControllerRest) insert(c *gin.Context) {
	asciiRequest := dto.ImageRequest{}
	err := c.BindJSON(&asciiRequest)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, this.service.Insert(asciiRequest))
}

func (this ControllerRest) modify(c *gin.Context) {
	code := c.Param("code")
	image:= this.service.Modify(code)
	c.JSON(200, &image)
}

func (this ControllerRest) delete(c *gin.Context) {
	code := c.Param("code")
	image:= this.service.Delete(code)
	c.JSON(200, &image)
}