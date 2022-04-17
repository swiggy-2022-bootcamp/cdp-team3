package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/dto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/models"
	"go.uber.org/zap"
)

const transactionCollection = "Transaction"
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
func AddTransactionAmtToCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId := c.Param("customerId")
		zap.L().Info("Inside AddTransactionAmtToCustomer Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var transaction models.Transaction
		transaction.CustomerID = customerId
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&transaction); err != nil {
			zap.L().Error("Error validating the request body"+err.Error())
			c.JSON(http.StatusBadRequest, dto.ResponseDTO{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&transaction); validationErr != nil {
				zap.L().Error("Required fields not present"+validationErr.Error())
				c.JSON(http.StatusBadRequest,  dto.ResponseDTO{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
				return
		}

		//TODO: Validate if customer is present
		//TODO: Add amount to customer DB as well
		newTransaction := models.Transaction{
			TransactionId: uuid.New().String(),
			Amount: transaction.Amount,
			Description: transaction.Description,
			CustomerID: customerId,
		}
		
		data, err := dynamodbattribute.MarshalMap(newTransaction)
		if err != nil {
			zap.L().Error("Marshalling of transaction failed - " + err.Error())
			c.JSON(http.StatusBadRequest,  dto.ResponseDTO{
				Status: http.StatusBadRequest, 
				Message: "Marshalling of transaction failed", 
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		query := &dynamodb.PutItemInput{
			Item:      data,
			TableName: aws.String(transactionCollection),
		}

		result, err := configs.DB.PutItem(query)
		if err != nil {
			zap.L().Error("Failed to add transaction - " + err.Error())
			c.JSON(http.StatusBadRequest,  dto.ResponseDTO{
				Status: http.StatusBadRequest, 
				Message: "Failed to add transaction", 
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		dataInBytes,_ := json.Marshal(result)
		zap.L().Info("Successfully added transaction"+string(dataInBytes))
		c.JSON(http.StatusCreated, dto.ResponseDTO{
			Status: http.StatusCreated, 
			Message: "success", 
			Data: map[string]interface{}{"data": result},
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