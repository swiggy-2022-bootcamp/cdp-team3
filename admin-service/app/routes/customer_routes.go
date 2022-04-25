package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/app/controllers"
)

type CustomerRoutes struct {
	customerController controllers.CustomerController
}

func NewCustomerRoutes(customerController controllers.CustomerController) CustomerRoutes {
	return CustomerRoutes{customerController: customerController}
}

func (cr CustomerRoutes) CustomerRoutes(router *gin.Engine) {

	// router.Use(middlewares.AuthenticateJWT())

	customerRoutes := router.Group("/customers")
	{
		// customerRoutes.Use(middlewares.OnlyAdmin())
		customerRoutes.GET("/", cr.customerController.GetAllCustomers())
		customerRoutes.POST("/", cr.customerController.AddCustomer)
		customerRoutes.GET("/:customerId", cr.customerController.GetCustomerById())
		customerRoutes.GET("/email/:emailId", cr.customerController.GetCustomerByEmail())
		customerRoutes.DELETE("/:customerId", cr.customerController.DeleteCustomerById())
		customerRoutes.PUT("/:customerId", cr.customerController.UpdateCustomerById())

	}

}
