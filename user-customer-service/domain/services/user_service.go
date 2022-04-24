package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/models"
)

type UserService interface {
	AddUser(reward *models.User) *errors.AppError
	GetAllUsers() ([]models.User, *errors.AppError)
	GetUserById(rewardId string) (*models.User, *errors.AppError)
	GetUsersByCustomerId(customerId string) ([]models.User, *errors.AppError)
}
