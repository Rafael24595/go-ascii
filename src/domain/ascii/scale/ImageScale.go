package dimensions

import "image"

type ImageScale struct {
	Image image.Image
	ScaleHeight float64
	ScaleWidth  float64
}

func NewImageScale(img image.Image, scaleHeight int, scaleWidth int) (scale ImageScale) {
	scale = ImageScale{Image: img, ScaleHeight: float64(scaleHeight), ScaleWidth: float64(scaleWidth)}
	return
}

func GetScaleX(this ImageScale) (scaleX float64) {
	scaleX = 0
	if this.ScaleHeight != 0 {
		scaleX = getImageHeight(this) / this.ScaleHeight
	} else if this.ScaleWidth != 0 {
		scaleY := GetScaleY(this)
		scaleX = (scaleY / getImageWidth(this)) *  getImageHeight(this)
	}
	return
}

func GetScaleY(this ImageScale) (scaleY float64) {
	scaleY = 0
	if this.ScaleWidth != 0 {
		scaleY = getImageWidth(this) / this.ScaleWidth
	} else  if this.ScaleHeight != 0 {
		scaleX := GetScaleX(this)
		scaleY = (scaleX / getImageHeight(this)) *  getImageWidth(this)
	}
	return
}

func getImageHeight(this ImageScale) (width float64) {
	width = float64(this.Image.Bounds().Dx())
	return
}

func getImageWidth(this ImageScale) (height float64) {
	height = float64(this.Image.Bounds().Dy())
	return
}