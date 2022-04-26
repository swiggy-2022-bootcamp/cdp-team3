package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/grpc/admin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/models"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/utils"
	"go.uber.org/zap"
)

var validate = validator.New()

type TransactionServiceImpl struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionServiceImpl(transactionRepository repository.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{transactionRepository: transactionRepository}
}

func (ts TransactionServiceImpl) GetTransactionByCustomerId(customerId string) ([]models.Transaction, *errors.AppError) {
	zap.L().Info("Inside GetTransactionByCustomerId Service")

	transactionList, err := ts.transactionRepository.GetTransactionByCustomerIdInDB(customerId)
	if err != nil {
		return nil, err
	}
	return transactionList, nil
}

func (ts TransactionServiceImpl) AddTransactionAmtToCustomer(transaction *models.Transaction) (*models.Transaction, *errors.AppError) {
	zap.L().Info("Inside AddTransactionAmtToCustomer Service")

	//use the validator library to validate required fields
	if validationErr := validate.Struct(transaction); validationErr != nil {
		zap.L().Error("Required fields not present" + validationErr.Error())
		return nil, errors.NewBadRequestError("Required fields not present" + validationErr.Error())
	}

	transaction = &models.Transaction{
		TransactionId: uuid.New().String(),
		Amount:        transaction.Amount,
		Description:   transaction.Description,
		CustomerID:    transaction.CustomerID,
	}

	transactionAmountAdmin := utils.ProtoConv(transaction)
	grpcResponse, err := admin.SendTransactionAmount(transactionAmountAdmin)

	if grpcResponse.IsAdded != "Success" || err != nil {
		zap.L().Error("Error Updating Transaction Amount for the Customer through Admin GRPC")
		return nil, errors.NewUnexpectedError("Error Updating Transaction Amount")
	}

	newtransaction, err_new := ts.transactionRepository.AddTransactionAmtToCustomerInDB(transaction)
	if err_new != nil {
		return nil, err_new
	}

	return newtransaction, nil
}
