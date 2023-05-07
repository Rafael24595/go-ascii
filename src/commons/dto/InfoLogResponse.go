package dto

import "time"

type InfoLogResponse struct {
	Id int `json:"id"`
	SessionId string `json:"session_id"`
	Category string `json:"category"`
	Message string `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}