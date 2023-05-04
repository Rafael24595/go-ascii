package dto

type InfoAsciiResponse struct {
	Name string
	Extension string
	Height int
	Width int
	Status string
	Message string
	Frames []string
}

func NewInfoAsciiResponse(name string, extension string, height int, width int, status string, message string, frames []string) InfoAsciiResponse {
	return InfoAsciiResponse{Name: name, Extension: extension, Height: height, Width: width, Status: status, Message: message, Frames: frames}
}