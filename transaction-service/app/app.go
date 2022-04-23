package app

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/app/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/app/routes"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/utils"
	"go.uber.org/zap"
)

var (
	transactionRepository repository.TransactionRepository
	transactionController controllers.TransactionController
	transactionRoutes routes.TransactionRoutes
)

func Start(){
	//Initialize Logger
	log := utils.InitializeLogger()

  zap.ReplaceGlobals(log)
  defer log.Sync()
  log.Info("Transaction Service Started")
	
	//Initialize DB
	transactionDB := configs.ConnectDB()
	configs.CreateTable(transactionDB)

	transactionRepository = repository.NewTransactionRepositoryImpl(transactionDB)
	transactionController = controllers.NewTransactionController(transactionRepository)
	transactionRoutes = routes.NewTransactionRoutes(transactionController)

	router := gin.Default()

	router.GET("/", controllers.HealthCheck())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	transactionRoutes.TransactionRoutes(router)
	router.Run(":"+configs.EnvPORT())
}