package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"github.com/cdp-team3/shipping-address-service/mocks"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(shippingRepository *mocks.MockShippingRepository)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Success",
			buildStubs: func(shippingRepository *mocks.MockShippingRepository) {
				// Call to MOCK DBHealthCheck() returning true.
				shippingRepository.EXPECT().
					DBHealthCheck().
					Times(1).
					Return(true)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchResposne(t, recorder.Body, HealthCheckResponse{Server: "Server is up",
					Database: "Database is up"})
			},
		},
		{
			name: "InternalServerError",
			buildStubs: func(shippingRepository *mocks.MockShippingRepository) {
				// Call to MOCK DBHealthCheck() returning false.
				shippingRepository.EXPECT().
					DBHealthCheck().
					Times(1).
					Return(false)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
				requireBodyMatchResposne(t, recorder.Body, HealthCheckResponse{Server: "Server is up",
					Database: "Database is down"})
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//Creating an object obj mock repository.
			shippingRepository := mocks.NewMockShippingRepository(ctrl)
			tc.buildStubs(shippingRepository)

			//Creating handler and setting up router
			healthCheckHandler := NewHealthCheckHandler(shippingRepository)
			server := gin.Default()
			router := server.Group("shipping-service/api")
			router.GET("/", healthCheckHandler.HealthCheck)

			// Making an HTTP call and recording the response
			recorder := httptest.NewRecorder()
			url := "/shipping-service/api/"
			request := httptest.NewRequest(http.MethodGet, url, nil)

			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

func TestNewHealthCheckHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	shippingRepository := mocks.NewMockShippingRepository(ctrl)

	healthCheckHandler := NewHealthCheckHandler(shippingRepository)
	assert.Equal(t, healthCheckHandler.shippingRepository, shippingRepository)
}

func requireBodyMatchResposne(t *testing.T, actualResponse *bytes.Buffer, expectedResponse HealthCheckResponse) {
	data, err := ioutil.ReadAll(actualResponse)
	require.NoError(t, err)

	var responseReceived HealthCheckResponse
	err = json.Unmarshal(data, &responseReceived)
	require.NoError(t, err)
	require.Equal(t, responseReceived, expectedResponse)
}