package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swiggy-ipp/cart-service/docs"
	"github.com/swiggy-ipp/cart-service/controllers"
)

const BaseURL string = "/cart_service";

func GenerateCartRoutes(router *gin.Engine) {
	cartServiceRouter := router.Group(BaseURL)
	cartServiceCartRouter := cartServiceRouter.Group("/cart")

	// Swagger Routes
	docs.SwaggerInfo.BasePath = BaseURL
	cartServiceRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	// Add Cart Item
	cartServiceCartRouter.POST("/", controllers.CreateCartItem)

	// Get Cart Items
	cartServiceCartRouter.GET("/", controllers.GetCartItems)

	// Update Cart Item
	cartServiceCartRouter.PUT("/", controllers.UpdateCartItem)

	// Delete Cart Items
	cartServiceCartRouter.DELETE("/:key", controllers.DeleteCartItem)

	// Delete Entire Cart
	cartServiceCartRouter.DELETE("/empty", controllers.EmptyCart)
}