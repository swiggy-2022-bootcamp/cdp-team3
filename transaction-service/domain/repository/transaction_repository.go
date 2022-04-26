package repository

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/models"
)

type TransactionRepository interface {
	GetTransactionByCustomerIdInDB(customerId string) ([]models.Transaction, *errors.AppError)
	AddTransactionAmtToCustomerInDB(transaction *models.Transaction) (*models.Transaction, *errors.AppError)
}