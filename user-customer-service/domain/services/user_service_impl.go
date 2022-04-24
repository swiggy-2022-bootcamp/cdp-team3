package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/models"

	"go.uber.org/zap"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func (rs UserServiceImpl) AddUser(user *models.User) *errors.AppError {
	zap.L().Info("Inside AddUser Service")

	err := rs.userRepository.AddUserToDB(user)
	if err != nil {
		zap.L().Error(err.Message)
		return err
	}
	return nil
}

func (rs UserServiceImpl) GetAllUsers() ([]models.User, *errors.AppError) {
	zap.L().Info("Inside GetAllUsers Service")
	result, err := rs.userRepository.GetAllUsersFromDB()
	if err != nil {
		zap.L().Error(err.Message)
		return nil, err
	}
	return result, nil
}

func (rs UserServiceImpl) GetUserById(userId string) (*models.User, *errors.AppError) {
	zap.L().Info("Inside GetUserById Service")
	result, err := rs.userRepository.GetUserByIdFromDB(userId)
	if err != nil {
		zap.L().Error(err.Message)
		return nil, err
	}
	return result, nil
}

func (rs UserServiceImpl) GetUsersByCustomerId(customerId string) ([]models.User, *errors.AppError) {
	zap.L().Info("Inside GetUsersByCustomerId Service")
	result, err := rs.userRepository.GetUsersByCustomerIdFromDB(customerId)
	if err != nil {
		zap.L().Error(err.Message)
		return nil, err
	}
	return result, nil
}
