package app

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	//"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"github.com/cdp-team3/shipping-address-service/app/handlers"
	"github.com/cdp-team3/shipping-address-service/app/routes"
    "github.com/cdp-team3/shipping-address-service/config"
	"github.com/cdp-team3/shipping-address-service/domain/repository"
	"github.com/cdp-team3/shipping-address-service/domain/services"
	"github.com/cdp-team3/shipping-address-service/utils/logger"
	//"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
//	apperrors "github.com/cdp-team3/categories-service/app-errors"
	//"github.com/cdp-team3/categories-service/domain/models"
	//"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/service/dynamodb"
)

const shippingCollection = "shippingCollection_Team3"
 var (
	server                *gin.Engine
	shippingRepository repository.ShippingRepository
	shippingService    services.ShippingService
	shippingHandler    handlers.ShippingHandler
	shippingRoutes     routes.ShippingRoutes
	shippingDB         *dynamodb.DynamoDB
	healthCheckHandler    handlers.HealthCheckHandler
)

func Start() {
//Variable initializations for DynamoDB
shippingDB = config.ConnectDB()
config.CreateTable(shippingDB)

// //Variable initializations to be used as dependency injections
shippingRepository = repository.NewShippingRepositoryImpl(shippingDB)
shippingService = services. NewShippingServiceImpl(shippingRepository)
shippingHandler = handlers.NewShippingHandler(shippingService)
 healthCheckHandler = handlers.NewHealthCheckHandler(shippingRepository)
 shippingRoutes = routes.NewShippingRoutes(shippingHandler, healthCheckHandler)

//Opening file for log collection
file, err := os.OpenFile("shipping-address-service-server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err == nil {
	gin.DefaultWriter = io.MultiWriter(file)
}

 server = gin.New()
 server.Use(logger.UseLogger(logger.DefaultLoggerFormatter), gin.Recovery())
 router := server.Group("shipping/api")
 shippingRoutes.InitRoutes(router)

//Starting server on port 3003
err = server.Run(":3003")
if err != nil {
	logger.Error(err.Error() + " - Failed to start server")
} else {
	logger.Info("Shipping Address Server started successfully.")
}
}
