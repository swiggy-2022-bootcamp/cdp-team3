package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/app/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/middlewares"
)

type AdminRoutes struct {
	adminController controllers.AdminController
}

func NewAdminRoutes(adminController controllers.AdminController) AdminRoutes {
	return AdminRoutes{adminController: adminController}
}

func (ar AdminRoutes) AdminRoutes(router *gin.Engine) {
	router.POST("/admin", ar.adminController.AddAdmin)
	router.Use(middlewares.AuthenticateJWT())
	router.Use(middlewares.OnlyAdmin())
	router.GET("/admin/user", ar.adminController.GetSelf())

}
