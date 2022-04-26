package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/app/controllers"
)

type AuthRoutes struct {
	authController controllers.AuthController
}

func NewAuthRouter(authController controllers.AuthController) AuthRoutes {
	return AuthRoutes{authController: authController}
}

func (ar AuthRoutes) AuthRoute(router *gin.Engine) {
	public := router.Group("/auth")
	public.POST("/login", ar.authController.Login())
	public.POST("/logout", ar.authController.Logout())
	public.POST("/verify-token", ar.authController.VerifyToken())
}
