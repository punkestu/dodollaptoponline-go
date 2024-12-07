package models

type SuccessResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    interface{}   `json:"data,omitempty"`
	Meta    *MetaResponse `json:"meta,omitempty"`
}

type MetaResponse struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

func NewSuccessResponse(status int, message string, data interface{}, meta *MetaResponse) *SuccessResponse {
	return &SuccessResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}
