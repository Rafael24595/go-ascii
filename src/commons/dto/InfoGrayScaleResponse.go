package dto

type InfoGrayScaleResponse struct {
	Id      string   `bson:"id" json:"id"`
	Description string   `bson:"description" json:"description"`
}

func NewInfoGrayScaleResponse(id string, description string) InfoGrayScaleResponse {
	return InfoGrayScaleResponse{Id: id, Description: description}
}