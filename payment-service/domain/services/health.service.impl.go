package services

import (
	"context"
	"time"

	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/configs"
)

type HealthServiceImpl struct{}

func NewHealthServiceImpl() HealthService {
	return &HealthServiceImpl{}
}

func (HealthServiceImpl) HealthCheck() *HealthCheckResponse {
	response := HealthCheckResponse{
		Status: "up",
	}
	return &response
}

func (HealthServiceImpl) DeepHealthCheck() *HealthCheckResponse {
	response := HealthCheckResponse{}
	response.Status = "up"
	response.Services = append(response.Services, dbHealthCheck())

	for _, services := range response.Services {
		if services.Status == "down" {
			response.Status = "down"
			break
		}
	}
	return &response
}

func dbHealthCheck() ServiceResponse {
	response := ServiceResponse{}
	response.Name = "DynamoDB"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := configs.DB.ListTablesWithContext(ctx, nil)

	if err != nil {
		response.Status = "down"
		response.Error = err.Error()
	} else {
		response.Status = "up"
		response.Error = ""
	}

	return response
}
