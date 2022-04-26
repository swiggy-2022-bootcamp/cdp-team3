package app

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/app/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/app/routes"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/configs"
	_ "github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/domain/services"

	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/utils"
	"go.uber.org/zap"
)

var (
	server           *gin.Engine
	rewardRepository repository.RewardRepository
	rewardService    services.RewardService
	rewardController controllers.RewardController
	rewardRoutes     routes.RewardRoutes
	rewardDB         *dynamodb.DynamoDB
)

func Start() {

	//Initialize Logger
	log := utils.InitializeLogger()

	zap.ReplaceGlobals(log)
	defer log.Sync()
	log.Info("Rewards Service Started")

	//Initialize DB
	rewardDB := configs.ConnectDB()
	configs.CreateTable(rewardDB)

	rewardRepository = repository.NewRewardRepositoryImpl(rewardDB)
	rewardService = services.NewRewardServiceImpl(rewardRepository)
	rewardController = controllers.NewRewardController(rewardService)
	rewardRoutes = routes.NewRewardRoutes(rewardController)

	router := StartRestServer()
	router.Run(configs.EnvRewardsHost() + ":" + configs.EnvRewardServicePORT())
}

func StartRestServer() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.HealthCheck())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rewardRoutes.RewardsRoute(router)
	return router
}
