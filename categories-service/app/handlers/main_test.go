package handlers

import (
	"github.com/gin-gonic/gin"
	"os"
	"github.com/cdp-team3/categories-service/domain/services"
	//"github.com/cdp-team3/categories-service/mocks"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func NewServer(categoryRepository *mocks.MockCategoryRepository) *gin.Engine {
	healthCheckHandler = handlers.NewHealthCheckHandler(categoryRepository)
	categoryService = services. NewCategoryServiceImpl(categoryRepository)
	categoryHandler = handlers.NewCategoryHandler(categoryService)
	

	server := gin.Default()
	router := server.Group("categories/api")

	router.GET("/", healthCheckHandler.HealthCheck)


	return server
}