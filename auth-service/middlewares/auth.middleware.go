package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/utils"
)

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken, err := c.Cookie("Authorization")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Auth token not found"})
			c.Abort()
		}
		token := strings.Fields(authToken)
		tokenString := token[1]
		verified, _ := utils.ValidateToken(tokenString)
		if !verified {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid auth token"})
			c.Abort()
			return
		}
		claims, err := utils.GetClaimsFromToken(tokenString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid auth token"})
			c.Abort()
			return
		}
		c.Set("user_details", claims)
	}
}

func OnlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDetails utils.SignedDetails = c.MustGet("user_details").(utils.SignedDetails)
		if !userDetails.IsAdmin {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Only admin can perform this action"})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
