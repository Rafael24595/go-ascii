package dto

type AsciiResponse struct {
	Name      string `bson:"name" json:"name"`
	Extension string `bson:"extension" json:"extension"`
	Status    string `bson:"status" json:"status"`
	Frames    []string `bson:"frames" json:"frames"`
}

func NewAsciiResponse(name string, extension string, status string, frames []string) AsciiResponse {
	return AsciiResponse{Name: name, Extension: extension, Status: status, Frames: frames}
}