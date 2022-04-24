package main

import (
	"net"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/swiggy-ipp/cart-service/configs"
	"github.com/swiggy-ipp/cart-service/configs/database"
	"github.com/swiggy-ipp/cart-service/controllers"
	"github.com/swiggy-ipp/cart-service/grpcs"
	"github.com/swiggy-ipp/cart-service/repositories"
	"github.com/swiggy-ipp/cart-service/routes"
	"github.com/swiggy-ipp/cart-service/services"
	"github.com/swiggy-ipp/cart-service/utils"
)

var (
	errChanGRPC  chan error = make(chan error)
	errChanKafka chan error = make(chan error)
	errChanREST  chan error = make(chan error)
)

/// Function with logic for starting GRPC server
func startGRPCServer(address string, cartService services.CartService) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", "0.0.0.0:" + address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		errChanGRPC <- err
	} else {
		// Start GRPC
		errChanGRPC <- grpcs.NewCartCheckoutGRPCServer(lis, cartService)
	}
}

/// Function with logic for starting REST Routes
func generateRESTRoutes(port string, cartController controllers.CartController) {
	cartRouter := gin.Default()
	routes.GenerateCartRoutes(cartRouter, cartController)

	// Run REST Microservice
	errChanREST <- cartRouter.Run("0.0.0.0:" + port)
}

func main() {
	// Make layered Architecture
	db := database.GetDynamoDBClient()                           // Database
	cartRepository := repositories.NewCartRepository(db, "cart") // Repository
	cartService := services.NewCartService(cartRepository)       // Service
	cartController := controllers.NewCartController(cartService) // Controller

	// Set up GRPC
	go startGRPCServer(configs.EnvServiceGRPCPort(), cartService)

	// Set up Kafka listeners
	kafkaCartDeleteListener := utils.NewKafkaCartConsumeService(
		configs.EnvKafkaBrokerAddress(),
		configs.EnvKafkaUserDeletedTopic(),
		cartService,
	) // Kafka Listener
	go kafkaCartDeleteListener.KafkaUserIDConsume()

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
