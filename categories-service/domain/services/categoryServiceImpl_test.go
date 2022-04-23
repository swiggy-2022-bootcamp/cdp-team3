package services

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
     app_erros "github.com/cdp-team3/categories-service/app-errors"
	"github.com/cdp-team3/categories-service/domain/models"
	 "github.com/cdp-team3/categories-service/mocks"
	 "testing"
)
func TestCategoryServiceImpl_GetCategory(t *testing.T) {
	gin.SetMode(gin.TestMode)
	category_id := "bb912edc-50d9-42d7-b7a1-9ce66d459thj"
    categoryDesc := models.CategoryDescription{
	
			Name    :"TestCategoryDescName",        
		
			Description     : "TestCategoryDescDescription",
			MetaDescription :"TestCategoryDescMetaDescription",
			MetaKeyword     :"TestCategoryDescMetaKeyword",
			MetaTitle       :"TestCategoryDescMetaTitle",
	
	}
	successCategory := &models.Category{
		
		CategoryId: "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
		CategoryDescription : []models.CategoryDescription{categoryDesc},
		
		
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockCategoryRepository)
		checkResponse func(t *testing.T, category interface{}, err interface{})
	}{
		{
			name: "SuccessCategoryFound",
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				    GetCategoryFromDB(category_id).
					Times(1).
					Return(successCategory, nil)
			},
			checkResponse: func(t *testing.T, category interface{}, err interface{}) {
				
				assert.NotNil(t,category)
				assert.Nil(t,err)
			},
		},
		{
			name: "FailureCategoriesNotFound",
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				GetCategoryFromDB(category_id).
					Times(1).
					Return(nil, app_erros.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, category interface{}, err interface{}) {
	
			    assert.Nil(t,category)
				assert.NotNil(t,err)
		
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				    GetCategoryFromDB(category_id).
					Times(1).
					Return(nil, app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, category interface{}, err interface{}) {
			
				assert.Nil(t,category)
				assert.NotNil(t,err)
		
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			categoryRepository := mocks.NewMockCategoryRepository(ctrl)
			tc.buildStubs(categoryRepository)

			categoryServiceImpl := NewCategoryServiceImpl(categoryRepository)
			category, err := categoryServiceImpl.GetCategory(category_id)
			tc.checkResponse(t,category, err)
		})
	}
}
func TestCategoryServiceImpl_GetAllCategory(t *testing.T) {
	gin.SetMode(gin.TestMode)
	categoryDesc1 := models.CategoryDescription{
	
		Name : "testname1",
					Description  : "testdesc1",
					MetaDescription: "testmetadesc1",
					MetaKeyword : "testmetakey1",
					MetaTitle: "testmetatitle1",
}
categoryDesc2 := models.CategoryDescription{
	
	Name : "testname2",
				Description  : "testdesc2",
				MetaDescription: "testmetadesc2",
				MetaKeyword : "testmetakey2",
				MetaTitle: "testmetatitle2",
}
	successCategories := []models.Category{
		{
			CategoryId: "423fec6b-8a0c-4f99-8b2b-6eeef7605a37",
			CategoryDescription:[]models.CategoryDescription{categoryDesc1},
		
		},
		{
			CategoryId: "523fec6b-8a0c-4f99-8b2b-6eeef7605a37",
			CategoryDescription:[]models.CategoryDescription{categoryDesc2},
		},
	}


	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockCategoryRepository)
		checkResponse func(t *testing.T,res []models.Category, err interface{})
	}{
		{
			name: "SuccessCategoriesFound",
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				    FindAllCategoryFromDB().
					Times(1).
					Return(successCategories, nil)
			},
			checkResponse: func(t *testing.T,res []models.Category, err interface{}) {
				
				assert.NotNil(t,res)
				assert.Nil(t,err)
			},
		},
		{
			name: "FailureCategoryNotFound",
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				FindAllCategoryFromDB().
					Times(1).
					Return(nil, app_erros.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T,res []models.Category, err interface{}) {
	
			assert.Nil(t,res)
				assert.NotNil(t,err)
		
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				FindAllCategoryFromDB().
					Times(1).
					Return(nil, app_erros.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T,res []models.Category, err interface{}) {
			
				assert.Nil(t,res)
				assert.NotNil(t,err)
		
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			categoryRepository := mocks.NewMockCategoryRepository(ctrl)
			tc.buildStubs(categoryRepository)

			categoryServiceImpl := NewCategoryServiceImpl(categoryRepository)
			res, err := categoryServiceImpl.GetAllCategory()
			tc.checkResponse(t,  res, err)
		})
	}
}
func TestCategoryServiceImpl_AddCategory(t *testing.T) {
	gin.SetMode(gin.TestMode)
	categoryDesc := models.CategoryDescription{
	
		Name    :"TestCategoryDescName",        
	
		Description     : "TestCategoryDescDescription",
		MetaDescription :"TestCategoryDescMetaDescription",
		MetaKeyword     :"TestCategoryDescMetaKeyword",
		MetaTitle       :"TestCategoryDescMetaTitle",

}
category := &models.Category{
	
	CategoryId: "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
	CategoryDescription : []models.CategoryDescription{categoryDesc},
	
	
}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockCategoryRepository)
		checkResponse func(t *testing.T,err interface{})
	}{
	
		{
			name: "SuccessAddCategory",
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				AddCategoryToDB(category).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T,err interface{}) {
				assert.Nil(t,err)
			},
		},
		{
			name: "FailureAddCategory",
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				    AddCategoryToDB(category).
					Times(1).
					Return(app_erros.NewUnexpectedError(""))

			},
			checkResponse: func(t *testing.T,err interface{}) {
				assert.NotNil(t,err)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				    AddCategoryToDB(category).
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

			categoryRepository := mocks.NewMockCategoryRepository(ctrl)
			tc.buildStubs(categoryRepository)

			categoryServiceImpl := NewCategoryServiceImpl(categoryRepository)
			err := categoryServiceImpl.AddCategory(category)
			tc.checkResponse(t, err)
		})
	}
}
func TestCategoryServiceImpl_DeleteCategoryByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	category_id := "bb912edc-50d9-42d7-b7a1-9ce66d459thj"
	

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockCategoryRepository)
		checkResponse func(t *testing.T, result bool , err interface{})
	}{
		{
			name: "SuccessCategoryDeleted",
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				    DeleteCategoryByIDFromDB(category_id).
					Times(1).
					Return(true, nil)
			},
			checkResponse: func(t *testing.T, result bool, err interface{}) {
				assert.Equal(t, true, result)
				assert.Nil(t,err)
			
			},
		},
		{
			name: "FailureCategoryNotDeleted",
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				    DeleteCategoryByIDFromDB(category_id).
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
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				    DeleteCategoryByIDFromDB(category_id).
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

			categoryRepository := mocks.NewMockCategoryRepository(ctrl)
			tc.buildStubs(categoryRepository)

			categoryServiceImpl := NewCategoryServiceImpl(categoryRepository)
			value, err := categoryServiceImpl.DeleteCategoryByID(category_id)
			tc.checkResponse(t,value, err)
		})
	}
}
func TestCategoryServiceImpl_DeleteCategories(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	
    categoriesId:= []string{"bb912edc-50d9-42d7-b7a1-9ce66d459thj","bb912edc-50d9-42d7-b7a1-9ce66d4522hj"}
	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockCategoryRepository)
		checkResponse func(t *testing.T, result bool , err interface{})
	}{
		{
			name: "SuccessCategoriesDeleted",
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				DeleteCategoriesFromDB(categoriesId).
					Times(1).
					Return(true, nil)
			},
			checkResponse: func(t *testing.T, result bool, err interface{}) {
				assert.Equal(t, true, result)
				assert.Nil(t,err)
			
			},
		},
		{
			name: "FailureCategoriesNotDeleted",
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				DeleteCategoriesFromDB(categoriesId).
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
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				DeleteCategoriesFromDB(categoriesId).
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

			categoryRepository := mocks.NewMockCategoryRepository(ctrl)
			tc.buildStubs(categoryRepository)

			categoryServiceImpl := NewCategoryServiceImpl(categoryRepository)
			value, err := categoryServiceImpl.DeleteCategories(categoriesId)
			tc.checkResponse(t,value, err)
		})
	}
}
func TestCategoryServiceImpl_UpdateCategoriesById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	category_id := "bb912edc-50d9-42d7-b7a1-9ce66d459thj"
	categoryDesc := models.CategoryDescription{
	
		Name    :"TestCategoryDescName",        
	
		Description     : "TestCategoryDescDescription",
		MetaDescription :"TestCategoryDescMetaDescription",
		MetaKeyword     :"TestCategoryDescMetaKeyword",
		MetaTitle       :"TestCategoryDescMetaTitle",

    }
     category := &models.Category{
	
	CategoryId: "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
	CategoryDescription : []models.CategoryDescription{categoryDesc},
	
	
    }
	

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockCategoryRepository)
		checkResponse func(t *testing.T, result bool , err interface{})
	}{
		{
			name: "SuccessCategoryUpdated",
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				    UpdateCategoryByIdFromDB(category_id,category).
					Times(1).
					Return(true, nil)
			},
			checkResponse: func(t *testing.T, result bool, err interface{}) {
				assert.Equal(t, true, result)
				assert.Nil(t,err)
				
			},
		},
		{
			name: "FailureCategoryNotUpdated",
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				UpdateCategoryByIdFromDB(category_id,category).
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
			buildStubs: func(repository *mocks.MockCategoryRepository) {
				repository.EXPECT().
				UpdateCategoryByIdFromDB(category_id,category).
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

		    categoryRepository := mocks.NewMockCategoryRepository(ctrl)
			tc.buildStubs(categoryRepository)

			categoryServiceImpl := NewCategoryServiceImpl(categoryRepository)
			value, err := categoryServiceImpl.UpdateCategoryByID(category_id,category)
			tc.checkResponse(t,value, err)
		})
	}
}