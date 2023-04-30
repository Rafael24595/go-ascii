package dto

type AsciiRequest struct {
	Code string `json:"code" binding:"required"`
	Image string `json:"image" binding:"required"`
}