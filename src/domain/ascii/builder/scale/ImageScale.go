package scale

import (
	"go-ascii/src/domain/ascii/builder/collection"
)

type ImageScale struct {
	Images collection.ImagesCollection
	ScaleHeight float64
	ScaleWidth  float64
}

func NewImageScale(imgs collection.ImagesCollection, scaleHeight int, scaleWidth int) (scale ImageScale) {
	scale = ImageScale{Images: imgs, ScaleHeight: float64(scaleHeight), ScaleWidth: float64(scaleWidth)}
	return
}

func GetScaleX(this ImageScale) (scaleX float64) {
	scaleX = 0
	if this.ScaleHeight != 0 {
		height := collection.GetImageHeight(this.Images)
		scaleX = height / this.ScaleHeight
	} else if this.ScaleWidth != 0 {
		height := collection.GetImageHeight(this.Images)
		width := collection.GetImageWidth(this.Images)
		scaleY := GetScaleY(this)
		scaleX = (scaleY / width) *  height
	}
	return
}

func GetScaleY(this ImageScale) (scaleY float64) {
	scaleY = 0
	if this.ScaleWidth != 0 {
		width := collection.GetImageWidth(this.Images)
		scaleY = width / this.ScaleWidth
	} else if this.ScaleHeight != 0 {
		width := collection.GetImageWidth(this.Images)
		height := collection.GetImageHeight(this.Images)
		scaleX := GetScaleX(this)
		scaleY = (scaleX / height) *  width
	}
	return
}