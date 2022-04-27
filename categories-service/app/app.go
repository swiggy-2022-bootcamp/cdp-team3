package app

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/gin-gonic/gin"
	"io"
	"os"
	"github.com/cdp-team3/categories-service/app/handlers"
	"github.com/cdp-team3/categories-service/app/routes"
    "github.com/cdp-team3/categories-service/config"
	"github.com/cdp-team3/categories-service/domain/repository"
	"github.com/cdp-team3/categories-service/domain/services"
	"github.com/cdp-team3/categories-service/utils/logger"
	
)

const categoryCollection = "categoryCollection_Team3"
 var (
	server                *gin.Engine
	categoryRepository repository.CategoryRepository
	categoryService    services.CategoryService
	categoryHandler    handlers.CategoryHandler
	categoriesRoutes     routes.CategoriesRoutes
	categoriesDB         *dynamodb.DynamoDB
	healthCheckHandler    handlers.HealthCheckHandler
)


func Start() {

//Variable initializations for DynamoDB
categoriesDB = config.ConnectDB()
config.CreateTable(categoriesDB)

// //Variable initializations to be used as dependency injections
 categoryRepository = repository.NewCategoryRepositoryImpl(categoriesDB)
 categoryService = services. NewCategoryServiceImpl(categoryRepository)
 categoryHandler = handlers.NewCategoryHandler(categoryService)
 healthCheckHandler = handlers.NewHealthCheckHandler(categoryRepository)
 categoriesRoutes = routes.NewCategoryRoutes(categoryHandler, healthCheckHandler)

//Opening file for log collection
file, err := os.OpenFile("categories-service-server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err == nil {
	gin.DefaultWriter = io.MultiWriter(file)
}

 server = gin.New()
server.Use(logger.UseLogger(logger.DefaultLoggerFormatter), gin.Recovery())
 router := server.Group("categories/api")
 categoriesRoutes.InitRoutes(router)

//Starting server on port 3002
err = server.Run(config.EnvCategoryHost()+":"+config.EnvCategoriesPort())
if err != nil {
	logger.Error(err.Error() + " - Failed to start server")
} else {
	logger.Info("Categories Server started successfully.")
}
}
