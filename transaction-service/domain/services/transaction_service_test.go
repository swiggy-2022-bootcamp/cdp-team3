package services

import (
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/models"
)

func TestOrderServiceImpl_GetTransactionByCustomerId(t *testing.T) {
  gin.SetMode(gin.TestMode)
  
  customerId := "5243c6b-8a0c-4f99-8b2b-6eeef7605a37"
  successTransactions := []models.Transaction{
    {
      TransactionId : "423fec6b-82c-4f99-8b2b-6eeef7605a37",
      Amount: 200.1,
      Description: "Trnsaction Points",
      CustomerID: "5243c6b-8a0c-4f99-8b2b-6eeef7605a37",
    },
    {
      TransactionId : "423dfvfv43b-82c-4f99-8b2b-6eeef7605a37",
      Amount: 250.1,
      Description: "Adding Points for Purchase of mobile",
      CustomerID: "5243c6b-8a0c-4f99-8b2b-6eeef7605a37",
    },
  }

  testCases := []struct {
    name          string
    buildStubs    func(repository *mocks.MockTransactionRepository)
    checkResponse func(t *testing.T,res []models.Transaction, err interface{})
  }{
    {
      name: "SuccessTransactionsFound",
      buildStubs: func(repository *mocks.MockTransactionRepository) {
        repository.EXPECT().
          GetTransactionByCustomerIdInDB(customerId).
          Times(1).
          Return(successTransactions, nil)
      },
      checkResponse: func(t *testing.T,res []models.Transaction, err interface{}) {
        assert.NotNil(t,res)
        assert.Nil(t,err)
      },
    },
    {
      name: "FailureTransactionNotFound",
      buildStubs: func(repository *mocks.MockTransactionRepository) {
        repository.EXPECT().
          GetTransactionByCustomerIdInDB(customerId).
          Times(1).
          Return(nil, errors.NewNotFoundError(""))
      },
      checkResponse: func(t *testing.T,res []models.Transaction, err interface{}) {
        assert.Nil(t,res)
        assert.NotNil(t,err)
      },
    },
    {
      name: "FailureUnexpectedError",
      buildStubs: func(repository *mocks.MockTransactionRepository) {
        repository.EXPECT().
          GetTransactionByCustomerIdInDB(customerId).
          Times(1).
          Return(nil, errors.NewUnexpectedError(""))
      },
      checkResponse: func(t *testing.T,res []models.Transaction, err interface{}) {
        assert.Nil(t,res)
        assert.NotNil(t,err)
      },
    },
  }

  for i := range testCases {
    tc := testCases[i]
    t.Run(tc.name, func(t *testing.T) {
      ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			transactionRepository := mocks.NewMockTransactionRepository(ctrl)
			tc.buildStubs(transactionRepository)

			transactionServiceImpl := NewTransactionServiceImpl(transactionRepository)
			res, err := transactionServiceImpl.GetTransactionByCustomerId(customerId)
			tc.checkResponse(t,  res, err)
    })
  }
}


// func TestTransactionServiceImpl_AddTransactionAmtToCustomer(t *testing.T) {
//   gin.SetMode(gin.TestMode)
	

// 	transaction := &models.Transaction{
// 		TransactionId : "423fec6b-82c-4f99-8b2b-6eeef7605a37",
// 		Amount: 200.1,
// 		Description: "Trnsaction Points",
// 		CustomerID: "5243c6b-8a0c-4f99-8b2b-6eeef7605a37",
// 	}
// 	testCases := []struct {
// 		name          string
// 		buildStubs    func(repository *mocks.MockTransactionRepository)
// 		checkResponse func(t *testing.T, transaction *models.Transaction, err interface{})
// 	}{
		
// 		{
// 			name: "SuccessAddTransaction",
// 			buildStubs: func(repository *mocks.MockTransactionRepository) {
// 				repository.EXPECT().
// 				AddTransactionAmtToCustomerInDB(transaction).
// 					Times(1).
// 					Return(transaction, nil)
// 			},
// 			checkResponse: func(t *testing.T, trans *models.Transaction,err interface{}) {
// 				assert.NotNil(t, trans)
// 				assert.Equal(t, trans, transaction)
// 				assert.Nil(t,err)
// 			},
// 		},
// 		{
// 			name: "FailureAddTransaction",
// 			buildStubs: func(repository *mocks.MockTransactionRepository) {
// 				repository.EXPECT().
// 				AddTransactionAmtToCustomerInDB(transaction).
// 					Times(1).
// 					Return(nil, errors.NewUnexpectedError(""))

// 			},
// 			checkResponse: func(t *testing.T, trans *models.Transaction, err interface{}) {
// 				assert.Nil(t,trans)
// 				assert.NotNil(t,err)
// 			},
// 		},
// 		{
// 			name: "FailureUnexpectedError",
// 			buildStubs: func(repository *mocks.MockTransactionRepository) {
// 				repository.EXPECT().
// 				AddTransactionAmtToCustomerInDB(transaction).
// 					Times(1).
// 					Return(nil,errors.NewUnexpectedError(""))
// 			},
// 			checkResponse: func(t *testing.T, trans *models.Transaction, err interface{}) {
// 				assert.Nil(t,trans)
// 				assert.NotNil(t,err)
// 			},
// 		},
		
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			transactionRepository := mocks.NewMockTransactionRepository(ctrl)
// 			tc.buildStubs(transactionRepository)

// 			transactionServiceImpl := NewTransactionServiceImpl(transactionRepository)
// 			res, err := transactionServiceImpl.AddTransactionAmtToCustomer(transaction)
// 			tc.checkResponse(t, res,err)
// 		})
// 	}
// }
