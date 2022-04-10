package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// AddTransactionAmtToCustomer godoc
// @Summary Adds transaction amount
// @Description This request will add transaction amount to a customer by customer ID
// @Tags Transacion Service
// @Schemes
// @Accept json
// @Produce json
// @Param customerId path string true "Customer Id"
// @Param req body models.Transaction true "Transaction details"
// @Success	201  {object} 	models.Transaction
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /transaction/{customerId} [POST]
func AddTransactionAmtToCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		c.JSON(http.StatusCreated, gin.H{
			"message": "Transaction amount added successfully",
		})
	}
}

// GetTransactionByCustomerId godoc
// @Summary Get transaction amount for a customer.
// @Description Get the total transaction amount credited to the customer based on customer ID.
// @Tags Transaction Service
// @Schemes
// @Param customerId path string true "Customer Id"
// @Produce json
// @Success	200  {object} models.Transaction
// @Failure	500  {number} http.StatusInternalServerError
// @Security Bearer Token
// @Router /transaction/{customerId} [GET]
func GetTransactionByCustomerId() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		c.JSON(http.StatusCreated, gin.H{
			"amount": "2000",
		})
	}
}

// HealthCheck godoc
// @Summary To check if the service is running or not.
// @Description This request will return 200 OK if server is up..
// @Tags Health
// @Schemes
// @Accept json
// @Produce json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router / [GET]
func HealthCheck() gin.HandlerFunc {
	//Check to be added for database.
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Transaction Service is running"})
	}
}
