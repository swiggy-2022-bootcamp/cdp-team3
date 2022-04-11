// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/domain/services (interfaces: TransactionService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	errors "github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/errors"
	models "github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/models"
)

// MockTransactionService is a mock of TransactionService interface.
type MockTransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionServiceMockRecorder
}

// MockTransactionServiceMockRecorder is the mock recorder for MockTransactionService.
type MockTransactionServiceMockRecorder struct {
	mock *MockTransactionService
}

// NewMockTransactionService creates a new mock instance.
func NewMockTransactionService(ctrl *gomock.Controller) *MockTransactionService {
	mock := &MockTransactionService{ctrl: ctrl}
	mock.recorder = &MockTransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionService) EXPECT() *MockTransactionServiceMockRecorder {
	return m.recorder
}

// AddTransactionAmtToCustomer mocks base method.
func (m *MockTransactionService) AddTransactionAmtToCustomer(arg0 *models.Transaction) (*models.Transaction, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTransactionAmtToCustomer", arg0)
	ret0, _ := ret[0].(*models.Transaction)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// AddTransactionAmtToCustomer indicates an expected call of AddTransactionAmtToCustomer.
func (mr *MockTransactionServiceMockRecorder) AddTransactionAmtToCustomer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransactionAmtToCustomer", reflect.TypeOf((*MockTransactionService)(nil).AddTransactionAmtToCustomer), arg0)
}

// GetTransactionByCustomerId mocks base method.
func (m *MockTransactionService) GetTransactionByCustomerId(arg0 string) ([]models.Transaction, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByCustomerId", arg0)
	ret0, _ := ret[0].([]models.Transaction)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// GetTransactionByCustomerId indicates an expected call of GetTransactionByCustomerId.
func (mr *MockTransactionServiceMockRecorder) GetTransactionByCustomerId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByCustomerId", reflect.TypeOf((*MockTransactionService)(nil).GetTransactionByCustomerId), arg0)
}