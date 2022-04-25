package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"
)

type AdminService interface {
	AddAdmin(admin *models.Admin) *errors.AppError
	GetSelf(adminId string) (*models.Admin, *errors.AppError)
}
