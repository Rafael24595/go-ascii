package controller

import (
	"go-ascii/src/infrastructure/dto"
	"go-ascii/src/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service service.Service
	RouterGroup gin.RouterGroup
}

func NewController(router *gin.Engine, service service.Service) (controller Controller) {
	controller = Controller{Service: service, RouterGroup: *router.Group("/api")}
	controller.RouterGroup.GET("/ascii", controller.findAllAscii)
	controller.RouterGroup.GET("/ascii/:code", controller.findAscii)
	controller.RouterGroup.POST("/ascii", controller.insertAscii)
	return
}

func (this Controller) findAllAscii(c *gin.Context, ) {
	body := gin.H{
		"message": this.Service.FindAllAscii(),
	}
	c.JSON(200, body)
}

func (this Controller) findAscii(c *gin.Context) {
	code := c.Param("code")
	image:= this.Service.FindAscii(code)
	c.JSON(200, &image)
}

func (this Controller) insertAscii(c *gin.Context) {
	asciiRequest := dto.AsciiRequest{}
	err := c.BindJSON(&asciiRequest)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, this.Service.InsertAscii(asciiRequest))
}