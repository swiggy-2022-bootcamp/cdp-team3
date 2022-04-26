package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/app/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/middlewares"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/register", controllers.CreateUser)

	router.Use(middlewares.AuthenticateJWT())

	router.GET("/account", controllers.GetUserById)
	router.POST("/account/:id", controllers.UpdateUser)
}
