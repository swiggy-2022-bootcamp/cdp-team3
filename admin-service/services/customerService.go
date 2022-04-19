package service

import (
	"time"

	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/errors"
	model "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"
	repository "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/repository"
)

type CustomerServiceInterface interface {
	AddCustomer(customer model.Customer) (*model.Customer, error)
	GetCustomerById(customerId string) (*model.Customer, error)
	GetCustomerByEmail(customerEmail string) (*model.Customer, error)
	UpdateCustomer(customerId string, customer model.Customer) (*model.Customer, error)
	DeleteCustomer(customerId string) (*string, error)
}

type CustomerService struct {
	customerCollection repository.CustomerCollectionInterface
}

func InitCustomerService(repositoryToInject repository.CustomerCollectionInterface) CustomerServiceInterface {
	customerService := new(CustomerService)
	customerService.customerCollection = repositoryToInject
	return customerService
}

func (customerService *CustomerService) AddCustomer(customer model.Customer) (*model.Customer, error) {
	fetchedCustomer, _ := customerService.GetCustomerByEmail(customer.Email)
	if fetchedCustomer != nil {
		return nil, errors.NewEmailAlreadyRegisteredError()
	}

	customer.DateAdded = time.Now()
	createdCustomer, err := customerService.customerCollection.Create(customer)
	if err != nil {
		return nil, err
	}
	return createdCustomer, nil
}

func (customerService *CustomerService) GetCustomerById(customerId string) (*model.Customer, error) {
	fetchedCustomer, err := customerService.customerCollection.GetById(customerId)
	if err != nil {
		return nil, err
	}
	return fetchedCustomer, nil
}

func (customerService *CustomerService) GetCustomerByEmail(customerEmail string) (*model.Customer, error) {
	fetchedCustomer, err := customerService.customerCollection.GetByEmail(customerEmail)
	if err != nil {
		return nil, err
	}
	return fetchedCustomer, nil
}

func (customerService *CustomerService) UpdateCustomer(customerId string, customer model.Customer) (*model.Customer, error) {
	customer.CustomerId = customerId
	updatedCustomer, err := customerService.customerCollection.Update(customer)
	if err != nil {
		return nil, err
	}
	return updatedCustomer, nil
}

func (customerService *CustomerService) DeleteCustomer(customerId string) (*string, error) {
	successMessage, err := customerService.customerCollection.Delete(customerId)
	if err != nil {
		return nil, err
	}
	return successMessage, nil
}
