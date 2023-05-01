package builder

import (
	"image"
	"image/color"
	"image/draw"
	"go-ascii/src/commons/constants"
	"go-ascii/src/domain/ascii"
	"go-ascii/src/domain/ascii/builder/collection"
	"go-ascii/src/domain/ascii/builder/scale"
)

type BuilderAscii struct {
	Images collection.ImagesCollection
	ImgeScale scale.ImageScale
	GrayScale string
}

func NewBuilderAscii(imgs []image.Image, scaleHeight int, scaleWidth int, grayScale string) BuilderAscii {
	collection := collection.NewImagesCollection(imgs)
	scale := scale.NewImageScale(collection, scaleHeight, scaleWidth)
	return BuilderAscii{Images:collection, ImgeScale: scale, GrayScale:grayScale}
}

func (this BuilderAscii) Build() (imageAscii ascii.ImageAscii) {
	imageAscii = ascii.NewImageAscii("", "", constants.SUCCES, []string{})
	for i := range this.Images.Images {
		frame := this.buildFrame(i)
		//frame = utils.CleanScapeChars(frame)
		imageAscii.Frames = append(imageAscii.Frames, frame)
	}
	return
}

func (this BuilderAscii) buildFrame(position int) (frame string) {
	height := int(this.Images.GetImageHeight())
	width := int(this.Images.GetImageWidth())
	
	scaleX := this.ImgeScale.GetScaleX()
	scaleY := this.ImgeScale.GetScaleY()

	grayscale := this.grayScale(position)
	for y := 0; y < height; y+=int(scaleY*2.5) {
		for x := 0; x < width; x+= int(scaleX) {
			c := grayscale.GrayAt(x, y).Y
			i := int(float64(c) / 255.0 * float64(len(this.GrayScale)-1))
			frame += string(this.GrayScale[i])
		}
		frame += "\n"
	}

	return
}

func (this BuilderAscii) resizeImage(position int) (resized *image.RGBA) {
	height := int(this.Images.GetImageHeight())
	width := int(this.Images.GetImageWidth())
	resized = image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(resized, resized.Bounds(), this.Images.Images[position], image.Point{0, 0}, draw.Src)
	return
}

func (this BuilderAscii) grayScale(position int) (grayscale *image.Gray) {
	resized := this.resizeImage(position)
	height := int(this.Images.GetImageHeight())
	width := int(this.Images.GetImageWidth())

	grayscale = image.NewGray(resized.Bounds())

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := color.GrayModel.Convert(resized.At(x, y)).(color.Gray)
			grayscale.Set(x, y, c)
		}
	}

	return
}