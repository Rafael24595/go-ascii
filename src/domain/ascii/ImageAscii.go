package ascii

type ImageAscii struct {
	Name string
	Type string
	Frames []string
}

func NewImageAscii(name string, typ string, frames []string) ImageAscii {
	return ImageAscii{Name: name, Type: typ, Frames: frames}
}