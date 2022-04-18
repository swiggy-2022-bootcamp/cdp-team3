package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/cdp-team3/categories-service/app/handlers"
	_ "github.com/cdp-team3/categories-service/docs"
)

func InitRoutes(router *gin.Engine) {
	newRouter := router.Group("categories-service/api")
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", handlers.HealthCheck())
}