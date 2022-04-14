package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/controllers"
)

func OrdersRoute(router *gin.Engine) {

	router.GET("/orders", controllers.GetAllOrders())
	router.GET("/orders/status/:status", controllers.GetOrdersByStatus())
	router.GET("/orders/:orderId", controllers.GetOrderById())
	router.PUT("/orders/:orderId", controllers.UpdateStatusById())
	router.DELETE("/orders/:orderId", controllers.DeleteOrderById())

	router.GET("/orders/user/:userId", controllers.GetOrdersByCustomerId())
	router.POST("/orders/invoice/:orderId", controllers.GenerateInvoiceById())

	//Front Store Route
	router.GET("/orders/:orderId/order_status", controllers.GetOrderStatusById())
}