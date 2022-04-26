package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/dto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/models"
)

type PaymentServiceImpl struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentServiceImpl(pr repository.PaymentRepository) PaymentService {
	return &PaymentServiceImpl{paymentRepository: pr}
}

func (ps *PaymentServiceImpl) Pay(pr dto.PaymentRequestDto) (string, *errors.AppError) {
	var payment models.Payment
	payment.OrderId = pr.OrderId
	payment.Details = pr.Details
	payment.UserId = pr.UserId
	payment.Amount = pr.Amount
	payment.Status = "Success"
	payment.TransactionId = uuid.New().String()
	payment.CreatedAt = time.Now()

	res, err := ps.paymentRepository.CreatePayment(payment)
	if err != nil {
		return "", err
	}
	return res, nil
}
