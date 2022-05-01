package services

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"
)

func TestCustomerServiceImpl_GetCustomerById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	customer_id := "423fec6b-82c-4f99-8b2b-6eeef7605a37"
	address := models.Address{

		ShippingAddressId: "123",
		Address1:          "Address1",
		Address2:          "Address2",
		City:              "hyderabad",
		CountryID:         22,
		PostCode:          500012,
		Default:           "0",
	}

	successCustomer := &models.Customer{
		CustomerId:        customer_id,
		IsAdmin:           false,
		Firstname:         "ExampleFirstName",
		Lastname:          "ExampleLastName",
		Username:          "ExampleUsername",
		Password:          "$2b$10$//DXiVVE59p7G5k/4Klx/ezF7BI42QZKmoOD0NDvUuqxRE5bFF",
		ConfirmPassword:   "$2b$10$//DXiVVE59p7G5k/4Klx/ezF7BI42QZKmoOD0NDvUuqxRE5bFF",
		Email:             "example@gmail.com",
		Telephone:         "123-454-6673",
		Address:           address,
		Status:            "1",
		Approved:          "1",
		DateAdded:         time.Now(),
		Rewards:           500,
		TransactionPoints: 0,
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockCustomerRepository)
		checkResponse func(t *testing.T, customer interface{}, err interface{})
	}{
		{
			name: "successCustomerFound",
			buildStubs: func(repository *mocks.MockCustomerRepository) {
				repository.EXPECT().
					GetCustomerByIdFromDB(customer_id).
					Times(1).
					Return(successCustomer, nil)
			},
			checkResponse: func(t *testing.T, customer interface{}, err interface{}) {
				assert.NotNil(t, customer)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailureCustomerNotFound",
			buildStubs: func(repository *mocks.MockCustomerRepository) {
				repository.EXPECT().
					GetCustomerByIdFromDB(customer_id).
					Times(1).
					Return(nil, errors.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, customer interface{}, err interface{}) {
				assert.Nil(t, customer)
				assert.NotNil(t, err)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockCustomerRepository) {
				repository.EXPECT().
					GetCustomerByIdFromDB(customer_id).
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, customer interface{}, err interface{}) {

				assert.Nil(t, customer)
				assert.NotNil(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			customerRepository := mocks.NewMockCustomerRepository(ctrl)
			tc.buildStubs(customerRepository)

			customerServiceImpl := NewCustomerServiceImpl(customerRepository)
			customer, err := customerServiceImpl.GetCustomerById(customer_id)
			tc.checkResponse(t, customer, err)
		})
	}
}

func TestCustomerServiceImpl_GetCustomerByEmail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	customer_email := "example@gmail.com"
	address := models.Address{

		ShippingAddressId: "123",
		Address1:          "Address1",
		Address2:          "Address2",
		City:              "hyderabad",
		CountryID:         22,
		PostCode:          500012,
		Default:           "0",
	}

	successCustomer := &models.Customer{
		CustomerId:        "423fec6b-82c-4f99-8b2b-6eeef7605a37",
		IsAdmin:           false,
		Firstname:         "ExampleFirstName",
		Lastname:          "ExampleLastName",
		Username:          "ExampleUsername",
		Password:          "$2b$10$//DXiVVE59p7G5k/4Klx/ezF7BI42QZKmoOD0NDvUuqxRE5bFF",
		ConfirmPassword:   "$2b$10$//DXiVVE59p7G5k/4Klx/ezF7BI42QZKmoOD0NDvUuqxRE5bFF",
		Email:             customer_email,
		Telephone:         "123-454-6673",
		Address:           address,
		Status:            "1",
		Approved:          "1",
		DateAdded:         time.Now(),
		Rewards:           500,
		TransactionPoints: 0,
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockCustomerRepository)
		checkResponse func(t *testing.T, customer interface{}, err interface{})
	}{
		{
			name: "successCustomerFound",
			buildStubs: func(repository *mocks.MockCustomerRepository) {
				repository.EXPECT().
					GetCustomerByEmailFromDB(customer_email).
					Times(1).
					Return(successCustomer, nil)
			},
			checkResponse: func(t *testing.T, customer interface{}, err interface{}) {
				assert.NotNil(t, customer)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailureCustomerNotFound",
			buildStubs: func(repository *mocks.MockCustomerRepository) {
				repository.EXPECT().
					GetCustomerByEmailFromDB(customer_email).
					Times(1).
					Return(nil, errors.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, customer interface{}, err interface{}) {
				assert.Nil(t, customer)
				assert.NotNil(t, err)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockCustomerRepository) {
				repository.EXPECT().
					GetCustomerByEmailFromDB(customer_email).
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, customer interface{}, err interface{}) {

				assert.Nil(t, customer)
				assert.NotNil(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			customerRepository := mocks.NewMockCustomerRepository(ctrl)
			tc.buildStubs(customerRepository)

			customerServiceImpl := NewCustomerServiceImpl(customerRepository)
			customer, err := customerServiceImpl.GetCustomerByEmail(customer_email)
			tc.checkResponse(t, customer, err)
		})
	}
}

func TestCustomerServiceImpl_GetAllCustomers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	address := models.Address{

		ShippingAddressId: "123",
		Address1:          "Address1",
		Address2:          "Address2",
		City:              "hyderabad",
		CountryID:         22,
		PostCode:          500012,
		Default:           "0",
	}
	address2 := models.Address{

		ShippingAddressId: "1234",
		Address1:          "Address12",
		Address2:          "Address22",
		City:              "hyderabad",
		CountryID:         22,
		PostCode:          500013,
		Default:           "0",
	}
	successCustomers := []models.Customer{
		{
			CustomerId:        "423fec6b-8a0c-4f99-8b2b-6eeef5478t54",
			IsAdmin:           false,
			Firstname:         "ExampleFirstName",
			Lastname:          "ExampleLastName",
			Username:          "ExampleUsername",
			Password:          "$2b$10$//DXiVVE59p7G5k/4Klx/ezF7BI42QZKmoOD0NDvUuqxRE5bFF",
			ConfirmPassword:   "$2b$10$//DXiVVE59p7G5k/4Klx/ezF7BI42QZKmoOD0NDvUuqxRE5bFF",
			Email:             "example@gmail.com",
			Telephone:         "123-454-6673",
			Address:           address,
			Status:            "1",
			Approved:          "1",
			DateAdded:         time.Now(),
			Rewards:           500,
			TransactionPoints: 0,
		},
		{
			CustomerId:        "423fec6b-8a0c-4f99-8b2b-6eeef445",
			IsAdmin:           false,
			Firstname:         "ExampleFirstName1",
			Lastname:          "ExampleLastName1",
			Username:          "ExampleUsername1",
			Password:          "$2b$10$//DXiVVE59p7G5k/4Klx/ezF7BI42QZKmoOD0NDvUuqxRE5erT",
			ConfirmPassword:   "$2b$10$//DXiVVE59p7G5k/4Klx/ezF7BI42QZKmoOD0NDvUuqxRE5erT",
			Email:             "example1@gmail.com",
			Telephone:         "123-454-6674",
			Address:           address2,
			Status:            "1",
			Approved:          "1",
			DateAdded:         time.Now(),
			Rewards:           500,
			TransactionPoints: 0,
		},
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockCustomerRepository)
		checkResponse func(t *testing.T, res []models.Customer, err interface{})
	}{
		{
			name: "SuccessCustomersFound",
			buildStubs: func(repository *mocks.MockCustomerRepository) {
				repository.EXPECT().
					GetAllCustomersFromDB().
					Times(1).
					Return(successCustomers, nil)
			},
			checkResponse: func(t *testing.T, res []models.Customer, err interface{}) {
				assert.NotNil(t, res)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailureCustomersNotFound",
			buildStubs: func(repository *mocks.MockCustomerRepository) {
				repository.EXPECT().
					GetAllCustomersFromDB().
					Times(1).
					Return(nil, errors.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, res []models.Customer, err interface{}) {
				assert.Nil(t, res)
				assert.NotNil(t, err)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockCustomerRepository) {
				repository.EXPECT().
					GetAllCustomersFromDB().
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, res []models.Customer, err interface{}) {
				assert.Nil(t, res)
				assert.NotNil(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			customerRepository := mocks.NewMockCustomerRepository(ctrl)
			tc.buildStubs(customerRepository)

			customerServiceImpl := NewCustomerServiceImpl(customerRepository)
			res, err := customerServiceImpl.GetAllCustomers()
			tc.checkResponse(t, res, err)
		})
	}
}

func TestCustomerServiceImpl_DeleteCustomerById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	customer_id := "423fec6b-8a0c-4f99-8b2b-6eeef445"

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockCustomerRepository)
		checkResponse func(t *testing.T, result bool, err interface{})
	}{
		{
			name: "SuccessCustomerDeleted",
			buildStubs: func(repository *mocks.MockCustomerRepository) {
				repository.EXPECT().
					DeleteCustomerByIdFromDB(customer_id).
					Times(1).
					Return(true, nil)
			},
			checkResponse: func(t *testing.T, result bool, err interface{}) {
				assert.Equal(t, true, result)
				assert.Nil(t, err)

			},
		},
		{
			name: "FailureCustomerNotDeleted",
			buildStubs: func(repository *mocks.MockCustomerRepository) {
				repository.EXPECT().
					DeleteCustomerByIdFromDB(customer_id).
					Times(1).
					Return(false, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, res bool, err interface{}) {
				assert.Equal(t, false, res)
				assert.NotNil(t, err)

			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockCustomerRepository) {
				repository.EXPECT().
					DeleteCustomerByIdFromDB(customer_id).
					Times(1).
					Return(false, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, res bool, err interface{}) {
				assert.Equal(t, false, res)
				assert.NotNil(t, err)

			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			customerRepository := mocks.NewMockCustomerRepository(ctrl)
			tc.buildStubs(customerRepository)

			customerServiceImpl := NewCustomerServiceImpl(customerRepository)
			value, err := customerServiceImpl.DeleteCustomerById(customer_id)
			tc.checkResponse(t, value, err)
		})
	}
}

func TestCustomerServiceImpl_UpdateCustomerById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	customer_id := "423fec6b-8a0c-4f99-8b2b-6eeef445"
	address := models.Address{

		ShippingAddressId: "123",
		Address1:          "Address1",
		Address2:          "Address2",
		City:              "hyderabad",
		CountryID:         22,
		PostCode:          500012,
		Default:           "0",
	}
	updatedCustomer := &models.Customer{

		CustomerId:        customer_id,
		IsAdmin:           false,
		Firstname:         "ExampleFirstName",
		Lastname:          "ExampleLastName",
		Username:          "ExampleUsername",
		Password:          "$2b$10$//DXiVVE59p7G5k/4Klx/ezF7BI42QZKmoOD0NDvUuqxRE5bFF",
		ConfirmPassword:   "$2b$10$//DXiVVE59p7G5k/4Klx/ezF7BI42QZKmoOD0NDvUuqxRE5bFF",
		Email:             "example@gmail.com",
		Telephone:         "123-454-6679",
		Address:           address,
		Status:            "1",
		Approved:          "1",
		DateAdded:         time.Now(),
		Rewards:           500,
		TransactionPoints: 0,
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockCustomerRepository)
		checkResponse func(t *testing.T, customer interface{}, err interface{})
	}{
		{
			name: "SuccessCustomerUpdated",
			buildStubs: func(repository *mocks.MockCustomerRepository) {
				repository.EXPECT().
					UpdateCustomerByIdFromDB(customer_id, updatedCustomer).
					Times(1).
					Return(updatedCustomer, nil)
			},
			checkResponse: func(t *testing.T, customer interface{}, err interface{}) {
				assert.NotNil(t, customer)
				assert.Equal(t, customer, updatedCustomer)
				assert.Nil(t, err)
			},
		},
		{
			name: "FailureCustomerNotUpdated",
			buildStubs: func(repository *mocks.MockCustomerRepository) {
				repository.EXPECT().
					UpdateCustomerByIdFromDB(customer_id, updatedCustomer).
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, customer interface{}, err interface{}) {
				assert.Nil(t, customer)
				assert.NotNil(t, err)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockCustomerRepository) {
				repository.EXPECT().
					UpdateCustomerByIdFromDB(customer_id, updatedCustomer).
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, customer interface{}, err interface{}) {
				assert.Nil(t, customer)
				assert.NotNil(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			customerRepository := mocks.NewMockCustomerRepository(ctrl)
			tc.buildStubs(customerRepository)

			customerServiceImpl := NewCustomerServiceImpl(customerRepository)
			value, err := customerServiceImpl.UpdateCustomerById(customer_id, updatedCustomer)
			tc.checkResponse(t, value, err)
		})
	}
}
