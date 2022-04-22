package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swiggy-ipp/checkout-service/controllers"
	"github.com/swiggy-ipp/checkout-service/docs"
	"github.com/swiggy-ipp/checkout-service/middlewares"
)

const BaseURL string = "/checkout_service"

func GenerateCheckoutRoutes(router *gin.Engine, checkoutController controllers.CheckoutController) {
	checkoutServiceRouter := router.Group(BaseURL)
	checkoutServiceRouter.Use(middlewares.AuthenticateJWT())
	checkoutServiceConfirmRouter := checkoutServiceRouter.Group("/confirm")

	// Health Check
	checkoutServiceRouter.GET("/", checkoutController.HealthCheck)

	// Swagger Routes
	docs.SwaggerInfo.BasePath = BaseURL
	checkoutServiceRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Get an overview of the order
	checkoutServiceConfirmRouter.GET("/", checkoutController.GetOrderOverview)
	checkoutServiceConfirmRouter.POST("/", checkoutController.GetOrderOverview)

	// Clear Cart and Unset Session Data
	checkoutServiceConfirmRouter.POST("/success", checkoutController.OrderCompleteWebhook)
}
