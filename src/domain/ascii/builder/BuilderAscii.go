package builder

import (
	"go-ascii/src/commons/utils"
	"go-ascii/src/domain/ascii"
	"go-ascii/src/domain/ascii/builder/collection"
	"go-ascii/src/domain/ascii/builder/scale"
	"image"
	"image/color"
	"image/draw"
)

type BuilderAscii struct {
	Images collection.ImagesCollection
	ImgScale scale.ImageScale
	GrayScale string
}

func NewBuilderAscii(imgs []image.Image, scaleHeight int, scaleWidth int, grayScale string) (ascii BuilderAscii) {
	collection := collection.NewImagesCollection(imgs)
	scale := scale.NewImageScale(collection, scaleHeight, scaleWidth)
	ascii = BuilderAscii{Images:collection, ImgScale: scale, GrayScale:grayScale}
	return
}

func Build(this BuilderAscii) (imageAscii ascii.ImageAscii) {
	imageAscii = ascii.NewImageAscii("", "", []string{})
	for i, _ := range this.Images.Images {
		frame := buildFrame(this, i)
		frame = utils.CleanScapeChars(frame)
		imageAscii.Frames = append(imageAscii.Frames, frame)
	}
	return
}

func buildFrame(this BuilderAscii, position int) (frame string) {
	height := int(collection.GetImageHeight(this.Images))
	width := int(collection.GetImageWidth(this.Images))
	
	scaleX := scale.GetScaleX(this.ImgScale)
	scaleY := scale.GetScaleY(this.ImgScale)

	grayscale := grayScale(this, position)
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

func resizeImage(this BuilderAscii, position int) (resized *image.RGBA) {
	height := int(collection.GetImageHeight(this.Images))
	width := int(collection.GetImageWidth(this.Images))
	resized = image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(resized, resized.Bounds(), this.Images.Images[position], image.Point{0, 0}, draw.Src)
	return
}

func grayScale(this BuilderAscii, position int) (grayscale *image.Gray) {
	resized := resizeImage(this, position)
	height := int(collection.GetImageHeight(this.Images))
	width := int(collection.GetImageWidth(this.Images))

	grayscale = image.NewGray(resized.Bounds())

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := color.GrayModel.Convert(resized.At(x, y)).(color.Gray)
			grayscale.Set(x, y, c)
		}
	}
	return
}