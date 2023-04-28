package main

import (
	"fmt"
	"go-ascii/src/commons/constants"
	"go-ascii/src/commons/temp-source"
	"go-ascii/src/commons/utils/image"
	"go-ascii/src/domain/ascii"
	"os"
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

	imageAscii := ascii.NewImageAscii(img, scaleHeight, scaleWidth, grayScale)
	result := ascii.Generate(imageAscii)

	fmt.Printf(result)

	tempsource.CleanSessionSources()
}