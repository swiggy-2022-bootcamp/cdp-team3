package main

import (
	_ "github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/routes"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/utils"
	"go.uber.org/zap"
)

// @title           Swagger Orders Service API
// @version         1.0
// @description
// @termsOfService  http://swagger.io/terms/

// @contact.name   Jaithun Mahira
// @contact.email  swiggyb1035@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3004

// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization
func main() {

	//Initialize Logger
	log := utils.InitializeLogger()

  zap.ReplaceGlobals(log)
  defer log.Sync()
  log.Info("Orders Service Started")

	//Initialize DB
	ordersDB := configs.ConnectDB()
	configs.CreateTable(ordersDB)
	
	router := gin.Default()

	router.GET("/", controllers.HealthCheck())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.OrdersRoute(router)
	router.Run(":3004")
}