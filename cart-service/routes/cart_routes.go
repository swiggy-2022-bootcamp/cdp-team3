package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swiggy-ipp/cart-service/docs"
	"github.com/swiggy-ipp/cart-service/controllers"
)

const BaseURL string = "/cart";

func GenerateCartRoutes(router *gin.Engine) {
	cartRouter := router.Group(BaseURL)

	// Swagger Routes
	docs.SwaggerInfo.BasePath = BaseURL
	cartRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	// Add Cart Item
	cartRouter.POST("/", controllers.CreateCartItem)

	// Get Cart Items
	cartRouter.GET("/", controllers.GetCartItems)

	// Update Cart Item
	cartRouter.PUT("/", controllers.UpdateCartItem)

	// Delete Cart Items
	cartRouter.DELETE("/:key", controllers.DeleteCartItem)

	// Delete Entire Cart
	cartRouter.DELETE("/empty", controllers.EmptyCart)
}