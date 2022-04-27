package app

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"fmt"
	"context"
	"net"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"github.com/cdp-team3/shipping-address-service/app/handlers"
	"github.com/cdp-team3/shipping-address-service/app/routes"
    "github.com/cdp-team3/shipping-address-service/config"
	"github.com/cdp-team3/shipping-address-service/domain/repository"
	"github.com/cdp-team3/shipping-address-service/domain/services"
	"github.com/cdp-team3/shipping-address-service/utils/logger"
    "github.com/cdp-team3/shipping-address-service/app/grpcs/shipping"
	"google.golang.org/grpc"
	"sync"

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
	ctx                 context.Context
	shippingServiceProto services.ShippingProtoServer
	wg                  sync.WaitGroup
)
var (
	errChanGRPC chan error = make(chan error)
	errChanKafka chan error = make(chan error)
	errChanREST chan error = make(chan error)
) 

func Start() {
	ctx = context.TODO()

//Variable initializations for DynamoDB
shippingDB = config.ConnectDB()
config.CreateTable(shippingDB)

// //Variable initializations to be used as dependency injections
shippingRepository = repository.NewShippingRepositoryImpl(shippingDB)
shippingService = services.NewShippingServiceImpl(shippingRepository)
shippingHandler = handlers.NewShippingHandler(shippingService)
 healthCheckHandler = handlers.NewHealthCheckHandler(shippingRepository)
 shippingRoutes = routes.NewShippingRoutes(shippingHandler, healthCheckHandler)
 shippingServiceProto = services.NewShippingProtoService(shippingRepository)


 //Opening file for log collection
file, err := os.OpenFile("shipping-address-service-server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err == nil {
	gin.DefaultWriter = io.MultiWriter(file)
}

go startGRPCServer(os.Getenv("SHIPPING_SERVICE_GRPC_PORT"))
go StartRESTServer()

// Listen to errors.
select {
case err := <-errChanGRPC:
fmt.Println("GRPC encountered an error: ", err)

case err := <-errChanREST:
	fmt.Println("RESTful Microservice encountered an error: ", err)
default:
	// Block main thread for this time so goroutines can run with their seperate microservices.
	select {}

}
 
}

func startGRPCServer(port string) {
	// Create a listener on TCP port
	
	lis, err := net.Listen("tcp",config.EnvShippingHost()+":"+config.EnvShippingServiceGRPCPort())
	if err != nil {
		fmt.Println("Failed to listen: ", err)
		errChanGRPC <- err
	} else {
		// Start GRPC
		errChanGRPC <- StartGRPCServer1(lis)
	}
}
func StartGRPCServer1(lis net.Listener) error {
	// Create a gRPC server object

	grpcServer := grpc.NewServer()
	shipping.RegisterShippingServer(grpcServer, shippingServiceProto)


	// Start serving requests
	err := grpcServer.Serve(lis)
	if err != nil {
	
		logger.Error("Failed to start the grpc server : %v", err)
	}
	if err==nil {fmt.Printf("Server started at %v", lis.Addr())}
	return nil
}
func StartRESTServer(){

	//Configuring gin server and router

	server = gin.New()
	server.Use(logger.UseLogger(logger.DefaultLoggerFormatter), gin.Recovery())
	router := server.Group("shipping-service/api")
	shippingRoutes.InitRoutes(router)

	//Starting server on port 3003
    err := server.Run(":"+config.EnvShippingPort())

    if err != nil {
		
	     logger.Error(err.Error() + " - Failed to start server")
    } else {
	logger.Info("Shipping Address Server started successfully.")

}
