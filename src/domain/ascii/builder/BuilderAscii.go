package builder

import (
	"image"
	"image/color"
	"image/draw"
	"go-ascii/src/commons/constants/gray-scale"
	"go-ascii/src/commons/constants/request-state"
	"go-ascii/src/domain/ascii"
	"go-ascii/src/domain/ascii/builder/collection"
	"go-ascii/src/domain/ascii/builder/scale"
)

type BuilderAscii struct {
	images collection.ImagesCollection
	imgeScale scale.ImageScale
	grayScale gray_scale.GrayScale
}

func NewBuilderAscii(imgs []image.Image, scaleHeight int, scaleWidth int, grayScale gray_scale.GrayScale) BuilderAscii {
	collection := collection.NewImagesCollection(imgs)
	scale := scale.NewImageScale(collection, scaleHeight, scaleWidth)
	return BuilderAscii{images:collection, imgeScale: scale, grayScale: grayScale}
}

func (this BuilderAscii) Build() ascii.ImageAscii {
	imageAscii := ascii.NewImageAscii("", "", request_state.SUCCES, []string{})
	for i := range this.images.GetImages() {
		frame := this.buildFrame(i)
		//frame = utils.CleanScapeChars(frame)
		imageAscii.AppendFrame(frame)
	}
	return imageAscii
}

func (this BuilderAscii) buildFrame(position int) (frame string) {
	height := int(this.images.GetImageHeight())
	width := int(this.images.GetImageWidth())
	
	scaleX := this.imgeScale.GetScaleX()
	scaleY := this.imgeScale.GetScaleY()

	grayscale := this.desaturateImage(position)
	for y := 0; y < height; y+=int(scaleY*2.5) {
		for x := 0; x < width; x+= int(scaleX) {
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