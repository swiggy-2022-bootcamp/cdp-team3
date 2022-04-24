package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/app/controllers"
)

type HealthRoutes struct {
	healthController controllers.HealthController
}

func NewHealthRouter(authController controllers.AuthController) AuthRoutes {
	return AuthRoutes{authController: authController}
}

func (HealthRoutes) HealthRoute(router *gin.Engine) {
	public := router.Group("")
	public.GET("/", controllers.HealthCheck())
	public.GET("/deep", controllers.DeepHealthCheck())
}
