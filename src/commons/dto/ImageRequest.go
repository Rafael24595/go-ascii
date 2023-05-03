package dto

type ImageRequest struct {
	Code string `json:"code"`
	Height string `json:"height"`
	Width string `json:"width"`
	SwWidthFix string `json:"sw_width_fix"`
	GrayScale string `json:"gray_scale"`
	Image string `json:"image" binding:"required"`
}