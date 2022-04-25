package repository

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, *errors.AppError)
}
