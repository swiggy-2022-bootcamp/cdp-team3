package service

import (
	"time"

	repository "github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/errors"
	model "github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/models"
)

type UserServiceInterface interface {
	AddUser(user model.User) (*model.User, error)
	GetUserById(userId string) (*model.User, error)
	GetUserByEmail(userEmail string) (*model.User, error)
	UpdateUser(userId string, user model.User) (*model.User, error)
	DeleteUser(userId string) (*string, error)
}

type UserService struct {
	userCollection repository.UserCollectionInterface
}

func InitUserService(repositoryToInject repository.UserCollectionInterface) UserServiceInterface {
	userService := new(UserService)
	userService.userCollection = repositoryToInject
	return userService
}

func (userService *UserService) AddUser(user model.User) (*model.User, error) {
	fetchedUser, _ := userService.GetUserByEmail(user.Email)
	if fetchedUser != nil {
		return nil, errors.NewEmailAlreadyRegisteredError()
	}

	user.DateAdded = time.Now()
	createdUser, err := userService.userCollection.Create(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (userService *UserService) GetUserById(userId string) (*model.User, error) {
	fetchedUser, err := userService.userCollection.GetById(userId)
	if err != nil {
		return nil, err
	}
	return fetchedUser, nil
}

func (userService *UserService) GetUserByEmail(userEmail string) (*model.User, error) {
	fetchedUser, err := userService.userCollection.GetByEmail(userEmail)
	if err != nil {
		return nil, err
	}
	return fetchedUser, nil
}

func (userService *UserService) UpdateUser(userId string, user model.User) (*model.User, error) {
	user.UserId = userId
	updatedUser, err := userService.userCollection.Update(user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (userService *UserService) DeleteUser(userId string) (*string, error) {
	successMessage, err := userService.userCollection.Delete(userId)
	if err != nil {
		return nil, err
	}
	return successMessage, nil
}
