package app

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/app/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/app/routes"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/configs"
	_ "github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/domain/services"
	orderGrpc "github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/grpc/order"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/utils"
	"go.uber.org/zap"
)

var (
	server                *gin.Engine
	orderRepository repository.OrderRepository
	orderService    services.OrderService
	orderController    controllers.OrderController
	orderRoutes     routes.OrderRoutes
	orderDB         *dynamodb.DynamoDB
)

func Start() {

	//Initialize Logger
	log := utils.InitializeLogger()

  zap.ReplaceGlobals(log)
  defer log.Sync()
  log.Info("Orders Service Started")

	//Initialize DB
	orderDB := configs.ConnectDB()
	configs.CreateTable(orderDB)
	
	orderRepository = repository.NewOrderRepositoryImpl(orderDB)
	orderService = services. NewOrderServiceImpl(orderRepository)
	orderController = controllers.NewOrderController(orderService)
	orderRoutes = routes.NewOrderRoutes(orderController)

	go orderGrpc.InitializeGRPCServer(configs.EnvGrpcOrderServerPORT())
	router := StartRestServer()
	router.Run(":"+configs.EnvPORT())
}

func StartRestServer() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.HealthCheck())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	orderRoutes.OrdersRoute(router)
	return router
}