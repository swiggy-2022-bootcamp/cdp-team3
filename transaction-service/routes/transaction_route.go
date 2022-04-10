package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/controllers"
)

func TransactionRoute(router *gin.Engine) {
	router.GET("/transaction/:customerId", controllers.GetTransactionByCustomerId())
	router.POST("/transaction/:customerId", controllers.AddTransactionAmtToCustomer())
}
