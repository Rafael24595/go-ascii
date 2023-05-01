package dto

type AsciiResponse struct {
	Name string
	Extension string
	Status string
	Frames []string
}

func NewImageAscii(name string, extension string, status string, frames []string) AsciiResponse {
	return AsciiResponse{Name: name, Extension: extension, Status: status, Frames: frames}
}