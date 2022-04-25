package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/app/controllers"
)

type HealthRoutes struct {
	healthController controllers.HealthController
}

func NewHealthRouter(hc controllers.HealthController) HealthRoutes {
	return HealthRoutes{healthController: hc}
}

func (hr HealthRoutes) HealthRoute(router *gin.Engine) {
	public := router.Group("")
	public.GET("/", hr.healthController.HealthCheck())
	public.GET("/deep", hr.healthController.DeepHealthCheck())
}
