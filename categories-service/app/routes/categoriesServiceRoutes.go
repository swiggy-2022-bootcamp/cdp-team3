package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/cdp-team3/categories-service/app/handlers"
	"github.com/cdp-team3/categories-service/middlewares"
	_ "github.com/cdp-team3/categories-service/docs"
)
type CategoriesRoutes struct {
	categoriesHandler handlers.CategoryHandler
	healthCheckhandler handlers.HealthCheckHandler
}
func NewCategoryRoutes(categoriesHandler handlers.CategoryHandler, healthCheckhandler handlers.HealthCheckHandler) CategoriesRoutes {
	return CategoriesRoutes{categoriesHandler: categoriesHandler, healthCheckhandler: healthCheckhandler}
}
func (tr CategoriesRoutes) InitRoutes(newRouter *gin.RouterGroup) {

	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", tr.healthCheckhandler.HealthCheck)
	
	// middleware that check if the user is admin and logged in
	newRouter.Use(middlewares.AuthenticateJWT())
	newRouter.Use(middlewares.OnlyAdmin())

	newRouter.POST("/categories", tr.categoriesHandler.AddCategory())
	newRouter.GET("/categories", tr.categoriesHandler.GetAllCategory())
	newRouter.GET("/categories/:category_id",  tr.categoriesHandler.GetCategory())
	newRouter.DELETE("/categories/:category_id",  tr.categoriesHandler.DeleteCategory())
    newRouter.DELETE("/categories/",  tr.categoriesHandler.DeleteCategories())
	newRouter.PUT("/categories/:category_id",  tr.categoriesHandler.UpdateCategory())
	
}