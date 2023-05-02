package dto

type AsciiResponse struct {
	Name string
	Extension string
	Status string
	Message string
	Frames []string
}

func NewImageAscii(name string, extension string, status string, message string, frames []string) AsciiResponse {
	return AsciiResponse{Name: name, Extension: extension, Status: status, Message: message, Frames: frames}
}