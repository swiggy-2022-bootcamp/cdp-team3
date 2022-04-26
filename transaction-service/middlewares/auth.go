package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/grpc/auth"
	authproto "github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/grpc/auth/proto"
)

type SignedDetails struct {
	UserId  string `json:"userId"`
	Name    string `json:"name"`
	EmailId string `json:"emailId"`
	IsAdmin bool   `json:"isAdmin"`
	jwt.StandardClaims
}

func AuthenticateJWT() gin.HandlerFunc {
	fmt.Println("Inside Authentiicate")
	return func(c *gin.Context) {
		fmt.Println("Inside Authentiicate function")
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
	fmt.Println("Inside only admin")
	return func(c *gin.Context) {
		fmt.Println("Inside only admin function")
		var userDetails = c.MustGet("user_details").(*authproto.VerifyTokenResponse)
		fmt.Println("Userdetails in only admin", userDetails)
		if !userDetails.IsAdmin {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Only admin can perform this action"})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
