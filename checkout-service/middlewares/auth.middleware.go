package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	auth "github.com/swiggy-ipp/checkout-service/grpcs/auth"
)

type SignedDetails struct {
	UserId  string `json:"userId"`
	Name    string `json:"name"`
	EmailId string `json:"emailId"`
	IsAdmin bool   `json:"isAdmin"`
	jwt.StandardClaims
}

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken, err := c.Cookie("token")
		if err != nil {
			fmt.Println("Error in getting cookie: ", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Auth token not found"})
			c.Abort()
			return
		}
		claims, err := auth.VerifyToken(authToken)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth token"})
			c.Abort()
			return
		}
		c.Set("user_details", claims)
		c.Next()
	}
}

func OnlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userDetails SignedDetails = c.MustGet("user_details").(SignedDetails)
		if !userDetails.IsAdmin {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Only admin can perform this action"})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
