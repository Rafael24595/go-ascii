package collection

import "image"

type ImagesCollection struct {
	Images []image.Image
}

func NewImagesCollection(imgs []image.Image) ImagesCollection {
	return ImagesCollection{Images: imgs}
}

func (this ImagesCollection) GetImageHeight() float64 {
	return float64(this.Images[0].Bounds().Dy())
}

func (this ImagesCollection) GetImageWidth() float64 {
	return float64(this.Images[0].Bounds().Dx())
}