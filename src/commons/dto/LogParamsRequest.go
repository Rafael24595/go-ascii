package dto

type LogParamsRequest struct {
	Category string
	Family string
	From string
	To string
}

func NewLogParamsRequest(category string, family string, from string, to string) LogParamsRequest {
	return LogParamsRequest{Category: category, Family: family, From: from, To: to}
}