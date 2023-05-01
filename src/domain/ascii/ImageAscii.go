package ascii

type ImageAscii struct {
	Name string
	Type string
	Status string
	Frames []string
}

func NewImageAscii(name string, typ string, status string, frames []string) ImageAscii {
	return ImageAscii{Name: name, Type: typ, Status: status, Frames: frames}
}