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
	app_errors "github.com/cdp-team3/categories-service/app-errors"
	"github.com/cdp-team3/categories-service/domain/models"
	"github.com/cdp-team3/categories-service/domain/services"
	"github.com/cdp-team3/categories-service/mocks"
	"testing"
)

// Test Handler : Add Category 
func TestCategoryHandler_AddCategory(t *testing.T) {


	categorydesc:=models.CategoryDescription{
		Name           : "testName",

		Description     : "testDescription" ,
		MetaDescription  : "testMetaDescription",
		MetaKeyword      : "testMetaKeyword",
		MetaTitle       : "testMetaTitle",
	}

	category := models.Category{

	CategoryDescription: []models.CategoryDescription{categorydesc}	,
	}

	testCases := []struct {
		name          string
		buildStubs    func(categoryService *mocks.MockCategoryService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		
		{
			// Missing category body in the request
			name: "BadRequestFailure",
			buildStubs: func(categoryService *mocks.MockCategoryService) {
				categoryService.EXPECT().
				AddCategory(&category).
					Times(0).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	
		{
			// Category added successfully.
			name: "Success",
			buildStubs: func(categoryService *mocks.MockCategoryService) {

				categoryService.EXPECT().
				AddCategory(&category).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			// Category addition failed with Internal server error.
			name: "Failure",
			buildStubs: func(categoryService *mocks.MockCategoryService) {

				categoryService.EXPECT().
				AddCategory(&category).
					Times(1).
					Return(app_errors.NewUnexpectedError(""))
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

			data, err := json.Marshal(category)
			require.NoError(t, err)

			// Mock category service
			categoryService := mocks.NewMockCategoryService(ctrl)
			tc.buildStubs(categoryService)

			server := NewServer(categoryService)

			// Making an HTTP request
			recorder := httptest.NewRecorder()
			url := "/categories/api/categories"
			var request *http.Request
			if tc.name == "BadRequestFailure" {
				request = httptest.NewRequest(http.MethodPost, url, nil)
			}   else {
				request = httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			}
		
			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

// Test Handler : Update Category
func TestCategoryHandler_UpdateCategory(t *testing.T) {

	category_id := "1232"
	categorydesc:=models.CategoryDescription{
		Name           : "testName",

		Description     : "testDescription" ,
		MetaDescription  : "testMetaDescription",
		MetaKeyword      : "testMetaKeyword",
		MetaTitle       : "testMetaTitle",
	}
	
	category := models.Category{

	CategoryDescription: []models.CategoryDescription{categorydesc}	,
	}

	testCases := []struct {
		name          string
		buildStubs    func(categoryService *mocks.MockCategoryService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		
		{
			// Missing category body in the request
			name: "BadRequestFailure",
			buildStubs: func(categoryService *mocks.MockCategoryService) {
				categoryService.EXPECT().
				UpdateCategoryByID(category_id,&category).
					Times(0).
					Return(false,nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	
		{
			// Category updated successfully.
			name: "Success",
			buildStubs: func(categoryService *mocks.MockCategoryService) {

				categoryService.EXPECT().
				UpdateCategoryByID(category_id,&category).
					Times(1).
					Return(true,nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			// Category addition failed with Internal server error.
			name: "Failure",
			buildStubs: func(categoryService *mocks.MockCategoryService) {

				categoryService.EXPECT().
				UpdateCategoryByID(category_id,&category).
					Times(1).
					Return(false,app_errors.NewUnexpectedError(""))
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

			data, err := json.Marshal(category)
			require.NoError(t, err)

			// Mock category service
			categoryService := mocks.NewMockCategoryService(ctrl)
			tc.buildStubs(categoryService)

			server := NewServer(categoryService)

			// Making an HTTP request
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/categories/api/categories/%s",category_id)
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

// Test Handler : Get Category
func TestCategoryHandler_GetCategory(t *testing.T) {
	category_id := "1234"
	categorydesc:=models.CategoryDescription{
		Name           : "testName",

		Description     : "testDescription" ,
		MetaDescription  : "testMetaDescription",
		MetaKeyword      : "testMetaKeyword",
		MetaTitle       : "testMetaTitle",
	}
	category := models.Category{
		CategoryId : category_id,
		CategoryDescription: []models.CategoryDescription{categorydesc},
	}


	testCases := []struct {
		name          string
		buildStubs    func(categoryService *mocks.MockCategoryService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "SuccessGetCategory",
			buildStubs: func(categoryService *mocks.MockCategoryService) {

				categoryService.EXPECT().
			     	GetCategory(category_id).
					Times(1).
					Return(&category, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCategory(t, recorder.Body, category)
			},
		},
		{
			name: "FailureCategoryNotFound",
			buildStubs: func(categoryService *mocks.MockCategoryService) {
					categoryService.EXPECT().
					GetCategory(category_id).
					Times(1).
					Return(nil, app_errors.NewNotFoundError(""))

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailureInternalServerError",
			buildStubs: func(categoryService *mocks.MockCategoryService) {
				categoryService.EXPECT().
				GetCategory(category_id).
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

			categoryService := mocks.NewMockCategoryService(ctrl)
			tc.buildStubs(categoryService)

			server := NewServer(categoryService)

			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/categories/api/categories/%s",category_id)
			fmt.Println(url)
			request := httptest.NewRequest(http.MethodGet, url, nil)
			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

// Test Handler : Delete Category
func TestCategoryHandler_DeleteCategory(t *testing.T) {
	category_id := "1234"
	


	testCases := []struct {
		name          string
		buildStubs    func(categoryService *mocks.MockCategoryService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "SuccessDeleteCategory",
			buildStubs: func(categoryService *mocks.MockCategoryService) {

				categoryService.EXPECT().
				DeleteCategoryByID(category_id).
					Times(1).
					Return(true, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				
			},
		},
		{
			name: "FailureCategoryNotFound",
			buildStubs: func(categoryService *mocks.MockCategoryService) {
					categoryService.EXPECT().
					DeleteCategoryByID(category_id).
					Times(1).
					Return(false, app_errors.NewNotFoundError(""))

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailureInternalServerError",
			buildStubs: func(categoryService *mocks.MockCategoryService) {
				categoryService.EXPECT().
				DeleteCategoryByID(category_id).
					Times(1).
					Return(false, app_errors.NewUnexpectedError(""))

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

			categoryService := mocks.NewMockCategoryService(ctrl)
			tc.buildStubs(categoryService)

			server := NewServer(categoryService)

			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/categories/api/categories/%s",category_id)
			fmt.Println(url)
			request := httptest.NewRequest(http.MethodDelete, url, nil)
			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

// Test Handler : Get All Category
func TestCategoryHandler_GetAllCategory(t *testing.T) {
	category_id1 := "1234"
	category_id2 := "1232"
	categorydesc1:=models.CategoryDescription{
		Name           : "testName",

		Description     : "testDescription" ,
		MetaDescription  : "testMetaDescription",
		MetaKeyword      : "testMetaKeyword",
		MetaTitle       : "testMetaTitle",
	}
	categorydesc2:=models.CategoryDescription{
		Name           : "testName",

		Description     : "testDescription" ,
		MetaDescription  : "testMetaDescription",
		MetaKeyword      : "testMetaKeyword",
		MetaTitle       : "testMetaTitle",
	}
	category1 := models.Category{
		CategoryId : category_id1,
		CategoryDescription: []models.CategoryDescription{categorydesc1},
	}
	category2 := models.Category{
		CategoryId : category_id2,
		CategoryDescription: []models.CategoryDescription{categorydesc2},
	}
   categoryList:=[]models.Category{category1,category2,}

	testCases := []struct {
		name          string
		buildStubs    func(categoryService *mocks.MockCategoryService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "SuccessGetCategories",
			buildStubs: func(categoryService *mocks.MockCategoryService) {

				categoryService.EXPECT().
				GetAllCategory().
					Times(1).
					Return(categoryList, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				require.NotNil(t, recorder.Body)
			},
		},
		{
			name: "FailureCategoriesNotFound",
			buildStubs: func(categoryService *mocks.MockCategoryService) {
					categoryService.EXPECT().
					GetAllCategory().
					Times(1).
					Return(nil, app_errors.NewNotFoundError(""))

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailureInternalServerError",
			buildStubs: func(categoryService *mocks.MockCategoryService) {
				categoryService.EXPECT().
				GetAllCategory().
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

			categoryService := mocks.NewMockCategoryService(ctrl)
			tc.buildStubs(categoryService)

			server := NewServer(categoryService)

			recorder := httptest.NewRecorder()
			url := "/categories/api/categories"
			fmt.Println(url)
			request := httptest.NewRequest(http.MethodGet, url, nil)
			server.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

func TestNewCategoryHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	categoryRepository := mocks.NewMockCategoryRepository(ctrl)

	categoryService := services.NewCategoryServiceImpl(categoryRepository)
	categoryHandler := NewCategoryHandler(categoryService)

	assert.Equal(t, categoryHandler.categoryService, categoryService)
}

func requireBodyMatchCategory(t *testing.T, body *bytes.Buffer, requiredResponse models.Category) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var responseReceived models.Category
	err = json.Unmarshal(data, &responseReceived)
	require.NoError(t, err)
	require.Equal(t, responseReceived, requiredResponse)
}