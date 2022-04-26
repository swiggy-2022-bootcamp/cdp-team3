package handlers

import (
	"github.com/gin-gonic/gin"
	"os"
	"github.com/cdp-team3/categories-service/mocks"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

// Creating test server with CategoriesService mock.
func NewServer(categoryService *mocks.MockCategoryService) *gin.Engine {
	categoriesHandler := NewCategoryHandler(categoryService)

	server := gin.Default()
	router := server.Group("categories/api")
     
	router.POST("/categories", categoriesHandler.AddCategory())
	router.GET("/categories", categoriesHandler.GetAllCategory())
	router.GET("/categories/:category_id",  categoriesHandler.GetCategory())
	router.DELETE("/categories/:category_id",  categoriesHandler.DeleteCategory())
    router.DELETE("/categories/",  categoriesHandler.DeleteCategories())
	router.PUT("/categories/:category_id",  categoriesHandler.UpdateCategory())
	
	return server
}