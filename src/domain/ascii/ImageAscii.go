package ascii

import (
	"go-ascii/src/commons/utils"
	"go-ascii/src/domain/ascii/scale"
	"image"
	"image/color"
	"image/draw"
)

type ImageAscii struct {
	Image []image.Image
	ImgScale dimensions.ImageScale
	GrayScale string
}

func NewImageAscii(img []image.Image, scaleHeight int, scaleWidth int, grayScale string) (ascii ImageAscii) {
	scale := dimensions.NewImageScale(img[0], scaleHeight, scaleWidth)
	ascii = ImageAscii{Image:img, ImgScale: scale, GrayScale:grayScale}
	return
}

func Generate(this ImageAscii) (ascii string) {
	ascii = build(this)
	ascii = utils.CleanScapeChars(ascii)
	return
}

func build(this ImageAscii) (ascii string) {
	height := getImageHeight(this)
	width := getImageWidth(this)
	
	scaleX := dimensions.GetScaleX(this.ImgScale)
	scaleY := dimensions.GetScaleY(this.ImgScale)

	for i, _ := range this.Image {
		grayscale := grayScale(this, i)
		for y := 0; y < height; y+=int(scaleY*2.5) {
			for x := 0; x < width; x+= int(scaleX) {
				c := grayscale.GrayAt(x, y).Y
				i := int(float64(c) / 255.0 * float64(len(this.GrayScale)-1))
				ascii += string(this.GrayScale[i])
			}
			ascii += "\n"
		}
		ascii += "\n\n"
	}

	return
}

func resizeImage(this ImageAscii, i int) (resized *image.RGBA) {
	height := getImageHeight(this)
	width := getImageWidth(this)
	resized = image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(resized, resized.Bounds(), this.Image[i], image.Point{0, 0}, draw.Src)
	return
}

func grayScale(this ImageAscii, i int) (grayscale *image.Gray) {
	resized := resizeImage(this, i)
	height := getImageHeight(this)
	width := getImageWidth(this)

	grayscale = image.NewGray(resized.Bounds())

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := color.GrayModel.Convert(resized.At(x, y)).(color.Gray)
			grayscale.Set(x, y, c)
		}
	}
	return
}

func getImageHeight(this ImageAscii) (width int) {
	width = int(float64(this.Image[0].Bounds().Dx()))
	return
}

func getImageWidth(this ImageAscii) (height int) {
	height = int(float64(this.Image[0].Bounds().Dy()))
	return
}