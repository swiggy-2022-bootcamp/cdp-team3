package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/models"
)

type TransactionService interface {
	GetTransactionByCustomerId(customerId string) ([]models.Transaction, *errors.AppError)
	AddTransactionAmtToCustomer(transaction *models.Transaction) (*models.Transaction, *errors.AppError)
}