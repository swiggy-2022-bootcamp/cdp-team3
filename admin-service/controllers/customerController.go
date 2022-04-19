package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/errors"
	model "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"
	repository "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/repository"
	service "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/services"

	"github.com/gin-gonic/gin"
)

// func CheckAuthorized(usergroup string) {

// }

var customerService service.CustomerServiceInterface

func init() {
	customerService = service.InitCustomerService(&repository.CustomerCollection{})
}

// CreateCustomer godoc
// @Summary creates a customer account
// @Description creates a customer account when the admin is verified
// @Tags Admin
// @Schemes
// @Accept json
// @Produce json
// @Param        Customer body models.Customer  true "customer details"
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /customers [POST]
func CreateCustomer(c *gin.Context) {
	newCustomer := model.Customer{}
	json.NewDecoder(c.Request.Body).Decode(&newCustomer)
	createdCustomer, err := customerService.AddCustomer(newCustomer)

	if err != nil {
		userErr, _ := err.(*errors.CustomerError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *createdCustomer)
}

// GetCustomerByID godoc
// @Summary fetches a customer account by ID
// @Description fetches the details of a customer based on the given ID
// @Tags Admin
// @Schemes
// @Accept json
// @Produce json
// @Param        CustomerID path string  true "customer id"
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /customers/{id} [GET]
func GetCustomerById(c *gin.Context) {
	fetchedCustomer, err := customerService.GetCustomerById(c.Param("customerId"))

	if err != nil {
		userErr, _ := err.(*errors.CustomerError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *fetchedCustomer)
}

// GetCustomerByEmail godoc
// @Summary fetches a customer account by email
// @Description fetches the details of a customer based on the given email
// @Tags Admin
// @Schemes
// @Accept json
// @Produce json
// @Param        CustomerEmail path string  true "customer email"
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /customers/email/{email} [GET]
func GetCustomerByEmail(c *gin.Context) {
	fetchedCustomer, err := customerService.GetCustomerByEmail(c.Param("customerEmail"))

	if err != nil {
		userErr, _ := err.(*errors.CustomerError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *fetchedCustomer)
}

// UpdateCustomer godoc
// @Summary Updates a customer account
// @Description Updates The Customer Details
// @Tags Admin
// @Schemes
// @Accept json
// @Produce json
// @Param        CustomerID path string  true "customer id"
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /customers/{id} [PUT]
func UpdateCustomer(c *gin.Context) {
	customer := model.Customer{}
	json.NewDecoder(c.Request.Body).Decode(&customer)
	updatedCustomer, err := customerService.UpdateCustomer(c.Param("customerId"), customer)

	if err != nil {
		userErr, _ := err.(*errors.CustomerError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *updatedCustomer)
}

// DeleteCustomer godoc
// @Summary deletes a customer account
// @Description deletes The Customer Details based on the given ID
// @Tags Admin
// @Schemes
// @Accept json
// @Produce json
// @Param        Customer 	body 	models.Customer	true  	"customer details"
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /customers/{id} [DELETE]
func DeleteCustomer(c *gin.Context) {
	successMessage, err := customerService.DeleteCustomer(c.Param("customerId"))

	if err != nil {
		userErr, _ := err.(*errors.CustomerError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *successMessage)
}

// Healthcheck godoc
// @Summary Checks whether the service is up & running
// @Description When a request is made to the / endpoint, if the service is running, it returns "Okay"
// @Tags Health
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router / [GET]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Okay",
	})
}
