package dto

type AsciiRequest struct {
	Code string `json:"code"`
	Image string `json:"image" binding:"required"`
}