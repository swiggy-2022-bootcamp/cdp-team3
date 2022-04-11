package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/swiggy-2022-bootcamp/cdp-team3/products-service/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/products-service/routes"
)

var port int

// @title           Swagger Products Service API
// @version         1.0
// @description     Admin can add, update, delete and search Products into DB and User can fetch products
// @termsOfService  http://swagger.io/terms/

// @contact.name   Sai Kumar Basaveswara
// @contact.email  swiggyb2030@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000

// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization
func main() {
	port = 3000
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	routes.ProductAdminRoutes(router)

	if err := router.Run(fmt.Sprintf(":%v", port)); err != nil {
		handleError(fmt.Sprintf("Unable to start on port %v. The error is : %v", port, err))
	}

}

func handleError(errText string) {
	log.Fatalf(errText)
}
