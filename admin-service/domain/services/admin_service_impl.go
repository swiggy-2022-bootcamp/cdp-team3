package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"
	"golang.org/x/crypto/bcrypt"

	"go.uber.org/zap"
)

type AdminServiceImpl struct {
	adminRepository repository.AdminRepository
}

func NewAdminServiceImpl(adminRepository repository.AdminRepository) AdminService {
	return &AdminServiceImpl{adminRepository: adminRepository}
}

func (as AdminServiceImpl) AddAdmin(admin *models.Admin) *errors.AppError {
	zap.L().Info("Inside AddAdmin Service")
	adminPassword, errhash := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if errhash != nil {
		return errors.ParseFail("Parse Error")
	}
	admin.Password = string(adminPassword)
	err := as.adminRepository.AddAdminToDB(admin)
	if err != nil {
		zap.L().Error(err.Message)
		return err
	}
	return nil
}

func (as AdminServiceImpl) GetSelf(adminId string) (*models.Admin, *errors.AppError) {
	zap.L().Info("Inside GetSelf Service")
	result, err := as.adminRepository.GetSelfFromDB(adminId)
	if err != nil {
		zap.L().Error(err.Message)
		return nil, err
	}
	return result, nil
}
