package dto

type ResponseDTO struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

type HealthCheckResponse struct {
	Server   string `json:"server"`
	Database string `json:"database"`
}