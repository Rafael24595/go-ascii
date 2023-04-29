package collection

import "image"

type ImagesCollection struct {
	Images []image.Image
}

func NewImagesCollection(imgs []image.Image) (collection ImagesCollection) {
	collection = ImagesCollection{Images: imgs}
	return
}

func GetImageHeight(this ImagesCollection) (width float64) {
	width = float64(this.Images[0].Bounds().Dy())
	return
}

func GetImageWidth(this ImagesCollection) (height float64) {
	height = float64(this.Images[0].Bounds().Dx())
	return
}