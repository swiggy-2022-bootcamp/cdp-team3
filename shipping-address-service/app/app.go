package app

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"fmt"
	//"time"
	"context"
	"net"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	//"os/signal"
	//"github.com/joho/godotenv"
	"github.com/cdp-team3/shipping-address-service/app/handlers"
	"github.com/cdp-team3/shipping-address-service/app/routes"
    "github.com/cdp-team3/shipping-address-service/config"
	"github.com/cdp-team3/shipping-address-service/domain/repository"
	"github.com/cdp-team3/shipping-address-service/domain/services"
	"github.com/cdp-team3/shipping-address-service/utils/logger"
    "github.com/cdp-team3/shipping-address-service/app/grpcs/shipping_checkout"
	"google.golang.org/grpc"
	"sync"
	//"syscall"
	//"google.golang.org/grpc/reflection"

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
// // gracefulStop logic to allow go routines to finish
// var gracefulStop = make(chan os.Signal)
// signal.Notify(gracefulStop, syscall.SIGTERM)
// signal.Notify(gracefulStop, syscall.SIGINT)
// go func() {
// 	sig := <-gracefulStop
// 	logger.Info("caught sig: %+v", sig)
// 	logger.Info("Wait for 2 second to finish processing")
// 	time.Sleep(2 * time.Second)
// 	os.Exit(0)
// }()

// wg.Add(1)
// go StartRESTServer()
// // go StartGRPCServer()
// wg.Wait()
go startGRPCServer(os.Getenv("GRPC_ADDRESSS_PORT"))
go StartRESTServer()
// Listen to errors.
select {
case err := <-errChanGRPC:
fmt.Println("GRPC encountered an error: ", err)
case err := <-errChanKafka:
	fmt.Println("Kafka encountered an error: ", err)
case err := <-errChanREST:
	fmt.Println("RESTful Microservice encountered an error: ", err)
default:
	// Block main thread for this time so goroutines can run with their seperate microservices.
	select {}

}
 
}

func startGRPCServer(port string) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		fmt.Println("Failed to listen: %v", err)
		errChanGRPC <- err
	} else {
		// Start GRPC
		errChanGRPC <- StartGRPCServer1(lis)
	}
}
func StartGRPCServer1(lis net.Listener) error {
	// Create a gRPC server object
	//s := grpc.NewServer()
	grpcServer := grpc.NewServer()
	shipping_checkout.RegisterShippingServer(grpcServer, shippingServiceProto)
//	cart_checkout.RegisterCheckoutServiceServer(s, &server{})
	fmt.Printf("Server listening at %v", lis.Addr())
	// Start serving requests
	err := grpcServer.Serve(lis)
	if err != nil {
		fmt.Println("Error in starting grpc server",err)
		logger.Error("Failed to start the grpc server : %v", err)
	}
	if err==nil {fmt.Printf("Server started at %v", lis.Addr())}
	return nil
}
func StartRESTServer(){
	fmt.Println("Inside creating REST server")
	//Configuring gin server and router

	server = gin.New()
	server.Use(logger.UseLogger(logger.DefaultLoggerFormatter), gin.Recovery())
	router := server.Group("shipping-service/api")
	shippingRoutes.InitRoutes(router)

	//Starting server on port 3003
    err := server.Run(":3003")
	fmt.Println("going to start server")
    if err != nil {
		fmt.Println("Shipping Address Server can not be started.",err)
	     logger.Error(err.Error() + " - Failed to start server")
    } else {
	logger.Info("Shipping Address Server started successfully.")
	fmt.Println("Shipping Address Server started successfully.")
    }
}

func StartGRPCServer() {
	fmt.Println("Inside creating GRPC server")
	GRPC_ADDRESSS_PORT := os.Getenv("GRPC_ADDRESSS_PORT")
	lis, err := net.Listen("tcp", GRPC_ADDRESSS_PORT)
	if err != nil {
		fmt.Println("failed to listen on port")
		logger.Error("Failed to listen on port %v with error %v", GRPC_ADDRESSS_PORT, err)
	}
	
	grpcServer := grpc.NewServer()
	shipping_checkout.RegisterShippingServer(grpcServer, shippingServiceProto)
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println("Error in starting grpc server",err)
		logger.Error("Failed to start the grpc server : %v", err)
	}
	logger.Info("GRPC server started successfully")
	fmt.Println("GRPC server started successfully")
	
}