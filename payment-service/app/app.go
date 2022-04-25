package app

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/app/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/app/routes"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/configs"
	_ "github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/domain/services"
)

var (
	paymentRepository repository.PaymentRepository
	paymentService    services.PaymentService
	paymentController controllers.PaymentController
	paymentRoutes     routes.PaymentRoutes
	healthRoutes      routes.HealthRoutes
)

func Start() {
	paymentDB := configs.ConnectDB()
	paymentRepository = repository.NewPaymentRepository(paymentDB)
	paymentService = services.NewPaymentServiceImpl(paymentRepository)
	paymentController = *controllers.NewPaymentController(paymentService)
	paymentRoutes = routes.NewPaymentRouter(paymentController)

	router := InitialiseRestServer()
	router.Run(":" + configs.EnvPORT())
}

func InitialiseRestServer() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	paymentRoutes.PaymentRoute(router)
	healthRoutes.HealthRoute(router)
	return router
}
