package controller

import (
	"github.com/gin-gonic/gin"
	"go-ascii/src/service"
	"go-ascii/src/commons/dto"
)

type ControllerRest struct {
	Service service.Service
	RouterGroup gin.RouterGroup
}

func NewControllerRest(router *gin.Engine, service service.Service) (controller ControllerRest) {
	controller = ControllerRest{Service: service, RouterGroup: *router.Group("/api")}
	controller.RouterGroup.GET("/ascii", controller.findAll)
	controller.RouterGroup.GET("/ascii/:code", controller.find)
	controller.RouterGroup.POST("/ascii", controller.insert)
	controller.RouterGroup.PUT("/ascii/:code", controller.modify)
	controller.RouterGroup.DELETE("/ascii/:code", controller.delete)
	return
}

func (this ControllerRest) findAll(c *gin.Context, ) {
	body := gin.H{
		"message": this.Service.FindAll(),
	}
	c.JSON(200, body)
}

func (this ControllerRest) find(c *gin.Context) {
	code := c.Param("code")
	image:= this.Service.Find(code)
	c.JSON(200, &image)
}

func (this ControllerRest) insert(c *gin.Context) {
	asciiRequest := dto.ImageRequest{}
	err := c.BindJSON(&asciiRequest)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, this.Service.Insert(asciiRequest))
}

func (this ControllerRest) modify(c *gin.Context) {
	code := c.Param("code")
	image:= this.Service.Modify(code)
	c.JSON(200, &image)
}

func (this ControllerRest) delete(c *gin.Context) {
	code := c.Param("code")
	image:= this.Service.Delete(code)
	c.JSON(200, &image)
}