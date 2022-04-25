package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/domain/services"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/dto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/utils"
	"go.uber.org/zap"
)

type AdminController struct {
	adminService services.AdminService
}

func NewAdminController(adminService services.AdminService) AdminController {
	return AdminController{adminService: adminService}
}

func dynamoModelConvAdmin(customer models.Admin) *models.Admin {
	return &models.Admin{

		AdminId:   uuid.New().String(),
		IsAdmin:   true,
		Firstname: customer.Firstname,
		Lastname:  customer.Lastname,
		Username:  customer.Username,
		Password:  customer.Password,
		Email:     customer.Email,
		Telephone: customer.Telephone,
		Status:    "1",
		DateAdded: time.Now(),
	}
}

// AddAdmin godoc
// @Summary Adds Admin User
// @Description We can create an admin user using this function
// @Tags Admin Service - Admin Operations
// @Schemes
// @Accept json
// @Produce json
// @Param        Admin Details body models.SwaggerAdmin true "admin details"
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /admin [POST]
func (ac AdminController) AddAdmin(c *gin.Context) {
	zap.L().Info("Inside AddAdmin Controller")

	var admin models.Admin

	if err := c.BindJSON(&admin); err != nil {
		c.Error(err)
		err_ := errors.NewBadRequestError(err.Error())
		zap.L().Error(err_.Message)
		c.JSON(err_.Code, gin.H{"message": err_.Message})
		return
	}

	adminRecord := dynamoModelConvAdmin(admin)
	err := ac.adminService.AddAdmin(adminRecord)
	if err != nil {
		zap.L().Error(err.Message)
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	zap.L().Info("Created Admin successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Admin added successfully"})
}

// GetSelf godoc
// @Summary Gets the details of the Admin User
// @Description Gets the details of the Admin User that is currently logged in
// @Tags Admin Service - Admin Operations
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {object} 	models.Admin
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /admin/user [GET]
func (ac AdminController) GetSelf() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetSelf Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		adminId := utils.UserId(c)
		admin, err := ac.adminService.GetSelf(adminId)

		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(err.Code, dto.ResponseDTO{
				Status:  err.Code,
				Message: err.Message,
			})
			return
		}
		zap.L().Info("Fetched Admin Details successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"admin": admin},
		})
	}
}
