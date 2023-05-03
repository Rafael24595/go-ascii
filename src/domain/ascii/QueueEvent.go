package ascii

import (
	"path/filepath"
	"go-ascii/src/commons/constants/gray-scale"
	"go-ascii/src/commons/dto"
	"go-ascii/src/commons/utils"
)

const DEFAULT_HEIGHT = 50

type QueueEvent struct {
	code  string
	height float64 
    width float64 
	swWidthFix bool
    grayScale gray_scale.GrayScale  
	image string
	path  string
	message string
}

func NewQueueEvent(dto dto.ImageRequest, path string) QueueEvent {
	var grayScale gray_scale.GrayScale
	if gray_scale.IsValidGrayScale(dto.GrayScale) {
		grayScale = gray_scale.GetGrayScale(dto.GrayScale)
	} else {
		grayScale = gray_scale.GetGrayScale(string(gray_scale.DEFAULT))
	}
	height, _ := utils.ParseFloat64(dto.Height)
	width, _ := utils.ParseFloat64(dto.Width)
	if height < 0 {
		height = 0
	}
	if width < 0 {
		width = 0
	}
	if height == 0 && height == width {
		height = DEFAULT_HEIGHT
	}
	swWidthFix, _ := utils.ParseBoolean(dto.SwWidthFix)
	
	return QueueEvent{code: filepath.Base(path), height: height, width: width, swWidthFix: swWidthFix, grayScale: grayScale, image: dto.Image, path: path, message: ""}
}

func (this QueueEvent) GetCode() string {
	return this.code
}

func (this QueueEvent) GetHeight() float64 {
	return this.height
}

func (this QueueEvent) GetWidth() float64 {
	return this.width
}

func (this QueueEvent) SwWidthFix() bool {
	return this.swWidthFix
}

func (this QueueEvent) GetGrayScale() gray_scale.GrayScale {
	return this.grayScale
}

func (this QueueEvent) GetImage() string {
	return this.image
}

func (this QueueEvent) GetPath() string {
	return this.path
}

func (this QueueEvent) GetMessage() string {
	return this.message
}

func (this *QueueEvent) SetMessage(message string) {
	this.message = message
}