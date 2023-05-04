package main

import (
	"os"
	"syscall"
	"os/signal"
	"github.com/gin-gonic/gin"
	"go-ascii/src/commons/temp-source"
	"go-ascii/src/infrastructure/controller"
	"go-ascii/src/infrastructure/controller/middleware"
	"go-ascii/src/infrastructure/repository"
	"go-ascii/src/service"
)

func main() {
	serve()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	tempsource.CleanSessionSources()
}

func serve() {
	router := gin.Default()
	router.Use(middleware.Cors())

	queryRepository := repository.NewQueryRepositoryInmemory()
	queryRepository.OnLoad()
	commandRepository := repository.NewCommandRepositoryMongo(queryRepository)
	commandRepository.OnLoad()
	serviceAscii := service.NewService(queryRepository, commandRepository)
	controller.NewControllerRest(router, serviceAscii)
	controller.NewControllerView(router, serviceAscii)

	go router.Run("0.0.0.0:8080")
}