package repository

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"
)

type AdminRepository interface {
	AddAdminToDB(admin *models.Admin) *errors.AppError
	GetSelfFromDB(adminId string) (*models.Admin, *errors.AppError)
}
