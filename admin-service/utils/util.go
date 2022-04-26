package utils

import (
	"github.com/gin-gonic/gin"
	authProto "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/grpc/auth/proto"
)

func IsAdmin(c *gin.Context) bool {
	var userDetails = c.MustGet("user_details").(*authProto.VerifyTokenResponse)
	return userDetails.IsAdmin
}

func UserId(c *gin.Context) string {
	var userDetails = c.MustGet("user_details").(*authProto.VerifyTokenResponse)
	return userDetails.UserId
}
