package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/models"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/utils"
	"go.mongodb.org/mongo-driver/bson"
)

const requestTimeout = time.Second * 5

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()

		var user models.User
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
		c.SetCookie("token", token, 3600, "/", "", false, true)
		c.JSON(200, gin.H{"message": "logged in"})
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "", false, true)
		c.JSON(200, gin.H{"message": "logged out"})
	}
}
