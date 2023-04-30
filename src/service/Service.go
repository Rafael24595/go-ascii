package service

import (
	"go-ascii/src/commons/constants"
	"go-ascii/src/commons/temp-source"
	"go-ascii/src/commons/utils"
	"go-ascii/src/commons/utils/image"
	"go-ascii/src/domain/ascii"
	"go-ascii/src/domain/ascii/builder"
	"go-ascii/src/infrastructure/dto"
	"go-ascii/src/infrastructure/repository"
	"os"
	"path/filepath"
)

type Service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return Service{repository: repository}
}

func (this Service) FindAllAscii() (result string) {
	return this.repository.FindAllAscii()
}

func (this Service) FindAscii(code string) ascii.ImageAscii {
	image := this.repository.FindAscii(code)
	return image
}

func (this Service) InsertAscii(dto dto.AsciiRequest) string {
	imageAscii := create(dto)
	return this.repository.InsertAscii(imageAscii)
}

func create(dto dto.AsciiRequest) ascii.ImageAscii {
	path := tempsource.Base64ToSource(dto.Image, dto.Code)
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

	return result
}