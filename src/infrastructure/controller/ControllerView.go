package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-ascii/src/commons/configurator/configuration"
	"go-ascii/src/commons/dto"
	"go-ascii/src/infrastructure/input-output/ascii-view"
	"go-ascii/src/infrastructure/input-output/form-view"
	"go-ascii/src/infrastructure/input-output/log-view"
	"go-ascii/src/infrastructure/input-output/menu-view"
	"go-ascii/src/service"
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
	status := http.StatusOK
	images := this.serviceAscii.FindAll()
	builder := menu_view.NewMenuViewBuilder(images)
	c.Data(status, "text/html; charset=utf-8", []byte(builder.Build()))
	logRequest(c, Family, status)
}

func (this ControllerView) findAscii(c *gin.Context) {
	status := http.StatusOK
	code := c.Param("code")
	image := this.serviceAscii.Find(code)
	builder := ascii_view.NewAsciiViewBuilder(image, this.findAsciiArgs(c))
	c.Data(status, "text/html; charset=utf-8", []byte(builder.Build()))
	logRequest(c, Family, status)
}

func (this ControllerView) insertAscii(c *gin.Context) {
	status := http.StatusOK
	builder := form_view.NewAsciiAsciiFormViewBuilder()
	c.Data(status, "text/html; charset=utf-8", []byte(builder.Build()))
	logRequest(c, Family, status)
}

func (this ControllerView) filterLog(c *gin.Context) {
	status := http.StatusOK
	dto := dto.NewLogParamsRequest(c.Query("category"), c.Query("family"), c.Query("from"), c.Query("to"))
	logs := this.serviceLog.FilterLog(dto)
	builder := log_view.NewLogViewBuilder(logs)
	c.Data(status, "text/html; charset=utf-8", []byte(builder.Build(dto)))
	logRequest(c, Family, status)
}

func (this ControllerView) findAsciiArgs(c *gin.Context) (args map[string]string) {
	args = map[string]string{}
	args["delay"] = c.Query("delay")
	return 
}