package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/controllers"
)

func AdminRoute(router *gin.Engine) {

	// router.Use(controllers.CheckAuthorized("admin"))

	router.GET("/user", controllers.GetSelf)
	router.POST("/customers", controllers.CreateCustomer)
	router.GET("/customers/:customerId", controllers.GetCustomerById)
	router.GET("/customers/email/:customerEmail", controllers.GetCustomerByEmail)
	router.PUT("/customers/:customerId", controllers.UpdateCustomer)
	router.DELETE("/customers/:customerId", controllers.DeleteCustomer)

}
