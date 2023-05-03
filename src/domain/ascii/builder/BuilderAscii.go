package builder

import (
	"go-ascii/src/commons/constants/gray-scale"
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/commons/utils"
	"go-ascii/src/commons/utils/encoder-decoder"
	"go-ascii/src/domain/ascii"
	"go-ascii/src/domain/ascii/builder/collection"
	"go-ascii/src/domain/ascii/builder/scale"
	"image"
	"image/color"
	"image/draw"
)

const DEFAULT_WIDTH_CORRECTION = 0.4935

type BuilderAscii struct {
	event ascii.QueueEvent
	images collection.ImagesCollection
	imgeScale scale.ImageScale
	grayScale gray_scale.GrayScale 
}

func NewBuilderAscii(event ascii.QueueEvent) (builder BuilderAscii, err error) {
	scaleHeight := event.GetHeight()
	scaleWidth := event.GetWidth()
	grayScale := event.GetGrayScale()

	images, err := encoder_decoder.DecodeByPath(event.GetPath())
	collection := collection.NewImagesCollection(images)
	scale := scale.NewImageScale(collection, scaleHeight, scaleWidth)
	builder = BuilderAscii{event: event, images:collection, imgeScale: scale, grayScale: grayScale}
	return 
}

func (this BuilderAscii) Build() ascii.ImageAscii {
	code := this.event.GetCode()
	extension := utils.FileExtensionByPath(this.event.GetPath())
	imageAscii := ascii.NewImageAscii(code, extension, request_state.SUCCES, []string{})
	for i := range this.images.GetImages() {
		frame := this.buildFrame(i)
		imageAscii.AppendFrame(frame)
	}
	return imageAscii
}

func (this BuilderAscii) buildFrame(position int) (frame string) {
	height := int(this.images.GetImageHeight())
	width := int(this.images.GetImageWidth())
	
	scaleX := this.imgeScale.GetScaleX()
	scaleY := this.imgeScale.GetScaleY()
	correction := this.getCorrection()

	grayscale := this.desaturateImage(position)
	for y := 0; y < height; y+=int(scaleY) {
		for x := 0; x < width; x+= int(scaleX*correction) {
			c := grayscale.GrayAt(x, y).Y
			i := int(float64(c) / 255.0 * float64(len(this.grayScale)-1))
			frame += string(this.grayScale[i])
		}
		frame += "\n"
	}

	return
}

func (this BuilderAscii) resizeImage(position int) (resized *image.RGBA) {
	height := int(this.images.GetImageHeight())
	width := int(this.images.GetImageWidth())
	resized = image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(resized, resized.Bounds(), this.images.GetImage(position), image.Point{0, 0}, draw.Src)
	return
}

func (this BuilderAscii) desaturateImage(position int) (grayscale *image.Gray) {
	resized := this.resizeImage(position)
	height := int(this.images.GetImageHeight())
	width := int(this.images.GetImageWidth())

	grayscale = image.NewGray(resized.Bounds())

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := color.GrayModel.Convert(resized.At(x, y)).(color.Gray)
			grayscale.Set(x, y, c)
		}
	}

	return
}

func (this BuilderAscii) getCorrection() (correction float64) {
	correction = 1.0
	if this.event.SwWidthFix() {
		correction = DEFAULT_WIDTH_CORRECTION
	}
	return
}