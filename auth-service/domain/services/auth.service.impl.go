package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/models"
)

type AuthServiceImpl struct {
	userRepository repository.UserRepository
}

func NewAuthServiceImpl(userRepository repository.UserRepository) AuthService {
	return &AuthServiceImpl{userRepository: userRepository}
}

func (asi AuthServiceImpl) GetUserByEmail(email string) (*models.User, *errors.AppError) {
	result, err := asi.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return result, nil
}
