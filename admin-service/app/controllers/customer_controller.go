package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/domain/services"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/dto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"

	"go.uber.org/zap"
)

type CustomerController struct {
	customerService services.CustomerService
}

func NewCustomerController(customerService services.CustomerService) CustomerController {
	return CustomerController{customerService: customerService}
}

func dynamoModelConv(customer models.Customer) *models.Customer {
	return &models.Customer{

		CustomerId:        uuid.New().String(),
		IsAdmin:           false,
		Firstname:         customer.Firstname,
		Lastname:          customer.Lastname,
		Username:          customer.Username,
		Password:          customer.Password,
		ConfirmPassword:   customer.ConfirmPassword,
		Email:             customer.Email,
		Telephone:         customer.Telephone,
		Address:           customer.Address,
		Status:            "1",
		Approved:          "1",
		DateAdded:         time.Now(),
		Rewards:           0,
		TransactionPoints: 0,
	}
}

// CreateCustomer godoc
// @Summary creates a customer account
// @Description creates a customer account when the admin is verified
// @Tags Admin Service - Customer Operations
// @Schemes
// @Accept json
// @Produce json
// @Param        Customer body models.SwaggerCustomer  true "customer details"
// @Success	200  {String} 	success
// @Failure	400  {number} 	409
// @Failure	500  {number} 	500
// @Security Bearer Token
// @Router /customers [POST]
func (cc CustomerController) AddCustomer(c *gin.Context) {
	zap.L().Info("Inside AddCustomer Controller")

	var customer models.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.Error(err)
		err_ := errors.NewBadRequestError(err.Error())
		zap.L().Error(err_.Message)
		c.JSON(err_.Code, gin.H{"message": err_.Message})
		return
	}

	customerRecord := dynamoModelConv(customer)
	if customerRecord.Password != customer.ConfirmPassword {
		c.JSON(409, gin.H{"message": "Password & Confirm Password Need To Be The Same"})
		return
	}
	err := cc.customerService.AddCustomer(customerRecord)
	if err != nil {
		zap.L().Error(err.Message)
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	zap.L().Info("Created Customer successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Customer added successfully"})
}

// GetAllCustomers godoc
// @Summary Fetch all the customers
// @Description This request will fetch all the customers
// @Tags Admin Service - Customer Operations
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {array} 	models.Customer
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /customers [GET]
func (cc CustomerController) GetAllCustomers() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetAllCustomer Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		customersList, err := cc.customerService.GetAllCustomers()

		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(err.Code, dto.ResponseDTO{
				Status:  err.Code,
				Message: err.Message,
			})
			return
		}

		zap.L().Info("Fetched all customers successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"customers": customersList},
		})
	}
}

// GetCustomerById godoc
// @Summary Get Customer based on customer ID.
// @Description Gets the customer details based on customer ID.
// @Tags Admin Service - Customer Operations
// @Schemes
// @Param customerId path string true "Customer Id"
// @Produce json
// @Success	200  {object} models.Customer
// @Failure	500  {number} http.StatusInternalServerError
// @Security Bearer Token
// @Router /customers/{customerId} [GET]
func (cc CustomerController) GetCustomerById() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetCustomerById Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		customerId := c.Param("customerId")
		customer, err := cc.customerService.GetCustomerById(customerId)

		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(err.Code, dto.ResponseDTO{
				Status:  err.Code,
				Message: err.Message,
			})
			return
		}
		zap.L().Info("Fetched customer successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"customer": customer},
		})
	}
}

// GetCustomerByEmail godoc
// @Summary fetches a customer account by email
// @Description fetches the details of a customer based on the given email
// @Tags Admin Service - Customer Operations
// @Schemes
// @Accept json
// @Produce json
// @Param        emailId path string  true "customer email"
// @Success	200  {object} 	models.Customer
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Security Bearer Token
// @Router /customers/email/{emailId} [GET]
func (cc CustomerController) GetCustomerByEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetCustomerByEmail Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		emailId := c.Param("emailId")
		customer, err := cc.customerService.GetCustomerByEmail(emailId)

		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(err.Code, dto.ResponseDTO{
				Status:  err.Code,
				Message: err.Message,
			})
			return
		}
		zap.L().Info("Fetched customer successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"customer": customer},
		})
	}
}

// DeleteCustomer godoc
// @Summary deletes a customer account
// @Description deletes The Customer Details based on the given ID
// @Tags Admin Service - Customer Operations
// @Schemes
// @Accept json
// @Produce json
// @Param        customerId path string  true "Customer Id"
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Security Bearer Token
// @Router /customers/{customerId} [DELETE]
func (cc CustomerController) DeleteCustomerById() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside DeleteCustomer Controller")
		customerId := c.Param("customerId")
		_, err := cc.customerService.DeleteCustomerById(customerId)
		if err != nil {
			zap.L().Error(err.Message)
			c.Error(err.Error())
			c.JSON(err.Code, gin.H{"message": err.Message})
			return
		}
		zap.L().Info("Created Deleted successfully")
		c.JSON(http.StatusOK, gin.H{"message": "Customer Deleted successfully"})
	}
}

// UpdateCustomer godoc
// @Summary Updates a customer account
// @Description Updates The Customer Details using given id
// @Tags Admin Service - Customer Operations
// @Schemes
// @Accept json
// @Produce json
// @Param        CustomerID path string  true "customer id"
// @Param        Customer Details body models.Customer  true "customer id"
// @Success	200  {object} 	models.Customer
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /customers/{customerId} [PUT]
func (cc CustomerController) UpdateCustomerById() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside UpdateStatusById Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		customerId := c.Param("customerId")

		var customer *models.Customer
		if err := c.BindJSON(&customer); err != nil {
			zap.L().Error("Invalid Request")
			c.JSON(http.StatusBadRequest, dto.ResponseDTO{
				Status:  http.StatusBadRequest,
				Message: "Invalid Request",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}
		if customer.Password != customer.ConfirmPassword {
			c.JSON(409, gin.H{"message": "Password & Confirm Password Need To Be The Same"})
			return
		}
		updatedCustomer, err := cc.customerService.UpdateCustomerById(customerId, customer)
		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(err.Code, dto.ResponseDTO{
				Status:  err.Code,
				Message: err.Message,
			})
			return
		}

		zap.L().Info("Updated Customer details successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"customer": updatedCustomer},
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
