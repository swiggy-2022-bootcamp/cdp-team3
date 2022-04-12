package main

import (
	_ "github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/routes"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/utils"
)

var logger = utils.NewLoggerService("main")

// @title           Swagger Auth Service API
// @version         1.0
// @description     Users ( Admin, Customer, etc ) can login and get a token and use it to access other APIs
// @termsOfService  http://swagger.io/terms/

// @contact.name   Rishabh Mishra
// @contact.email  swiggyb2026@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3012

// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in cookie
// @name Authorization
func main() {
	router := gin.Default()

	configs.ConnectDB()
	logger.Log("Connected to DB")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.AuthRoutes(router)
	routes.HealthCheckRoutes(router)

	router.Run("localhost:" + configs.EnvPORT())
}
