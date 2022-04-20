package main

import (
	"context"
	"net"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/swiggy-ipp/cart-service/configs"
	"github.com/swiggy-ipp/cart-service/controllers"
	"github.com/swiggy-ipp/cart-service/grpcs"
	"github.com/swiggy-ipp/cart-service/models"
	"github.com/swiggy-ipp/cart-service/repositories"
	"github.com/swiggy-ipp/cart-service/routes"
)

var (
	errChanGRPC chan error = make(chan error)
	errChanKafka chan error = make(chan error)
	errChanREST chan error = make(chan error)
) 

/// Function with logic for starting GRPC server
func startGRPCServer(port string, cartRepository repositories.CartRepository) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
		errChanGRPC <- err
	} else {
		// Start GRPC
		errChanGRPC <- grpcs.NewCartCheckoutGRPCServer(lis, cartRepository)
	}
}

/// Function with logic for starting REST Routes
func generateRESTRoutes(port string, cartController controllers.CartController) {
	cartRouter := gin.Default()
	routes.GenerateCartRoutes(cartRouter, cartController)

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

	// Make layered Architecture
	db := configs.GetDynamoDBClient(); // Database
	cartRepository := repositories.NewCartRepository(db, "cart") // Repository
	cartController := controllers.NewCartController(cartRepository) // Controller
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cartRepository.Create(ctx, &models.Cart{
		UserID: "user1",
		Items: []models.CartItem{
			{
				ProductID: "product1",
				Quantity:  "1",
			},
		},
	})

	// Set up GRPC
	go startGRPCServer(configs.EnvServiceGRPCPort(), cartRepository)

	// Set up Kafka listener
	go startKafka(kafkaTopic)

	// Set up routes for Cart Microservice
	go generateRESTRoutes(configs.EnvServicePort(), cartController)

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
