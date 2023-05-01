package collection

import "image"

type ImagesCollection struct {
	images []image.Image
}

func NewImagesCollection(imgs []image.Image) ImagesCollection {
	return ImagesCollection{images: imgs}
}

func (this ImagesCollection) GetImageHeight() float64 {
	return float64(this.images[0].Bounds().Dy())
}

func (this ImagesCollection) GetImageWidth() float64 {
	return float64(this.images[0].Bounds().Dx())
}

func (this ImagesCollection) GetImages() []image.Image {
	return this.images
}

func (this ImagesCollection) GetImage(index int) image.Image {
	return this.images[index]
}