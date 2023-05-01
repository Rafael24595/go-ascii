package dto

type ImageRequest struct {
	Code string `json:"code"`
	Image string `json:"image" binding:"required"`
}