package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/dto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/errors"
)

type PaymentService interface {
	Pay(pr dto.PaymentRequestDto) (string, *errors.AppError)
}
