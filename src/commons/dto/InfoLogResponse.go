package dto

type InfoLogResponse struct {
	Id int `json:"id"`
	SessionId string `json:"session_id"`
	Category string `json:"category"`
	Family string `json:"family"`
	Message string `json:"message"`
	Timestamp int64 `json:"timestamp"`
}

func NewInfoLogResponse(id int, sessionId string, category string, family string, message string, timestamp int64) InfoLogResponse {
	return InfoLogResponse{Id: id, SessionId: sessionId, Category: category, Family: family, Message: message, Timestamp: timestamp}
}