package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/dto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/models"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/utils"
	"go.mongodb.org/mongo-driver/bson"
)

const requestTimeout = time.Second * 5

// Login godoc
// @Summary Login
// @Description This request is used to login a user and get a token in cookies
// @Tags Auth Service
// @Schemes
// @Accept json
// @Produce json
// @Param req body dto.LoginDto true "Login Details"
// @Success	200  {string} 	message
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	401  {number} 	http.StatusUnauthorized
// @Failure	404  {number} 	http.StatusNotFound
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /auth/login [POST]
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		var user dto.LoginDto
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := models.UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User with this email doesn't exists "})
			return
		}

		isValidPassword, msg := utils.VerifyPassword(user.Password, foundUser.Password)

		if !isValidPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
			return
		}
		token, err := utils.CreateToken(foundUser.Id, foundUser.Email, foundUser.Name, foundUser.IsAdmin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		token = "Bearer " + token
		c.SetCookie("token", token, 3600, "/", "", false, true)
		c.JSON(200, gin.H{"message": "logged in"})
	}
}

// Logout godoc
// @Summary Logout
// @Description This request is used to logout a user
// @Tags Auth Service
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {string} 	message
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /auth/logout [POST]
func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "", false, true)
		c.JSON(200, gin.H{"message": "logged out"})
	}
}

// VerifyToken godoc
// @Summary VerifyToken
// @Description This request is used to verify a token internally or by frontend
// @Tags Auth Service
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {object} utils.SignedDetails
// @Failure	401  {number} http.StatusUnauthorized
// @Router /auth/verify-token [POST]
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		isValid, _ := utils.ValidateToken(token)
		if isValid != true {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid"})
			return
		}
		claims, err := utils.GetClaimsFromToken(token)
		if err != nil {
			fmt.Println("error", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid"})
			return
		}
		c.JSON(200, gin.H{"message": "token valid", "claims": claims})
	}
}
