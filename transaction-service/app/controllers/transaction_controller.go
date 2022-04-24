package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/domain/services"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/dto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/kafka"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/models"
	"go.uber.org/zap"
)

func init() {
	go kafka.AddTransactionAmountConsumer()
}
type TransactionController struct {
	transactionService services.TransactionService
}

func NewTransactionController(transactionService services.TransactionService) TransactionController {
	return TransactionController{transactionService : transactionService}
}


var validate = validator.New()

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
func (tc TransactionController) AddTransactionAmtToCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside AddTransactionAmtToCustomer Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		customerId := c.Param("customerId")
		transactionFromClient := &models.Transaction {
			CustomerID: customerId,
		}

		//validate the request body
		if err := c.BindJSON(&transactionFromClient); err != nil {
			zap.L().Error("Error validating the request body"+err.Error())
			c.JSON(http.StatusBadRequest, dto.ResponseDTO{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		transaction, err := tc.transactionService.AddTransactionAmtToCustomer(transactionFromClient)
		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(err.Code, dto.ResponseDTO{
				Status:  err.Code,
				Message: err.Message,
			})
			return
		}
		zap.L().Info("Successfully added transaction to customer"+customerId)
		c.JSON(http.StatusCreated, dto.ResponseDTO{
			Status: http.StatusCreated, 
			Message: "success", 
			Data: map[string]interface{}{"transaction": transaction},
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
func (tc TransactionController) GetTransactionByCustomerId() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetTransactionByCustomerId Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		customerId := c.Param("customerId")

		// if !utils.IsAdmin(c) && !utils.CheckLoggedInUserWithTransactionCustomerId(c, customerId) {
		// 	zap.L().Error("Not authorized to view transaction points of customer")
		// 	c.JSON(http.StatusUnauthorized, dto.ResponseDTO{
		// 		Status:  http.StatusUnauthorized,
		// 		Message: "Not authorized to view transaction points of customer",
		// 	})
		// }

		transactionList, err := tc.transactionService.GetTransactionByCustomerId(customerId)

		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(err.Code, dto.ResponseDTO{
				Status:  err.Code,
				Message: err.Message,
			})
			return
		}

		zap.L().Info("Fetched all transaction for customer " + customerId + "successfully")

		var totalTransactionAmount float64;
		for _, trans := range transactionList {
			totalTransactionAmount += trans.Amount
		}
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{
				"transactions": transactionList,
				"totalAmount": totalTransactionAmount,
			},
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
// @Failure      500  {number} 	http.StatusInternalServerError
// @Router / [GET]
func HealthCheck() gin.HandlerFunc {

	//Ping DB
	_, err := configs.DB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		zap.L().Error("Database connection is down.")
		return func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, dto.HealthCheckResponse{Server: "Server is up", Database: "Database is down"})
		}
	}
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, dto.HealthCheckResponse{Server: "Server is up", Database: "Database is up"})
	}
}