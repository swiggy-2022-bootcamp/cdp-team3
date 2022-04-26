package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/mode-of-payment-service/controllers"
)

func PaymentModeRoutes(r *gin.Engine) {
	r.GET("/paymentmethods", controllers.GetAllPaymentMethods())

}
