package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/products-service/controllers"
)

func ProductAdminRoutes(router *gin.Engine) {
	router.GET("/products", controllers.GetAllProducts())
	router.POST("/products", controllers.AddProduct())
	router.PUT("/products/:productId", controllers.UpdateProduct())
	router.DELETE("/products/:productId", controllers.DeleteProduct())
	router.GET("/products/search/:search", controllers.SearchProducts())

}
