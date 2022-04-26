package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/middlewares"
)

func IsAdmin(c *gin.Context) bool {
	var userDetails middlewares.SignedDetails = c.MustGet("user_details").(middlewares.SignedDetails)
	return userDetails.IsAdmin
}
