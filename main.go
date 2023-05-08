package main

import (
	"os"
	"syscall"
	"os/signal"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-ascii/src/commons/configurator"
	"go-ascii/src/commons/configurator/configuration"
	"go-ascii/src/commons/dependency-container"
	"go-ascii/src/commons/log"
	"go-ascii/src/commons/temp-source"
	"go-ascii/src/infrastructure/controller"
	"go-ascii/src/infrastructure/controller/middleware"
	"go-ascii/src/service"
	"go-ascii/src/commons/constants/log-categories"
)

func init() {
    err := godotenv.Load(".env")
    if err != nil {
        panic(err)
    }
}

func main() {
	onLoad()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	onExit()
}

func onLoad() {
	configurator.LoadConfiguration()
	serve()
}

func onExit() {
	dependencyContainer := dependency_container.GetInstance()

	tempsource.CleanSessionSources()

	dependencyContainer.OnExit()
}

func serve() {
	router := gin.Default()
	router.Use(middleware.Cors())
	
	dependencyContainer := dependency_container.GetInstance()
	logRepository := dependencyContainer.GetLogRepository()
	queryRepository := dependencyContainer.GetQueryRepository()
	commandRepository:= dependencyContainer.GetCommandRepository()

	serviceLog := service.NewServiceLog(logRepository)
	controller.NewControllerLog(router, serviceLog)

	serviceAscii := service.NewServiceAscii(queryRepository, commandRepository)
	controller.NewControllerRest(router, serviceAscii)
	controller.NewControllerView(router, serviceAscii, serviceLog)

	configuration := configuration.GetInstance()

	go router.Run(configuration.GetAddr())
	
	log.LogFam(log_categories.INFO, controller.Family, "Listening and serving HTTP on " + configuration.GetAddr() +".")
}