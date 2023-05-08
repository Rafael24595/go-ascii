package dto

type AsciiResponse struct {
	Name      string `bson:"name" json:"name"`
	Extension string `bson:"extension" json:"extension"`
	Status    string `bson:"status" json:"status"`
	Timestamp int64 `bson:"timestamp" json:"timestamp"`
	Frames    []string `bson:"frames" json:"frames"`
}

func NewAsciiResponse(name string, extension string, status string, timestamp int64, frames []string) AsciiResponse {
	return AsciiResponse{Name: name, Extension: extension, Status: status, Timestamp: timestamp, Frames: frames}
}