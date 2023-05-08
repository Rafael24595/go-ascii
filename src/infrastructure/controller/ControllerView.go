package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-ascii/src/service"
	"go-ascii/src/commons/dto"
	"go-ascii/src/commons/log"
	"go-ascii/src/commons/configurator/configuration"
	"go-ascii/src/infrastructure/input-output/log-view"
	"go-ascii/src/infrastructure/input-output/form-view"
	"go-ascii/src/infrastructure/input-output/menu-view"
	"go-ascii/src/infrastructure/input-output/ascii-view"
)

type ControllerView struct {
	serviceAscii service.ServiceAscii
	serviceLog service.ServiceLog
	routerGroup gin.RouterGroup
}

func NewControllerView(router *gin.Engine, service service.ServiceAscii, serviceLog service.ServiceLog) (controller ControllerView) {
	configuration := configuration.GetInstance()
	
	controller = ControllerView{serviceAscii: service, serviceLog: serviceLog, routerGroup: *router.Group("api/view")}
	controller.routerGroup.GET("/ascii", controller.findAllAscii)
	controller.routerGroup.GET("/ascii/:code", controller.findAscii)
	controller.routerGroup.GET("/form", controller.insertAscii)
	if configuration.IsDebugSession() {
		controller.routerGroup.GET("/log", controller.filterLog)
	}
	return
}

func (this ControllerView) findAllAscii(c *gin.Context) {
	log.Log("INFO", "Get petition to end point '" + c.FullPath() + "'.")
	images := this.serviceAscii.FindAll()
	builder := menu_view.NewMenuViewBuilder(images)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(builder.Build()))
}

func (this ControllerView) findAscii(c *gin.Context) {
	log.Log("INFO", "Get petition to end point '" + c.FullPath() + "'.")
	code := c.Param("code")
	image := this.serviceAscii.Find(code)
	builder := ascii_view.NewAsciiViewBuilder(image, this.findAsciiArgs(c))
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(builder.Build()))
}

func (this ControllerView) insertAscii(c *gin.Context) {
	log.Log("INFO", "Get petition to end point '" + c.FullPath() + "'.")
	builder := form_view.NewAsciiAsciiFormViewBuilder()
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(builder.Build()))
}

func (this ControllerView) filterLog(c *gin.Context) {
	log.Log("INFO", "Get petition to end point '" + c.FullPath() + "'.")
	dto := dto.NewLogParamsRequest(c.Query("category"), c.Query("from"), c.Query("to"))
	logs := this.serviceLog.FindAll(dto)
	builder := log_view.NewLogViewBuilder(logs)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(builder.Build()))
}

func (this ControllerView) findAsciiArgs(c *gin.Context) (args map[string]string) {
	args = map[string]string{}
	args["delay"] = c.Query("delay")
	return 
}