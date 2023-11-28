package dto

type InfoAsciiResponse struct {
	Name      string   `bson:"name" json:"name"`
	Extension string   `bson:"extension" json:"extension"`
	Height    int      `bson:"height" json:"height"`
	Width     int      `bson:"width" json:"width"`
	Status    string   `bson:"status" json:"status"`
	Message   string   `bson:"message" json:"message"`
	Frames    []string `bson:"frames" json:"frames"`
}

func NewInfoAsciiResponse(name string, extension string, height int, width int, status string, message string, frames []string) InfoAsciiResponse {
	return InfoAsciiResponse{Name: name, Extension: extension, Height: height, Width: width, Status: status, Message: message, Frames: frames}
}