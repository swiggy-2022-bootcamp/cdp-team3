package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/swiggy-ipp/checkout-service/configs"
	"github.com/swiggy-ipp/checkout-service/grpcs"
	"github.com/swiggy-ipp/checkout-service/grpcs/cart_checkout"
	"github.com/swiggy-ipp/checkout-service/routes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	errChanGRPC chan error = make(chan error)
	errChanKafka chan error = make(chan error)
	errChanREST chan error = make(chan error)
) 

/// Function with logic for GRPC client
func startGRPCClient(address string) {
	// Create a listener on TCP port
	conn, err := grpc.Dial(":" + address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		errChanGRPC <- err
	} else {
		// Start GRPC
		grpcs.CartCheckoutGRPCClient = cart_checkout.NewCheckoutServiceClient(conn)
	}
}

/// Function with logic for starting REST Routes
func generateRESTRoutes(port string) {
	checkoutRouter := gin.Default()
	routes.GenerateCheckoutRoutes(checkoutRouter)

	// Run REST Microservice
	errChanREST <- checkoutRouter.Run(":" + port)
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
	go startGRPCClient(configs.EnvServiceGRPCPort())

	// Set up Kafka listener
	go startKafka(kafkaTopic)

	// Set up routes for Checkout Microservice
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
