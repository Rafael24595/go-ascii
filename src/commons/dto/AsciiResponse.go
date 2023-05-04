package dto

type AsciiResponse struct {
	Name string
	Extension string
	Height int
	Width int
	Status string
	Message string
	Frames []string
}

func NewImageAscii(name string, extension string, height int, width int, status string, message string, frames []string) AsciiResponse {
	return AsciiResponse{Name: name, Extension: extension, Height: height, Width: width, Status: status, Message: message, Frames: frames}
}