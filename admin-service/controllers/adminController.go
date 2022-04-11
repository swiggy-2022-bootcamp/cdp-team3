package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func CheckAuthorized(usergroup string) {

// }

// GetAdminDetails godoc
// @Summary Gets the details of the admin that is logged in
// @Description When a request is made, it returns the admin details
// @Tags Admin
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /user [GET]
func GetSelf(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Successfully Fetched Details"})
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
	c.JSON(200, gin.H{"message": "Successfully Created Customer"})
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
	c.JSON(200, gin.H{"message": "Successfully Fetched Details"})
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
	c.JSON(200, gin.H{"message": "Successfully Fetched Details"})
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
	c.JSON(200, gin.H{"message": "Successfully Updated Details"})
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
	c.JSON(200, gin.H{"message": "Successfully Deleted Customer"})
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
