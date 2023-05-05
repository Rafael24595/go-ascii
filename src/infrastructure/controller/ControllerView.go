package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-ascii/src/service"
	"go-ascii/src/infrastructure/input-output/ascii-view"
	"go-ascii/src/infrastructure/input-output/form-view"
	"go-ascii/src/infrastructure/input-output/menu-view"
)

type ControllerView struct {
	Service     service.Service
	RouterGroup gin.RouterGroup
}

func NewControllerView(router *gin.Engine, service service.Service) (controller ControllerView) {
	controller = ControllerView{Service: service, RouterGroup: *router.Group("api/view")}
	controller.RouterGroup.GET("/ascii", controller.findAllAscii)
	controller.RouterGroup.GET("/ascii/:code", controller.findAscii)
	controller.RouterGroup.GET("/form", controller.insertAscii)
	return
}

func (this ControllerView) findAllAscii(c *gin.Context) {
	images := this.Service.FindAll()
	builder := menu_view.NewMenuViewBuilder(images)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(builder.Build()))
}

func (this ControllerView) findAscii(c *gin.Context) {
	code := c.Param("code")
	image := this.Service.Find(code)
	builder := ascii_view.NewAsciiViewBuilder(image, this.findAsciiArgs(c))
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(builder.Build()))
}

func (this ControllerView) insertAscii(c *gin.Context) {
	builder := form_view.NewAsciiAsciiFormViewBuilder()
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(builder.Build()))
}

func (this ControllerView) findAsciiArgs(c *gin.Context) (args map[string]string) {
	args = map[string]string{}
	args["delay"] = c.Query("delay")
	return 
}