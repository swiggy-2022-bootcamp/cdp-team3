package service

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"
	repository "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/repository"
)

type AdminCollection struct {
	adminRepository repository.AdminRepository
}

type AdminServices interface {
	AddCustomer(customer models.Customer) (*models.Customer, error)
}

func (customerService *AdminCollection) AddCustomer(customer models.Customer) (*models.Customer, error) {
	createdCustomer, err := customerService.adminRepository.Create(customer)
	if err != nil {
		return nil, err
	}
	return createdCustomer, nil
}
