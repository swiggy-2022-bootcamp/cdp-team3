package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swiggy-ipp/checkout-service/docs"
	"github.com/swiggy-ipp/checkout-service/controllers"
)

const BaseURL string = "/checkout_service";

func GenerateCheckoutRoutes(router *gin.Engine) {
	checkoutServiceRouter := router.Group(BaseURL)
	checkoutServiceConfirmRouter := checkoutServiceRouter.Group("/confirm")

	// Swagger Routes
	docs.SwaggerInfo.BasePath = BaseURL
	checkoutServiceRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	// Get an overview of the order
	checkoutServiceConfirmRouter.GET("/", controllers.GetOrderOverview)
	checkoutServiceConfirmRouter.POST("/", controllers.GetOrderOverview)

	// Clear Cart and Unset Session Data
	checkoutServiceConfirmRouter.PUT("/success", controllers.OrderCompleteWebhook)
}