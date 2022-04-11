package main

import (
	// "context"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-ipp/checkout-service/configs"
	"github.com/swiggy-ipp/checkout-service/routes"
)

func main() {
	// Set up routes for Bookkeeping Microservice
	cartRouter := gin.Default()
	routes.GenerateCheckoutRoutes(cartRouter)

	// Set up Kafka listener
	// ctx := context.Background()
	// go services.Consume(configs.KafkaDiagnosisTopic(), services.DeserializeAndSaveDiseaseDiagnosis, ctx)
	
	// Run Microservice
	cartRouter.Run(":" + configs.EnvServicePort())
}
