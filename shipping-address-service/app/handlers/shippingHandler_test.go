package handlers

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"github.com/golang/mock/gomock"
	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	app_errors "github.com/cdp-team3/shipping-address-service/app-errors"
	"github.com/cdp-team3/shipping-address-service/domain/models"
	// "github.com/cdp-team3/shipping-address-service/domain/services"
	"github.com/cdp-team3/shipping-address-service/mocks"
	"testing"
)

func TestShippingHandler_AddShippingAddress(t *testing.T) {

	//ShippingAddressID := "bb912edc-50d9-42d7-b7a1-9ce66d459thj"
	var ShippingAddress *models.ShippingAddress
		


//	ShippingAddress.Id= "bb912edc-50d9-42d7-b7a1-9ce66d459thj"
	
    // ShippingAddress := models.ShippingAddress{
		
	// 	Id: "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
	// 	FirstName: "naveen",
	// 	LastName: "Sharma",
	// 	City: "Banglore",
	// 	Address1: "Address1",
	// 	Address2: "Address2",
	// 	CountryID: 81,
	// 	PostCode: 560012,
	
	
}
shippingAddressValidationError:= models.ShippingAddress{
		
	Id: "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
	FirstName: "naveen",
	City: "Banglore",
	CountryID: 81,
	PostCode: 560012,


}
	testCases := []struct {
		name          string
		buildStubs    func(shippingService *mocks.MockShippingService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		// {
		// 	// Missing payment mode body in the request
		// 	name: "BadRequestFailure",
		// 	buildStubs: func(shippingService *mocks.MockShippingService) {
		// 		shippingService.EXPECT().
		// 		InsertShippingAddress(*ShippingAddress).
		// 			Times(0).
		// 			Return(nil)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusBadRequest, recorder.Code)
		// 	},
		// },
		// {
		// 	// Required field not present in request body.
		// 	name: "ValidationFailure",
		// 	buildStubs: func(shippingService *mocks.MockShippingService) {
		// 		shippingService.EXPECT().
		// 	    InsertShippingAddress(*ShippingAddress).
		// 			Times(0).
		// 			Return(nil)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusBadRequest, recorder.Code)
		// 	},
		// },
		// {
		// 	// Shipping Address added successfully.
		// 	name: "Success",
		// 	buildStubs: func(shippingService *mocks.MockShippingService) {
		// 		shippingService.EXPECT().
		// 		InsertShippingAddress(ShippingAddress).
		// 			Times(1).
		// 			Return(nil)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusOK, recorder.Code)
		// 	},
		// },
		{
			// Shipping Address addition failed with Internal server error.
			name: "Failure",
			buildStubs: func(shippingService *mocks.MockShippingService) {
				shippingService.EXPECT().
				InsertShippingAddress(ShippingAddress).
					Times(1).
					Return(app_errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			data, err := json.Marshal(&ShippingAddress)
			require.NoError(t, err)

			// Mock shipping service
			shippingService := mocks.NewMockShippingService(ctrl)
			tc.buildStubs(shippingService)

			server := NewServer(shippingService)

			// Making an HTTP request
			recorder := httptest.NewRecorder()
			url := "/shipping-service/api/shippingaddress"
			var request *http.Request
			if tc.name == "BadRequestFailure" {
				request = httptest.NewRequest(http.MethodPost, url, nil)
			} else if tc.name == "ValidationFailure" {
				data, err := json.Marshal(shippingAddressValidationError)
				require.NoError(t, err)
				request = httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			} else {
				request = httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			}

			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}


func requireBodyMatchPaymentMode(t *testing.T, body *bytes.Buffer, requiredResponse models.ShippingAddress) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var responseReceived models.ShippingAddress
	err = json.Unmarshal(data, &responseReceived)
	require.NoError(t, err)
	require.Equal(t, responseReceived, requiredResponse)
}