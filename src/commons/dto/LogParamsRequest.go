package dto

type LogParamsRequest struct {
	Category string
	From string
	To string
}

func NewLogParamsRequest(category string, from string, to string) LogParamsRequest {
	return LogParamsRequest{Category: category, From: from, To: to}
}