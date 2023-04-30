package main

import (
	"go-ascii/src/commons/temp-source"
	"go-ascii/src/infrastructure/controller"
	"go-ascii/src/infrastructure/repository"
	"go-ascii/src/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"github.com/gin-gonic/gin"
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

	repositoryAscii := repository.NewRepositoryInmemory()
	serviceAscii := service.NewService(repositoryAscii)
	controller.NewController(router, serviceAscii)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	go server.ListenAndServe()
}
