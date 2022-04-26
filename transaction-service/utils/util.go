package utils

import (
	"github.com/gin-gonic/gin"
	adminProto "github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/grpc/admin/proto"
	authproto "github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/grpc/auth/proto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/models"
)

func CheckLoggedInUserWithTransactionCustomerId(c *gin.Context, orderCudtomerId string) bool {
	var userDetails = c.MustGet("user_details").(*authproto.VerifyTokenResponse)
	return userDetails.UserId == orderCudtomerId
}

func IsAdmin(c *gin.Context) bool {
	var userDetails = c.MustGet("user_details").(*authproto.VerifyTokenResponse)
	return userDetails.IsAdmin
}

func ProtoConv(transaction *models.Transaction) *adminProto.TransactionDetails {
	return &adminProto.TransactionDetails{
		UserId:            transaction.CustomerID,
		TransactionAmount: float32(transaction.Amount),
	}
}
