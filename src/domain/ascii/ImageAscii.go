package ascii

type ImageAscii struct {
	Name string
	Type string
	Frames []string
}

func NewImageAscii(name string, typ string, frames []string) (imageAscii ImageAscii) {
	imageAscii = ImageAscii{Name: name, Type: typ, Frames: frames}
	return
}