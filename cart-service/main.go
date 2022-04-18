package main

import (
	"net"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/swiggy-ipp/cart-service/configs"
	"github.com/swiggy-ipp/cart-service/grpcs"
	"github.com/swiggy-ipp/cart-service/routes"
)

var (
	errChanGRPC chan error = make(chan error)
	errChanKafka chan error = make(chan error)
	errChanREST chan error = make(chan error)
) 

/// Function with logic for starting GRPC server
func startGRPCServer(port string) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
		errChanGRPC <- err
	} else {
		// Start GRPC
		errChanGRPC <- grpcs.StartGRPCServer(lis)
	}
}

/// Function with logic for starting REST Routes
func generateRESTRoutes(port string) {
	cartRouter := gin.Default()
	routes.GenerateCartRoutes(cartRouter)

	// Run REST Microservice
	errChanREST <- cartRouter.Run(":" + port)
}

/// Function with logic for starting Kafka listener
func startKafka(topic string) {
	// Set up Kafka listener
	// ctx := context.Background()
	// go services.Consume(topic, services.DeserializeAndSaveDiseaseDiagnosis, ctx)
}

func main() {
	// Get configs
	kafkaTopic := ""

	// Set up GRPC
	go startGRPCServer(configs.EnvServiceGRPCPort())

	// Set up Kafka listener
	go startKafka(kafkaTopic)

	// Set up routes for Cart Microservice
	go generateRESTRoutes(configs.EnvServicePort())

	// Listen to errors.
	select {
		case err := <-errChanGRPC:
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
