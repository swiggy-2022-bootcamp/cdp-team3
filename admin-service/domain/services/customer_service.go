package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"
)

type CustomerService interface {
	AddCustomer(customer *models.Customer) *errors.AppError
	GetAllCustomers() ([]models.Customer, *errors.AppError)
	GetCustomerById(customerId string) (*models.Customer, *errors.AppError)
	GetCustomerByEmail(emailId string) (*models.Customer, *errors.AppError)
	DeleteCustomerById(customerId string) (bool, *errors.AppError)
	UpdateCustomerById(customerId string, customer *models.Customer) (*models.Customer, *errors.AppError)
	// UpdateRewardPointsById(customerId string, rewards int32) (bool, *errors.AppError)
}
