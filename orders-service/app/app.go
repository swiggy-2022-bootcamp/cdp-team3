package app

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/controllers"
	_ "github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/docs"
	orderGrpc "github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/grpc/order"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/routes"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/utils"
	"go.uber.org/zap"
)

func Start() {

	//Initialize Logger
	log := utils.InitializeLogger()

  zap.ReplaceGlobals(log)
  defer log.Sync()
  log.Info("Orders Service Started")

	//Initialize DB
	ordersDB := configs.ConnectDB()
	configs.CreateTable(ordersDB)
	
	go orderGrpc.InitializeGRPCServer(configs.EnvGrpcPORT())
	router := StartRestServer()
	router.Run(":3004")
}

func StartRestServer() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.HealthCheck())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.OrdersRoute(router)
	return router
}