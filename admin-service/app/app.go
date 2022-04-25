package app

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/app/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/app/routes"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/configs"
	_ "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/domain/services"
	transactionGrpc "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/grpc/admin"
	rewardGrpc "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/grpc/reward"

	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/utils"
	"go.uber.org/zap"
)

var (
	router             *gin.Engine
	adminRepository    repository.AdminRepository
	adminService       services.AdminService
	adminController    controllers.AdminController
	adminRoutes        routes.AdminRoutes
	adminDB            *dynamodb.DynamoDB
	customerRepository repository.CustomerRepository
	customerService    services.CustomerService
	customerController controllers.CustomerController
	customerRoutes     routes.CustomerRoutes
	customerDB         *dynamodb.DynamoDB
)

func Start() {

	//Initialize Logger
	log := utils.InitializeLogger()

	zap.ReplaceGlobals(log)
	defer log.Sync()
	log.Info("Admin Service Started")

	//Initialize DB
	customerDB = configs.ConnectDB()
	adminDB = configs.ConnectDB()
	configs.CreateCustomerTable(customerDB)
	configs.CreateAdminTable(adminDB)

	adminRepository = repository.NewAdminRepositoryImpl(adminDB)
	adminService = services.NewAdminServiceImpl(adminRepository)
	adminController = controllers.NewAdminController(adminService)
	adminRoutes = routes.NewAdminRoutes(adminController)
	customerRepository = repository.NewCustomerRepositoryImpl(customerDB)
	customerService = services.NewCustomerServiceImpl(customerRepository)
	customerController = controllers.NewCustomerController(customerService)
	customerRoutes = routes.NewCustomerRoutes(customerController)
	go transactionGrpc.InitialiseTransactionsServer()
	go rewardGrpc.InitialiseRewardsServer()

	router := StartRestServer()
	router.Run(":" + configs.EnvPORT())
}

func StartRestServer() *gin.Engine {
	router = gin.Default()

	router.GET("/", controllers.HealthCheck())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	customerRoutes.CustomerRoutes(router)
	adminRoutes.AdminRoutes(router)
	return router
}
