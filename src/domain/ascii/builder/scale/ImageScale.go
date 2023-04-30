package scale

import (
	"go-ascii/src/domain/ascii/builder/collection"
)

type ImageScale struct {
	Images collection.ImagesCollection
	ScaleHeight float64
	ScaleWidth  float64
}

func NewImageScale(imgs collection.ImagesCollection, scaleHeight int, scaleWidth int) ImageScale {
	return ImageScale{Images: imgs, ScaleHeight: float64(scaleHeight), ScaleWidth: float64(scaleWidth)}
}

func (this ImageScale) GetScaleX() (scaleX float64) {
	scaleX = 0
	if this.ScaleHeight != 0 {
		height := this.Images.GetImageHeight()
		scaleX = height / this.ScaleHeight
	} else if this.ScaleWidth != 0 {
		height := this.Images.GetImageHeight()
		width := this.Images.GetImageWidth()
		scaleY := this.GetScaleY()
		scaleX = (scaleY / width) *  height
	}
	return
}

func (this ImageScale) GetScaleY() (scaleY float64) {
	scaleY = 0
	if this.ScaleWidth != 0 {
		width := this.Images.GetImageWidth()
		scaleY = width / this.ScaleWidth
	} else if this.ScaleHeight != 0 {
		width := this.Images.GetImageWidth()
		height := this.Images.GetImageHeight()
		scaleX := this.GetScaleX()
		scaleY = (scaleX / height) *  width
	}
	return
}