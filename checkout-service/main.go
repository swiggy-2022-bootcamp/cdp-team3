package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/swiggy-ipp/checkout-service/configs"
	"github.com/swiggy-ipp/checkout-service/controllers"
	"github.com/swiggy-ipp/checkout-service/grpcs"
	"github.com/swiggy-ipp/checkout-service/routes"
)

var (
	// Error Channels
	errChanKafka chan error = make(chan error)
	errChanREST  chan error = make(chan error)
)

/// Function with logic for starting REST Routes
func generateRESTRoutes(port string) {
	checkoutController := controllers.NewCheckoutController(
		<-grpcs.CartCheckoutGRPCChannel,
		<-grpcs.ShippingCheckoutGRPCChannel,
		<-grpcs.OrderCheckoutGRPCChannel,
	) // Controller
	checkoutRouter := gin.Default()
	routes.GenerateCheckoutRoutes(checkoutRouter, checkoutController)

	// Run REST Microservice
	errChanREST <- checkoutRouter.Run(":" + port)
}

func main() {
	// Set up GRPC
	go grpcs.BecomeGRPCClient(
		configs.EnvCartServiceGRPCAddress(),
		configs.EnvShippingServiceGRPCAddress(),
		configs.EnvOrderGRPCAddress(),
	)

	// Set up routes for Checkout Microservice
	go generateRESTRoutes(configs.EnvServicePort())

	// Listen to errors.
	select {
	case err := <-grpcs.ErrChanGRPC:
		log.Fatal("GRPC encountered an error: ", err)
	case err := <-errChanKafka:
		log.Fatal("Kafka encountered an error: ", err)
	case err := <-errChanREST:
		log.Fatal("RESTful Microservice encountered an error: ", err)
	default:
		// Block main thread for this time so goroutines can run with their seperate microservices.
		select {}

	}
}
