package repository

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"
)

type CustomerRepository interface {
	AddCustomerToDB(customer *models.Customer) *errors.AppError
	GetAllCustomersFromDB() ([]models.Customer, *errors.AppError)
	GetCustomerByIdFromDB(customerId string) (*models.Customer, *errors.AppError)
	GetCustomerByEmailFromDB(emailId string) (*models.Customer, *errors.AppError)
	DeleteCustomerByIdFromDB(customerId string) (bool, *errors.AppError)
	UpdateCustomerByIdFromDB(customerId string, customer *models.Customer) (*models.Customer, *errors.AppError)
}
