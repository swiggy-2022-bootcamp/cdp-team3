package middlewares

import (
	"fmt"
	"net/http"
	auth "github.com/cdp-team3/categories-service/app/grpcs/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
    services "github.com/cdp-team3/categories-service/domain/services"
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
		claims, err := services.VerifyToken(authToken)

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
		var userDetails = c.MustGet("user_details").(*auth.VerifyTokenResponse)
		fmt.Println("Userdetails in only admin",userDetails)
		if !userDetails.IsAdmin {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Only admin can perform this action"})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}