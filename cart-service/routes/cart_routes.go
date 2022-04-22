package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swiggy-ipp/cart-service/controllers"
	"github.com/swiggy-ipp/cart-service/docs"
	"github.com/swiggy-ipp/cart-service/middlewares"
)

const BaseURL string = "/cart_service"

func GenerateCartRoutes(router *gin.Engine, cartController controllers.CartController) {
	cartServiceRouter := router.Group(BaseURL)
	cartServiceRouter.Use(middlewares.AuthenticateJWT())
	cartServiceCartRouter := cartServiceRouter.Group("/cart")

	// Health Check
	cartServiceRouter.GET("/", cartController.HealthCheck)

	// Swagger Routes
	docs.SwaggerInfo.BasePath = BaseURL
	cartServiceRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Add Cart Item
	cartServiceCartRouter.POST("/", cartController.CreateCartItem)

	// Get Cart Items
	cartServiceCartRouter.GET("/", cartController.GetCartItems)

	// Update Cart Item
	cartServiceCartRouter.PUT("/", cartController.UpdateCartItem)

	// Delete Cart Items
	cartServiceCartRouter.DELETE("/:key", cartController.DeleteCartItem)

	// Delete Entire Cart
	cartServiceCartRouter.DELETE("/empty", cartController.EmptyCart)
}
