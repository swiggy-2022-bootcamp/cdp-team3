package main

import (
	docs "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/routes"
)

// @title           BuyItNow Admin Services
// @version         1.0
// @description     Admin can get their details and create, read, update and delete customers.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Uttej Immadi
// @contact.email  swiggyb3014@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3009

// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization
func main() {
	router := gin.Default()

	router.GET("/", controllers.HealthCheck)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.Title = "BuyItNow Admin Service"

	routes.AdminRoute(router)
	router.Run(":3009")
}
