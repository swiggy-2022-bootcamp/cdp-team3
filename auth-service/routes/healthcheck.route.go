package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/controllers"
)

func HealthCheckRoutes(router *gin.Engine) {
	public := router.Group("/")
	public.GET("/", controllers.HealthCheck())
	public.GET("/deep", controllers.DeepHealthCheck())
}
