package repository

import "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"

type AdminRepo struct {
}

type AdminRepository interface {
	Create(customer models.Customer) (*models.Customer, error)
}

func (adminRepository *AdminRepo) Create(customer models.Customer) (*models.Customer, error) {
	return nil, nil
}
