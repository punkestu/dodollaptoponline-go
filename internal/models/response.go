package models

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(status int, message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
