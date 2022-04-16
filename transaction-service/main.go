package main

import (
	_ "github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/routes"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/utils"

	"go.uber.org/zap"
)

// @title           Swagger Transaction Service API
// @version         1.0
// @description     Admin can add and view transaction amount for a customer based on custmer ID
// @termsOfService  http://swagger.io/terms/

// @contact.name   Jaithun Mahira
// @contact.email  swiggyb1035@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3005

// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization
func main() {

	//Initialize Logger
	log := utils.InitializeLogger()

  zap.ReplaceGlobals(log)
  defer log.Sync()
  log.Info("Transaction Service Started")
	
	router := gin.Default()

	router.GET("/", controllers.HealthCheck())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.TransactionRoute(router)
	router.Run(":3005")
}
