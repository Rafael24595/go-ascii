package main

import (
	"os"
	"syscall"
	"os/signal"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"go-ascii/src/commons/configurator"
	"go-ascii/src/commons/dependency-container"
	"go-ascii/src/commons/temp-source"
	"go-ascii/src/infrastructure/controller"
	"go-ascii/src/infrastructure/controller/middleware"
	"go-ascii/src/service"

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
	queryRepository := dependencyContainer.GetQueryRepository()
	commandRepository:= dependencyContainer.GetCommandRepository()

	serviceAscii := service.NewService(queryRepository, commandRepository)
	controller.NewControllerRest(router, serviceAscii)
	controller.NewControllerView(router, serviceAscii)

	configuration := configurator.GetInstance()

	go router.Run(configuration.GetAddr())
}