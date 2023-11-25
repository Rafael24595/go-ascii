package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-ascii/src/commons/dto"
	"go-ascii/src/service"
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
	status := http.StatusOK
	body := gin.H{
		"message": this.service.FindAll(),
	}
	c.JSON(status, body)
	logRequest(c, Family, status)
}

func (this ControllerRest) find(c *gin.Context) {
	status := http.StatusOK
	code := c.Param("code")
	image, ok:= this.service.Find(code)
	if !ok {
		c.JSON(404, &image)
	}
	c.JSON(status, &image)
	logRequest(c, Family, status)
}

func (this ControllerRest) insert(c *gin.Context) {
	status := http.StatusOK
	asciiRequest := dto.ImageRequest{}
	err := c.BindJSON(&asciiRequest)
	if err != nil {
		status = 500
		c.JSON(status, err)
	}
	c.JSON(status, this.service.Insert(asciiRequest))
	logRequest(c, Family, status)
}

func (this ControllerRest) modify(c *gin.Context) {
	status := http.StatusOK
	code := c.Param("code")
	image, ok := this.service.Modify(code)
	if ok {
		c.JSON(200, &image)
	} else {
		c.JSON(404, &image)
	}
	logRequest(c, Family, status)
}

func (this ControllerRest) delete(c *gin.Context) {
	status := http.StatusOK
	code := c.Param("code")
	image, ok := this.service.Delete(code)
	if ok {
		c.JSON(200, &image)
	} else {
		c.JSON(404, &image)
	}
	c.JSON(status, &image)
	logRequest(c, Family, status)
}