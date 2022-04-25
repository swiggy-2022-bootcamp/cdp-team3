package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/errors"
	kafkaProducer "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/kafka"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"

	"golang.org/x/crypto/bcrypt"

	"go.uber.org/zap"
)

type CustomerServiceImpl struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerServiceImpl(customerRepository repository.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{customerRepository: customerRepository}
}

func (cs CustomerServiceImpl) AddCustomer(customer *models.Customer) *errors.AppError {
	zap.L().Info("Inside AddCustomer Service")
	customerPassword, errhash := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if errhash != nil {
		return errors.ParseFail("Parse Error")
	}
	customer.Password = string(customerPassword)
	customer.ConfirmPassword = customer.Password
	err := cs.customerRepository.AddCustomerToDB(customer)
	if err != nil {
		zap.L().Error(err.Message)
		return err
	}
	go kafkaProducer.UserCreationProducer(customer.CustomerId)
	return nil
}

func (cs CustomerServiceImpl) GetAllCustomers() ([]models.Customer, *errors.AppError) {
	zap.L().Info("Inside GetAllCustomers Service")
	result, err := cs.customerRepository.GetAllCustomersFromDB()
	if err != nil {
		zap.L().Error(err.Message)
		return nil, err
	}
	return result, nil
}

func (cs CustomerServiceImpl) GetCustomerById(customerId string) (*models.Customer, *errors.AppError) {
	zap.L().Info("Inside GetCustomerById Service")
	result, err := cs.customerRepository.GetCustomerByIdFromDB(customerId)
	if err != nil {
		zap.L().Error(err.Message)
		return nil, err
	}
	return result, nil
}

func (cs CustomerServiceImpl) GetCustomerByEmail(emailId string) (*models.Customer, *errors.AppError) {
	zap.L().Info("Inside GetCustomerByEmail Service")
	result, err := cs.customerRepository.GetCustomerByEmailFromDB(emailId)
	if err != nil {
		zap.L().Error(err.Message)
		return nil, err
	}
	return result, nil
}

func (cs CustomerServiceImpl) DeleteCustomerById(customerId string) (bool, *errors.AppError) {
	zap.L().Info("Inside DeleteCustomerById Service")
	result, err := cs.customerRepository.DeleteCustomerByIdFromDB(customerId)
	if err != nil {
		zap.L().Error(err.Message)
		return result, err
	}
	go kafkaProducer.UserDeletionProducer(customerId)

	return result, nil
}

func (cs CustomerServiceImpl) UpdateCustomerById(customerId string, customer *models.Customer) (*models.Customer, *errors.AppError) {

	customer.CustomerId = customerId

	customerPassword, errhash := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if errhash != nil {
		return nil, errors.ParseFail("Parse Error")
	}
	customer.Password = string(customerPassword)
	customer.ConfirmPassword = customer.Password
	result, err := cs.customerRepository.UpdateCustomerByIdFromDB(customerId, customer)
	if err != nil {
		zap.L().Error(err.Message)
		return nil, err
	}
	return result, nil
}

// func (cs CustomerServiceImpl) UpdateRewardPointsById(customerId string, rewards int32) (bool, *errors.AppError) {
// 	_, err := cs.customerRepository.UpdateRewardPointsByIdFromDB(customerId, rewards)
// 	if err != nil {
// 		zap.L().Error(err.Message)
// 		return false, err
// 	}
// 	return false, nil
// }
