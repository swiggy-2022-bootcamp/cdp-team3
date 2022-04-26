package utils

import (
	"github.com/gin-gonic/gin"
	adminProto "github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/grpc/admin/proto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/middlewares"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/models"
)

func CheckLoggedInUserWithTransactionCustomerId(c *gin.Context , orderCudtomerId string) bool {
	var userDetails middlewares.SignedDetails = c.MustGet("user_details").(middlewares.SignedDetails)
	return userDetails.UserId == orderCudtomerId
}

func IsAdmin(c *gin.Context) bool {
	var userDetails middlewares.SignedDetails = c.MustGet("user_details").(middlewares.SignedDetails)
	return userDetails.IsAdmin
}


func ProtoConv(transaction *models.Transaction) *adminProto.TransactionDetails {
	return &adminProto.TransactionDetails{
		UserId: transaction.CustomerID,
		TransactionAmount: float32(transaction.Amount),
	}
}