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
	controller.RouterGroup.GET("/ascii", controller.findAllAscii)
	controller.RouterGroup.GET("/ascii/:code", controller.findAscii)
	controller.RouterGroup.POST("/ascii", controller.insertAscii)
	return
}

func (this ControllerRest) findAllAscii(c *gin.Context, ) {
	body := gin.H{
		"message": this.Service.FindAllAscii(),
	}
	c.JSON(200, body)
}

func (this ControllerRest) findAscii(c *gin.Context) {
	code := c.Param("code")
	image:= this.Service.FindAscii(code)
	c.JSON(200, &image)
}

func (this ControllerRest) insertAscii(c *gin.Context) {
	asciiRequest := dto.ImageRequest{}
	err := c.BindJSON(&asciiRequest)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, this.Service.InsertAscii(asciiRequest))
}