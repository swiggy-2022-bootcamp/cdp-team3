package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/dto"
	auth "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/grpc/auth"
	"go.uber.org/zap"
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
			zap.L().Error("Error in getting cookie:" + err.Error())
			c.JSON(http.StatusUnauthorized, dto.ResponseDTO{
				Status:  http.StatusUnauthorized,
				Message: "Auth token not found",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			c.Abort()
			return
		}

		claims, err := auth.VerifyToken(authToken)

		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ResponseDTO{
				Status:  http.StatusUnauthorized,
				Message: "Invalid auth token",
				Data:    map[string]interface{}{"data": err.Error()},
			})
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
			c.JSON(http.StatusUnauthorized, dto.ResponseDTO{
				Status:  http.StatusUnauthorized,
				Message: "Only admin can perform this action",
			})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
