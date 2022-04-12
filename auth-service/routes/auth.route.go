package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/controllers"
)

func AuthRoutes(router *gin.Engine) {
	public := router.Group("/auth")
	public.POST("/login", controllers.Login())
	public.POST("/logout", controllers.Logout())
	public.POST("/verify-token", controllers.VerifyToken())
}
