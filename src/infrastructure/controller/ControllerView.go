package controller

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"go-ascii/src/service"
	"go-ascii/src/view/ascii-view"
	"go-ascii/src/view/menu-view"
)

type ControllerView struct {
	Service     service.Service
	RouterGroup gin.RouterGroup
}

func NewControllerView(router *gin.Engine, service service.Service) (controller ControllerView) {
	controller = ControllerView{Service: service, RouterGroup: *router.Group("api/view")}
	controller.RouterGroup.GET("/ascii", controller.findAllAscii)
	controller.RouterGroup.GET("/ascii/:code", controller.findAscii)
	return
}

func (this ControllerView) findAllAscii(c *gin.Context) {
	images := this.Service.FindAllAscii()
	print(len(images))
	builder := menu_view.NewMenuViewBuilder(images, c.Request.Host)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(builder.Build()))
}

func (this ControllerView) findAscii(c *gin.Context) {
	code := c.Param("code")
	image := this.Service.FindAscii(code)
	builder := ascii_view.NewAsciiViewBuilder(image)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(builder.Build()))
}