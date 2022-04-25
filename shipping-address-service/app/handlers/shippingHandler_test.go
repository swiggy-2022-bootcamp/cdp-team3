package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	app_errors "github.com/cdp-team3/shipping-address-service/app-errors"
	"github.com/cdp-team3/shipping-address-service/domain/models"
	"github.com/cdp-team3/shipping-address-service/domain/services"
	"github.com/cdp-team3/shipping-address-service/mocks"
	"testing"
)


func TestShippingHandler_UpdateShippingAddress(t *testing.T) {

	shipping_id := "1234"
	shippingAddress := models.ShippingAddress{
	//	Id:   shipping_id,
		FirstName :"testFirstName",
		LastName  :"testLastName",
		City     :"testCity",
		Address1  :"testAddress1",
		Address2  :"testAddress2",
		CountryID :61,
		PostCode  :1222,
		// UserID     :"testUserId",
		// DefaultAddress  :"testDefaultAddress",
	}
	
	testCases := []struct {
		name          string
		buildStubs    func(shippingService *mocks.MockShippingService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		
		{
			// Missing shipping body in the request
			name: "BadRequestFailure",
			buildStubs: func(shippingService *mocks.MockShippingService) {
				shippingService.EXPECT().
				UpdateShippingAddressById(shipping_id,&shippingAddress).
					Times(0).
					Return(false,nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	
		{
			// Shipping updated successfully.
			name: "Success",
			buildStubs: func(shippingService *mocks.MockShippingService) {

				shippingService.EXPECT().
				UpdateShippingAddressById(shipping_id,&shippingAddress).
					Times(1).
					Return(true,nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusAccepted, recorder.Code)
			},
		},
		{
			//Shipping  failed with Internal server error.
			name: "Failure",
			buildStubs: func(shippingService *mocks.MockShippingService) {

				shippingService.EXPECT().
				UpdateShippingAddressById(shipping_id,&shippingAddress).
					Times(1).
					Return(false,app_errors.NewUnexpectedError(""))
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

			data, err := json.Marshal(shippingAddress)
			require.NoError(t, err)

			// Mock shipping service
			shippingService := mocks.NewMockShippingService(ctrl)
			tc.buildStubs(shippingService)

			server := NewServer(shippingService)

			// Making an HTTP request
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/shipping-service/api/shippingaddress/%s",shipping_id)
			var request *http.Request
			if tc.name == "BadRequestFailure" {
				request = httptest.NewRequest(http.MethodPut, url, nil)
			}   else {
				request = httptest.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			}
		
			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}
func TestShippingHandler_GetShippingAddress(t *testing.T) {
	shipping_id := "1234"
	shippingAddress := models.ShippingAddress{
		Id:   shipping_id,
		FirstName :"testFirstName",
		LastName  :"testLastName",
		City     :"testCity",
		Address1  :"testAddress1",
		Address2  :"testAddress2",
		CountryID :61,
		PostCode  :1222,
		UserID     :"testUserId",
		DefaultAddress  :"testDefaultAddress",
	}


	testCases := []struct {
		name          string
		buildStubs    func(shippingService *mocks.MockShippingService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "SuccessGetShippingAddress",
			buildStubs: func(shippingService *mocks.MockShippingService) {

				shippingService.EXPECT().
				FindShippingAddressById(shipping_id).
					Times(1).
					Return(&shippingAddress, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchShippingAddress(t, recorder.Body, shippingAddress)
			},
		},
		{
			name: "FailureShippingAddressNotFound",
			buildStubs: func(shippingService *mocks.MockShippingService) {
					shippingService.EXPECT().
					FindShippingAddressById(shipping_id).
					Times(1).
					Return(nil, app_errors.NewNotFoundError(""))

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailureInternalServerError",
			buildStubs: func(shippingService *mocks.MockShippingService) {
				shippingService.EXPECT().
				FindShippingAddressById(shipping_id).
					Times(1).
					Return(nil, app_errors.NewUnexpectedError(""))

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			shippingService := mocks.NewMockShippingService(ctrl)
			tc.buildStubs(shippingService)

			server := NewServer(shippingService)

			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/shipping-service/api/shippingaddress/%s",shipping_id)
			fmt.Println(url)
			request := httptest.NewRequest(http.MethodGet, url, nil)
			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}



func TestNewShippingHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	shippingRepository := mocks.NewMockShippingRepository(ctrl)

	shippingService := services.NewShippingServiceImpl(shippingRepository)
	shippingHandler := NewShippingHandler(shippingService)

	assert.Equal(t, shippingHandler.shippingService, shippingService)
}

func requireBodyMatchShippingAddress(t *testing.T, body *bytes.Buffer, requiredResponse models.ShippingAddress) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var responseReceived models.ShippingAddress
	err = json.Unmarshal(data, &responseReceived)
	require.NoError(t, err)
	require.Equal(t, responseReceived, requiredResponse)
}