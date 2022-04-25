package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/app/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/middlewares"
)

type PaymentRoutes struct {
	paymentController controllers.PaymentController
}

func NewPaymentRouter(pc controllers.PaymentController) PaymentRoutes {
	return PaymentRoutes{paymentController: pc}
}

func (pr PaymentRoutes) PaymentRoute(router *gin.Engine) {
	private := router.Group("/pay")
	private.Use(middlewares.AuthenticateJWT())
	private.POST("", pr.paymentController.Pay())
}
