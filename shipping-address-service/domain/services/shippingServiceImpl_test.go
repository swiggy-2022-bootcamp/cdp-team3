package services

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	app_erros "github.com/cdp-team3/shipping-address-service/app-errors"
	"github.com/cdp-team3/shipping-address-service/domain/models"
	"github.com/cdp-team3/shipping-address-service/mocks"
    "testing"
)
func TestTransactionServiceImpl_FindShippingAddressById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ShippingAddressID := "bb912edc-50d9-42d7-b7a1-9ce66d459thj"

	successShippingAddress := &models.ShippingAddress{
		
			Id: "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
			FirstName: "naveen",
			LastName: "Sharma",
			City: "Banglore",
			Address1: "Address1",
			Address2: "Address2",
			CountryID: 81,
			PostCode: 560012,
		
		
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockShippingRepository)
		checkResponse func(t *testing.T, shippingAddress interface{}, err interface{})
	}{
		{
			name: "SuccessShippingAddressFound",
			buildStubs: func(repository *mocks.MockShippingRepository) {
				repository.EXPECT().
				    FindShippingAddressByIdFromDB(ShippingAddressID).
					Times(1).
					Return(successShippingAddress, nil)
			},
			checkResponse: func(t *testing.T, shippingAddress interface{}, err interface{}) {
				
				assert.NotNil(t,shippingAddress)
				assert.Nil(t,err)
			},
		},
		{
			name: "FailureShippingAddressNotFound",
			buildStubs: func(repository *mocks.MockShippingRepository) {
				repository.EXPECT().
				FindShippingAddressByIdFromDB(ShippingAddressID).
					Times(1).
					Return(nil, app_erros.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, shippingAddress interface{}, err interface{}) {
	
			assert.Nil(t,shippingAddress)
				assert.NotNil(t,err)
		
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockShippingRepository) {
				repository.EXPECT().
				    FindShippingAddressByIdFromDB(ShippingAddressID).
					Times(1).
					Return(nil, app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, shippingAddress interface{}, err interface{}) {
			
				assert.Nil(t,shippingAddress)
				assert.NotNil(t,err)
		
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			shippingRepository := mocks.NewMockShippingRepository(ctrl)
			tc.buildStubs(shippingRepository)

			shippingServiceImpl := NewShippingServiceImpl(shippingRepository)
			shippingAddress, err := shippingServiceImpl.FindShippingAddressById(ShippingAddressID)
			tc.checkResponse(t,  shippingAddress, err)
		})
	}
}

func TestTransactionServiceImpl_AddShippingAddress(t *testing.T) {
	gin.SetMode(gin.TestMode)
	shippingAddress := &models.ShippingAddress{
		
		Id: "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
		FirstName: "naveen",
		LastName: "Sharma",
		City: "Banglore",
		Address1: "Address1",
		Address2: "Address2",
		CountryID: 81,
		PostCode: 560012,
	
	
}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockShippingRepository)
		checkResponse func(t *testing.T,err interface{})
	}{
	
		{
			name: "SuccessAddShippingAddress",
			buildStubs: func(repository *mocks.MockShippingRepository) {
				repository.EXPECT().
				InsertShippingAddressToDB(shippingAddress).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T,err interface{}) {
				assert.Nil(t,err)
			},
		},
		{
			name: "FailureAddShippingAddress",
			buildStubs: func(repository *mocks.MockShippingRepository) {
				repository.EXPECT().
				    InsertShippingAddressToDB(shippingAddress).
					Times(1).
					Return(app_erros.NewUnexpectedError(""))

			},
			checkResponse: func(t *testing.T,err interface{}) {
				assert.NotNil(t,err)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockShippingRepository) {
				repository.EXPECT().
				InsertShippingAddressToDB(shippingAddress).
					Times(1).
					Return(app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, err interface{}) {
				assert.NotNil(t,err)
			
			},
		},
	
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			shippingRepository := mocks.NewMockShippingRepository(ctrl)
			tc.buildStubs(shippingRepository)

			shippingServiceImpl := NewShippingServiceImpl(shippingRepository)
			err := shippingServiceImpl.InsertShippingAddress(shippingAddress)
			tc.checkResponse(t, err)
		})
	}
}

func TestShippingServiceImpl_DeleteShippingAddressById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ShippingAddressID := "bb912edc-50d9-42d7-b7a1-9ce66d459thj"
	

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockShippingRepository)
		checkResponse func(t *testing.T, result bool , err interface{})
	}{
		{
			name: "SuccessShippingAddressDeleted",
			buildStubs: func(repository *mocks.MockShippingRepository) {
				repository.EXPECT().
				    DeleteShippingAddressByIdFromDB(ShippingAddressID).
					Times(1).
					Return(true, nil)
			},
			checkResponse: func(t *testing.T, result bool, err interface{}) {
				assert.Equal(t, true, result)
				assert.Nil(t,err)
			
			},
		},
		{
			name: "FailureShippingAddressNotDeleted",
			buildStubs: func(repository *mocks.MockShippingRepository) {
				repository.EXPECT().
				DeleteShippingAddressByIdFromDB(ShippingAddressID).
					Times(1).
					Return(false, app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, res bool, err interface{}) {
				assert.Equal(t, false, res)
				assert.NotNil(t,err)
				
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockShippingRepository) {
				repository.EXPECT().
				DeleteShippingAddressByIdFromDB(ShippingAddressID).
					Times(1).
					Return(false, app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, res bool, err interface{}) {
				assert.Equal(t, false, res)
				assert.NotNil(t,err)

			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			shippingRepository := mocks.NewMockShippingRepository(ctrl)
			tc.buildStubs(shippingRepository)

			shippingServiceImpl := NewShippingServiceImpl(shippingRepository)
			value, err := shippingServiceImpl.DeleteShippingAddressById(ShippingAddressID)
			tc.checkResponse(t,value, err)
		})
	}
}
func TestShippingServiceImpl_UpdateShippingAddressById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ShippingAddressID := "bb912edc-50d9-42d7-b7a1-9ce66d459thj"

	shippingAddress := &models.ShippingAddress{
		
			Id: "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
			FirstName: "naveen",
			LastName: "Sharma",
			City: "Banglore",
			Address1: "Address1",
			Address2: "Address2",
			CountryID: 81,
			PostCode: 560012,
		
		
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockShippingRepository)
		checkResponse func(t *testing.T, result bool , err interface{})
	}{
		{
			name: "SuccessShippingAddressUpdated",
			buildStubs: func(repository *mocks.MockShippingRepository) {
				repository.EXPECT().
				    UpdateShippingAddressByIdFromDB(ShippingAddressID,shippingAddress).
					Times(1).
					Return(true, nil)
			},
			checkResponse: func(t *testing.T, result bool, err interface{}) {
				assert.Equal(t, true, result)
				assert.Nil(t,err)
				
			},
		},
		{
			name: "FailureShippingAddressNotUpdated",
			buildStubs: func(repository *mocks.MockShippingRepository) {
				repository.EXPECT().
				UpdateShippingAddressByIdFromDB(ShippingAddressID,shippingAddress).
					Times(1).
					Return(false, app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, res bool, err interface{}) {
				assert.Equal(t, false, res)
				assert.NotNil(t,err)
			
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockShippingRepository) {
				repository.EXPECT().
				UpdateShippingAddressByIdFromDB(ShippingAddressID,shippingAddress).
					Times(1).
					Return(false, app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, res bool, err interface{}) {
				assert.Equal(t, false, res)
				assert.NotNil(t,err)
			
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			shippingRepository := mocks.NewMockShippingRepository(ctrl)
			tc.buildStubs(shippingRepository)

			shippingServiceImpl := NewShippingServiceImpl(shippingRepository)
			value, err := shippingServiceImpl.UpdateShippingAddressById(ShippingAddressID,shippingAddress)
			tc.checkResponse(t,value, err)
		})
	}
}
