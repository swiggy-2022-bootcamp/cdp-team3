package repository

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/models"
)

type UserRepository interface {
	AddUserToDB(reward *models.User) *errors.AppError
	GetAllUsersFromDB() ([]models.User, *errors.AppError)
	GetUserByIdFromDB(rewardId string) (*models.User, *errors.AppError)
	GetUsersByCustomerIdFromDB(customerId string) ([]models.User, *errors.AppError)
}
