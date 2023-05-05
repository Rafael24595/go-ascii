package main

import (
	"os"
	"syscall"
	"os/signal"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"go-ascii/src/commons/temp-source"
	"go-ascii/src/infrastructure/controller"
	"go-ascii/src/infrastructure/controller/middleware"
	"go-ascii/src/infrastructure/repository"
	"go-ascii/src/service"

)

var queryRepository repository.QueryRepository
var commandRepository repository.CommandRepository

func init() {
    err := godotenv.Load(".env")
    if err != nil {
        panic(err)
    }
}

func main() {
	serve()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	tempsource.CleanSessionSources()

	queryRepository.OnExit()
	commandRepository.OnExit()
}

func serve() {
	router := gin.Default()
	router.Use(middleware.Cors())

	queryRepository = repository.NewQueryRepositoryInmemory()
	queryRepository.OnLoad()
	commandRepository = repository.NewCommandRepositoryMongo(queryRepository)
	commandRepository.OnLoad()
	serviceAscii := service.NewService(queryRepository, commandRepository)
	controller.NewControllerRest(router, serviceAscii)
	controller.NewControllerView(router, serviceAscii)

	domain := os.Getenv("GO_ASCII_DOMAIN")
	port := os.Getenv("GO_ASCII_PORT")

	go router.Run( domain + ":" + port)
}