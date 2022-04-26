// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/swiggy-ipp/cart-service/repositories (interfaces: CartRepository)

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/swiggy-ipp/cart-service/models"
)

// MockCartRepository is a mock of CartRepository interface.
type MockCartRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCartRepositoryMockRecorder
}

// MockCartRepositoryMockRecorder is the mock recorder for MockCartRepository.
type MockCartRepositoryMockRecorder struct {
	mock *MockCartRepository
}

// NewMockCartRepository creates a new mock instance.
func NewMockCartRepository(ctrl *gomock.Controller) *MockCartRepository {
	mock := &MockCartRepository{ctrl: ctrl}
	mock.recorder = &MockCartRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartRepository) EXPECT() *MockCartRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCartRepository) Create(arg0 context.Context, arg1 *models.Cart) (*models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCartRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCartRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockCartRepository) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCartRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCartRepository)(nil).Delete), arg0, arg1)
}

// Read mocks base method.
func (m *MockCartRepository) Read(arg0 context.Context, arg1 string) (*models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1)
	ret0, _ := ret[0].(*models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockCartRepositoryMockRecorder) Read(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockCartRepository)(nil).Read), arg0, arg1)
}

// ReadByUserID mocks base method.
func (m *MockCartRepository) ReadByUserID(arg0 context.Context, arg1 string) (*models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadByUserID", arg0, arg1)
	ret0, _ := ret[0].(*models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByUserID indicates an expected call of ReadByUserID.
func (mr *MockCartRepositoryMockRecorder) ReadByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadByUserID", reflect.TypeOf((*MockCartRepository)(nil).ReadByUserID), arg0, arg1)
}

// UpdateCartItems mocks base method.
func (m *MockCartRepository) UpdateCartItems(arg0 context.Context, arg1 *models.Cart) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCartItems", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCartItems indicates an expected call of UpdateCartItems.
func (mr *MockCartRepositoryMockRecorder) UpdateCartItems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCartItems", reflect.TypeOf((*MockCartRepository)(nil).UpdateCartItems), arg0, arg1)
}