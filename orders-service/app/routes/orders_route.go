package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/app/controllers"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/middlewares"
)

type OrderRoutes struct {
	ordersController controllers.OrderController
}
func NewOrderRoutes(ordersController controllers.OrderController) OrderRoutes {
	return OrderRoutes{ordersController: ordersController}
}

func (or OrderRoutes)OrdersRoute(router *gin.Engine) {

	router.Use(middlewares.AuthenticateJWT())

	adminRoutes := router.Group("/orders")
	{
		adminRoutes.Use(middlewares.OnlyAdmin())
		adminRoutes.GET("/", or.ordersController.GetAllOrders())
		adminRoutes.GET("/status/:status", or.ordersController.GetOrdersByStatus())
		adminRoutes.PUT("/:orderId", or.ordersController.UpdateStatusById())
		adminRoutes.DELETE("/:orderId", or.ordersController.DeleteOrderById())
	
		adminRoutes.GET("/user/:userId", or.ordersController.GetOrdersByCustomerId())
		adminRoutes.POST("/invoice/:orderId", or.ordersController.GenerateInvoiceById())
	}

	//Should be available to front store only if the order is placed by the customer
	router.GET("/:orderId", or.ordersController.GetOrderById())

	//Front Store Route
	router.GET("/orders/:orderId/order_status", or.ordersController.GetOrderStatusById())
}