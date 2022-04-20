package utils

import (
	"context"
	"time"

	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/configs"
)

type ServiceResponse struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Error  string `json:"error"`
}

type HealthCheckResponse struct {
	Status   string            `json:"status"`
	Services []ServiceResponse `json:"services"`
}

func HealthCheck() HealthCheckResponse {
	response := HealthCheckResponse{
		Status: "up",
	}
	return response
}

func DeepHealthCheck() HealthCheckResponse {
	response := HealthCheckResponse{}
	response.Status = "up"
	response.Services = append(response.Services, DBHealthCheck())

	for _, services := range response.Services {
		if services.Status == "down" {
			response.Status = "down"
			break
		}
	}
	return response
}

func DBHealthCheck() ServiceResponse {
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
