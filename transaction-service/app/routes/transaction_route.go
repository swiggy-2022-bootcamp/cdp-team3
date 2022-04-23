package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/app/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/middlewares"
)

type TransactionRoutes struct {
	transactionController controllers.TransactionController
}

func NewTransactionRoutes(transactionController controllers.TransactionController) TransactionRoutes {
	return TransactionRoutes{transactionController: transactionController}
}

func (tr TransactionRoutes)TransactionRoutes(router *gin.Engine) {

	router.Use(middlewares.AuthenticateJWT())
	router.GET("/transaction/:customerId", tr.transactionController.GetTransactionByCustomerId())
	router.POST("/transaction/:customerId", middlewares.OnlyAdmin(), tr.transactionController.AddTransactionAmtToCustomer())
}
