package services

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
     app_errors "github.com/cdp-team3/shipping-address-service/app-errors"
	"github.com/cdp-team3/shipping-address-service/domain/models"
	"github.com/cdp-team3/shipping-address-service/mocks"
	"github.com/cdp-team3/shipping-address-service/app/grpcs/shipping_checkout"
	"testing"
)


func TestShippingProtoServer_GetShippingAddress(t *testing.T) {
	gin.SetMode(gin.TestMode)
	request := &shipping_checkout.ShippingAddressRequest{
		ShippingAddressID: "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
	}
	shippingAddress := &shipping_checkout.ShippingAddressResponse{
	
		Firstname: "naveen",
		Lastname: "Sharma",
		City: "Banglore",
		Address1: "Address1",
		Address2: "Address2",
		Countryid: 81,
		Postcode: 560012,
	}
	
	shippingAddressFromDbResponse := &models.ShippingAddress{
		
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
		checkResponse func(t *testing.T, shippingAddress *shipping_checkout.ShippingAddressResponse, err error)
	}{
		{
			name: "SuccessShippingAddressFound",
			buildStubs: func(repository *mocks.MockShippingRepository) {
				repository.EXPECT().
				    FindShippingAddressByIdFromDB(request.GetShippingAddressID()).
					Times(1).
					Return(shippingAddressFromDbResponse, nil)
			},
			checkResponse: func(t *testing.T, address *shipping_checkout.ShippingAddressResponse, err error) {
				require.Equal(t, shippingAddress,shippingAddress)
				require.Nil(t, err)
			},
		},
		{
			name: "FailureShippingAddressNotFound",
			buildStubs: func(repository *mocks.MockShippingRepository) {
				repository.EXPECT().
				FindShippingAddressByIdFromDB(request.GetShippingAddressID()).
					Times(1).
					Return(nil, app_errors.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, address *shipping_checkout.ShippingAddressResponse, err error) {
				require.Equal(t, app_errors.NewNotFoundError("").Error(), err)
				require.NotNil(t,err)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockShippingRepository) {
				repository.EXPECT().
				FindShippingAddressByIdFromDB(request.GetShippingAddressID()).
					Times(1).
					Return(nil, app_errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, address *shipping_checkout.ShippingAddressResponse, err error) {
				require.Equal(t, app_errors.NewUnexpectedError("").Error(), err)
				require.NotNil(t,err)
				
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

			shippingProtoServer := NewShippingProtoService(shippingRepository)
			address, err :=shippingProtoServer.GetShippingAddress(context.Background(), request)
			tc.checkResponse(t, address, err)
		})
	}
}

