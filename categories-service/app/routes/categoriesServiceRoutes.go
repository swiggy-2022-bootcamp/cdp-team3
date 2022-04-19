package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/cdp-team3/categories-service/app/handlers"
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

	newRouter.POST("/categories", tr.categoriesHandler.AddCategory)
	newRouter.GET("/categories", tr.categoriesHandler.GetAllCategory)
	newRouter.GET("/categories/:category_id",  tr.categoriesHandler.GetCategory)
	newRouter.DELETE("/categories/",  tr.categoriesHandler.DeleteCategories)
	newRouter.PUT("/categories/:category_id",  tr.categoriesHandler.UpdateCategory)
	
}