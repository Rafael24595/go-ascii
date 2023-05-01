package scale

import (
	"go-ascii/src/domain/ascii/builder/collection"
)

type ImageScale struct {
	images collection.ImagesCollection
	scaleHeight float64
	scaleWidth  float64
}

func NewImageScale(images collection.ImagesCollection, scaleHeight int, scaleWidth int) ImageScale {
	return ImageScale{images: images, scaleHeight: float64(scaleHeight), scaleWidth: float64(scaleWidth)}
}

func (this ImageScale) GetScaleX() (scaleX float64) {
	scaleX = 0
	if this.scaleHeight != 0 {
		height := this.images.GetImageHeight()
		scaleX = height / this.scaleHeight
	} else if this.scaleWidth != 0 {
		height := this.images.GetImageHeight()
		width := this.images.GetImageWidth()
		scaleY := this.GetScaleY()
		scaleX = (scaleY / width) *  height
	}
	return
}

func (this ImageScale) GetScaleY() (scaleY float64) {
	scaleY = 0
	if this.scaleWidth != 0 {
		width := this.images.GetImageWidth()
		scaleY = width / this.scaleWidth
	} else if this.scaleHeight != 0 {
		width := this.images.GetImageWidth()
		height := this.images.GetImageHeight()
		scaleX := this.GetScaleX()
		scaleY = (scaleX / height) *  width
	}
	return
}