package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/middlewares"
)

func PaymentRoutes(router *gin.Engine) {
	private := router.Group("/pay")
	private.Use(middlewares.AuthenticateJWT())
	private.POST("", controllers.Pay())
}
