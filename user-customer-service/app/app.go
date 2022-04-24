package app

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/app/routes"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/configs"
	_ "github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/utils"
	"go.uber.org/zap"
)

var (
	server *gin.Engine
	userDB *dynamodb.DynamoDB
)

func Start() {

	//Initialize Logger
	log := utils.InitializeLogger()

	zap.ReplaceGlobals(log)
	defer log.Sync()
	log.Info("Users Service Started")

	//Initialize DB
	userDB := configs.ConnectDB()
	configs.CreateTable(userDB)

	router := StartRestServer()
	router.Run(":" + configs.EnvPORT())
}

func StartRestServer() *gin.Engine {
	router := gin.Default()

	// router.GET("/", controllers.HealthCheck())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UserRoutes(router)
	return router
}
