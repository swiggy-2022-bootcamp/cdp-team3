package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swiggy-2022-bootcamp/cdp-team3/mode-of-payment-service/configs"

	_ "github.com/swiggy-2022-bootcamp/cdp-team3/mode-of-payment-service/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/mode-of-payment-service/routes"
	"github.com/swiggy-2022-bootcamp/cdp-team3/mode-of-payment-service/utils"
	"go.uber.org/zap"
)

// @title           BuyItNow Payment Mode Service
// @version         1.0
// @description
// @termsOfService  http://swagger.io/terms/

// @contact.name   Aman Gupta
// @contact.email  swiggyb1010@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3007

// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization
func main() {

	//Initialize Logger
	log := utils.InitializeLogger()

	zap.ReplaceGlobals(log)
	defer log.Sync()
	log.Info("Payment Method Service Started")

	router := StartServer()
	router.Run(":" + configs.EnvPORT())
}

func StartServer() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.PaymentModeRoutes(router)
	return router
}
