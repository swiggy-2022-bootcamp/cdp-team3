package main

import (
	// "context"

	// docs "github.com/swiggy-ipp/cart-service/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swiggy-ipp/cart-service/configs"
)

func main() {
	// docs.SwaggerInfo.BasePath = "/cart-service"
	// Set up routes for Bookkeeping Microservice
	cartRouter := gin.Default()
	cartRouter.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "Okay") })
	cartRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Set up Kafka listener
	// ctx := context.Background()
	// go services.Consume(configs.KafkaDiagnosisTopic(), services.DeserializeAndSaveDiseaseDiagnosis, ctx)
	
	// Run Microservice
	cartRouter.Run(":" + configs.EnvServicePort())
}
