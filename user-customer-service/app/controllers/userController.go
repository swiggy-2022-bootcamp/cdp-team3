package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	repository "github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/domain/repository"
	service "github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/domain/services"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/middlewares"
	model "github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/models"
)

var customerService service.UserServiceInterface

func init() {
	customerService = service.InitUserService(&repository.UserCollection{})
}

// CreateUser godoc
// @Summary creates a customer account
// @Description creates a customer account
// @Tags User
// @Schemes
// @Accept json
// @Produce json
// @Param        User body models.User  true "user details"
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /customers [POST]
func CreateUser(c *gin.Context) {
	newUser := model.User{}
	json.NewDecoder(c.Request.Body).Decode(&newUser)
	createdUser, err := customerService.AddUser(newUser)

	if err != nil {
		userErr, _ := err.(*errors.UserError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *createdUser)
}

// GetUserByID godoc
// @Summary fetches a customer account by ID
// @Description fetches the details of a customer based on the given ID
// @Tags User
// @Schemes
// @Accept json
// @Produce json
// @Param        UserID path string  true "user id"
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /customers/{id} [GET]
func GetUserById(c *gin.Context) {
	// fetchedUser, err := customerService.GetUserById(c.Param("id"))

	// if err != nil {
	// 	userErr, _ := err.(*errors.UserError)
	// 	c.JSON(userErr.Status, userErr.ErrorMessage)
	// 	return
	// }

	// c.JSON(200, *fetchedUser)
	userClaim := middlewares.AuthenticateJWT()
	c.JSON(200, c.Bind(userClaim))
}

// UpdateUser godoc
// @Summary Updates a customer account
// @Description Updates the customer details
// @Tags User
// @Schemes
// @Accept json
// @Produce json
// @Param        UserID path string  true "customer id"
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /customers/{id} [PUT]
func UpdateUser(c *gin.Context) {
	customer := model.User{}
	json.NewDecoder(c.Request.Body).Decode(&customer)
	updatedUser, err := customerService.UpdateUser(c.Param("id"), customer)

	if err != nil {
		userErr, _ := err.(*errors.UserError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *updatedUser)
}

// DeleteUser godoc
// @Summary Deletes a customer account
// @Description Deletes the User details based on the given ID
// @Tags User
// @Schemes
// @Accept json
// @Produce json
// @Param        User 	body 	models.User	true  	"customer details"
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /customers/{id} [DELETE]
func DeleteUser(c *gin.Context) {
	successMessage, err := customerService.DeleteUser(c.Param("id"))

	if err != nil {
		userErr, _ := err.(*errors.UserError)
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
