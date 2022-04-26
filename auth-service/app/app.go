package app

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/app/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/app/routes"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/configs"
	_ "github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/domain/services"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/grpc/auth"
)

var (
	userRepository repository.UserRepository
	authService    services.AuthService
	authController controllers.AuthController
	authRoutes     routes.AuthRoutes
	healthRoutes   routes.HealthRoutes
)

func Start() {
	userDB := configs.ConnectDB()
	userRepository = repository.NewUserRepository(userDB)
	authService = services.NewAuthServiceImpl(userRepository)
	authController = *controllers.NewAuthController(authService)
	authRoutes = routes.NewAuthRouter(authController)

	go auth.InitialiseAuthServer()

	router := InitialiseRestServer()
	router.Run(configs.EnvAuthHost() + ":" + configs.EnvPORT())
}

func InitialiseRestServer() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRoutes.AuthRoute(router)
	healthRoutes.HealthRoute(router)
	return router
}
