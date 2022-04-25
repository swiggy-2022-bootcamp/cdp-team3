// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cdp-team3/categories-service/domain/repository (interfaces: CategoryRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	app_erros "github.com/cdp-team3/categories-service/app-errors"
	models "github.com/cdp-team3/categories-service/domain/models"
	gomock "github.com/golang/mock/gomock"
)

// MockCategoryRepository is a mock of CategoryRepository interface.
type MockCategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryRepositoryMockRecorder
}

// MockCategoryRepositoryMockRecorder is the mock recorder for MockCategoryRepository.
type MockCategoryRepositoryMockRecorder struct {
	mock *MockCategoryRepository
}

// NewMockCategoryRepository creates a new mock instance.
func NewMockCategoryRepository(ctrl *gomock.Controller) *MockCategoryRepository {
	mock := &MockCategoryRepository{ctrl: ctrl}
	mock.recorder = &MockCategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryRepository) EXPECT() *MockCategoryRepositoryMockRecorder {
	return m.recorder
}

// AddCategoryToDB mocks base method.
func (m *MockCategoryRepository) AddCategoryToDB(arg0 *models.Category) *app_erros.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCategoryToDB", arg0)
	ret0, _ := ret[0].(*app_erros.AppError)
	return ret0
}

// AddCategoryToDB indicates an expected call of AddCategoryToDB.
func (mr *MockCategoryRepositoryMockRecorder) AddCategoryToDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCategoryToDB", reflect.TypeOf((*MockCategoryRepository)(nil).AddCategoryToDB), arg0)
}

// DBHealthCheck mocks base method.
func (m *MockCategoryRepository) DBHealthCheck() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBHealthCheck")
	ret0, _ := ret[0].(bool)
	return ret0
}

// DBHealthCheck indicates an expected call of DBHealthCheck.
func (mr *MockCategoryRepositoryMockRecorder) DBHealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBHealthCheck", reflect.TypeOf((*MockCategoryRepository)(nil).DBHealthCheck))
}

// DeleteCategoriesFromDB mocks base method.
func (m *MockCategoryRepository) DeleteCategoriesFromDB(arg0 []string) (bool, *app_erros.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategoriesFromDB", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*app_erros.AppError)
	return ret0, ret1
}

// DeleteCategoriesFromDB indicates an expected call of DeleteCategoriesFromDB.
func (mr *MockCategoryRepositoryMockRecorder) DeleteCategoriesFromDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategoriesFromDB", reflect.TypeOf((*MockCategoryRepository)(nil).DeleteCategoriesFromDB), arg0)
}

// DeleteCategoryByIDFromDB mocks base method.
func (m *MockCategoryRepository) DeleteCategoryByIDFromDB(arg0 string) (bool, *app_erros.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategoryByIDFromDB", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*app_erros.AppError)
	return ret0, ret1
}

// DeleteCategoryByIDFromDB indicates an expected call of DeleteCategoryByIDFromDB.
func (mr *MockCategoryRepositoryMockRecorder) DeleteCategoryByIDFromDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategoryByIDFromDB", reflect.TypeOf((*MockCategoryRepository)(nil).DeleteCategoryByIDFromDB), arg0)
}

// FindAllCategoryFromDB mocks base method.
func (m *MockCategoryRepository) FindAllCategoryFromDB() ([]models.Category, *app_erros.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllCategoryFromDB")
	ret0, _ := ret[0].([]models.Category)
	ret1, _ := ret[1].(*app_erros.AppError)
	return ret0, ret1
}

// FindAllCategoryFromDB indicates an expected call of FindAllCategoryFromDB.
func (mr *MockCategoryRepositoryMockRecorder) FindAllCategoryFromDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllCategoryFromDB", reflect.TypeOf((*MockCategoryRepository)(nil).FindAllCategoryFromDB))
}

// GetCategoryFromDB mocks base method.
func (m *MockCategoryRepository) GetCategoryFromDB(arg0 string) (*models.Category, *app_erros.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryFromDB", arg0)
	ret0, _ := ret[0].(*models.Category)
	ret1, _ := ret[1].(*app_erros.AppError)
	return ret0, ret1
}

// GetCategoryFromDB indicates an expected call of GetCategoryFromDB.
func (mr *MockCategoryRepositoryMockRecorder) GetCategoryFromDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryFromDB", reflect.TypeOf((*MockCategoryRepository)(nil).GetCategoryFromDB), arg0)
}

// UpdateCategoryByIdFromDB mocks base method.
func (m *MockCategoryRepository) UpdateCategoryByIdFromDB(arg0 string, arg1 *models.Category) (bool, *app_erros.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategoryByIdFromDB", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*app_erros.AppError)
	return ret0, ret1
}

// UpdateCategoryByIdFromDB indicates an expected call of UpdateCategoryByIdFromDB.
func (mr *MockCategoryRepositoryMockRecorder) UpdateCategoryByIdFromDB(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategoryByIdFromDB", reflect.TypeOf((*MockCategoryRepository)(nil).UpdateCategoryByIdFromDB), arg0, arg1)
}
