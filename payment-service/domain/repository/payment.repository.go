package repository

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/models"
)

type PaymentRepository interface {
	CreatePayment(models.Payment) (string, *errors.AppError)
}
