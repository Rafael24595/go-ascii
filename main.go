package main

import (
	"fmt"
	"go-ascii/src/commons/constants"
	"go-ascii/src/commons/temp-source"
	"go-ascii/src/commons/utils"
	"go-ascii/src/commons/utils/image"
	"go-ascii/src/domain/ascii/builder"
	"go-ascii/src/infrastructure/controller"
	"go-ascii/src/infrastructure/repository"
	"go-ascii/src/service"
	"net/http"
	"os"
	/*"os/signal"
	"syscall"*/

	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	/*** Base64 request ***/
	f, err := os.Open("golang.gif")
		if err != nil {
			panic(err)
		}
		defer f.Close()

	    encoded := image.Encoder(f)
	/***                           ***/

	path := tempsource.Base64ToSource(encoded)
	temp, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	img := image.Decode(temp)
	scaleHeight := 115
	scaleWidth := 0
	grayScale := constants.GrayScaleLevels["default"]

	builderAscii := builder.NewBuilderAscii(img, scaleHeight, scaleWidth, grayScale)
	result := builderAscii.Build()

	temp, err = os.Open(path)
	if err != nil {
		panic(err)
	}

	result.Name = filepath.Base(temp.Name())
	result.Type = utils.FileExtension(temp)

	temp.Close()

	for _, frame := range result.Frames {
		fmt.Printf(frame)
		fmt.Printf("\n\n")
	}

	fmt.Println(result.Name)
	fmt.Println(result.Type)

	tempsource.CleanSessionSources()
	/*serve()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit*/
}

func serve() {
	router := gin.Default()

	repositoryAscii := repository.NewRepositoryInmemory()
	serviceAscii := service.Service{Repository: repositoryAscii}
	controller.NewController(router, serviceAscii)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	go server.ListenAndServe()
}
