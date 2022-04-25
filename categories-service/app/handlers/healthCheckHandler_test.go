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
	"github.com/cdp-team3/categories-service/mocks"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(categoryRepository *mocks.MockCategoryRepository)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Success",
			buildStubs: func(categoryRepository *mocks.MockCategoryRepository) {
				// Call to MOCK DBHealthCheck() returning true.
				categoryRepository.EXPECT().
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
			buildStubs: func(categoryRepository *mocks.MockCategoryRepository) {
				// Call to MOCK DBHealthCheck() returning false.
				categoryRepository.EXPECT().
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
			categoryRepository := mocks.NewMockCategoryRepository(ctrl)
			tc.buildStubs(categoryRepository)

			//Creating handler and setting up router
			healthCheckHandler := NewHealthCheckHandler(categoryRepository)
			server := gin.Default()
			router := server.Group("categories/api")
			router.GET("/", healthCheckHandler.HealthCheck)

			// Making an HTTP call and recording the response
			recorder := httptest.NewRecorder()
			url := "/categories/api/"
			request := httptest.NewRequest(http.MethodGet, url, nil)

			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

func TestNewHealthCheckHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	categoriesRepository := mocks.NewMockCategoryRepository(ctrl)

	healthCheckHandler := NewHealthCheckHandler(categoriesRepository)
	assert.Equal(t, healthCheckHandler.categoryRepository, categoriesRepository)
}

func requireBodyMatchResposne(t *testing.T, actualResponse *bytes.Buffer, expectedResponse HealthCheckResponse) {
	data, err := ioutil.ReadAll(actualResponse)
	require.NoError(t, err)

	var responseReceived HealthCheckResponse
	err = json.Unmarshal(data, &responseReceived)
	require.NoError(t, err)
	require.Equal(t, responseReceived, expectedResponse)
}