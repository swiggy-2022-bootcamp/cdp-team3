package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/app/controllers"
)

type OrderRoutes struct {
	ordersController controllers.OrderController
}
func NewOrderRoutes(ordersController controllers.OrderController) OrderRoutes {
	return OrderRoutes{ordersController: ordersController}
}

func (or OrderRoutes)OrdersRoute(router *gin.Engine) {

	router.GET("/orders", or.ordersController.GetAllOrders())
	router.GET("/orders/status/:status", or.ordersController.GetOrdersByStatus())
	router.GET("/orders/:orderId", or.ordersController.GetOrderById())
	router.PUT("/orders/:orderId", or.ordersController.UpdateStatusById())
	router.DELETE("/orders/:orderId", or.ordersController.DeleteOrderById())

	router.GET("/orders/user/:userId", or.ordersController.GetOrdersByCustomerId())
	router.POST("/orders/invoice/:orderId", or.ordersController.GenerateInvoiceById())

	//Front Store Route
	router.GET("/orders/:orderId/order_status", or.ordersController.GetOrderStatusById())
}