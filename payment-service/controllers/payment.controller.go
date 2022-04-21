package controllers

import (
	"time"

	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/utils"
)

const requestTimeout = time.Second * 5

var logger = utils.NewLoggerService("payment-controller")

// Pay godoc
// @Summary Payment
// @Description This request is used for Payment
// @Tags Payment Service
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {string} 	message
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	401  {number} 	http.StatusUnauthorized
// @Failure	404  {number} 	http.StatusNotFound
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /pay [POST]
func Pay() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Log("Payment controller called")
		c.JSON(200, gin.H{
			"message": "Payment controller called",
		})
	}
}
