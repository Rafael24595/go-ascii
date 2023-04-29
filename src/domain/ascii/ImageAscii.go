package ascii

import (
	"go-ascii/src/commons/utils"
	"go-ascii/src/domain/ascii/collection"
	"go-ascii/src/domain/ascii/scale"
	"image"
	"image/color"
	"image/draw"
)

type ImageAscii struct {
	Images collection.ImagesCollection
	ImgScale scale.ImageScale
	GrayScale string
}

func NewImageAscii(imgs []image.Image, scaleHeight int, scaleWidth int, grayScale string) (ascii ImageAscii) {
	collection := collection.NewImagesCollection(imgs)
	scale := scale.NewImageScale(collection, scaleHeight, scaleWidth)
	ascii = ImageAscii{Images:collection, ImgScale: scale, GrayScale:grayScale}
	return
}

func Generate(this ImageAscii) (ascii string) {
	ascii = buildImage(this)
	ascii = utils.CleanScapeChars(ascii)
	return
}

func buildImage(this ImageAscii) (ascii string) {
	for i, _ := range this.Images.Images {
		ascii += buildFrame(this, i)
		ascii += "\n\n"
	}
	return
}

func buildFrame(this ImageAscii, position int) (frame string) {
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

func resizeImage(this ImageAscii, position int) (resized *image.RGBA) {
	height := int(collection.GetImageHeight(this.Images))
	width := int(collection.GetImageWidth(this.Images))
	resized = image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(resized, resized.Bounds(), this.Images.Images[position], image.Point{0, 0}, draw.Src)
	return
}

func grayScale(this ImageAscii, position int) (grayscale *image.Gray) {
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