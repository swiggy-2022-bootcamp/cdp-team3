package services

type HealthService interface {
	HealthCheck() *HealthCheckResponse
	DeepHealthCheck() *HealthCheckResponse
}

type ServiceResponse struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Error  string `json:"error"`
}

type HealthCheckResponse struct {
	Status   string            `json:"status"`
	Services []ServiceResponse `json:"services"`
}
