package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/models"
)

// GetAllOrders godoc
// @Summary Fetch all the orders
// @Description This request will fetch all the orders
// @Tags Orders Service
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {object} 	models.Order
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /orders [GET]
func GetAllOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var order models.Order;
		fmt.Print(order)
		c.JSON(http.StatusCreated, gin.H{
			"message": "All Orders returned",
		})
	}
}

// GetOrdersByStatus godoc
// @Summary Get orders based on order status
// @Description Get all the orders in the application based on the order status for admin to view.
// @Tags Orders Service
// @Schemes
// @Param status path string true "Order Status"
// @Produce json
// @Success	200  {object} models.Order
// @Failure	500  {number} http.StatusInternalServerError
// @Security Bearer Token
// @Router /orders/status/{status} [GET]
func GetOrdersByStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		c.JSON(http.StatusCreated, gin.H{
			"message": "All Orders with requested status returned",
		})
	}
}

// GetOrderById godoc
// @Summary Get order based on order ID.
// @Description Get order details based on Order ID.
// @Tags Orders Service
// @Schemes
// @Param orderId path string true "Order Id"
// @Produce json
// @Success	200  {object} models.Order
// @Failure	500  {number} http.StatusInternalServerError
// @Security Bearer Token
// @Router /orders/{orderId} [GET]
func GetOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		c.JSON(http.StatusCreated, gin.H{
			"message": "Order details for a Order ID returned",
		})
	}
}

// UpdateStatusById godoc
// @Summary Update Order Status by Order ID
// @Description This request will update the order status
// @Tags Orders Service
// @Schemes
// @Accept json
// @Produce json
// @Param orderId path string true "Order Id"
// @Param req body models.OrderStatus true "Order Status"
// @Success	200  {string} 	models.OrderStatus.Status
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /orders/{orderId} [PUT]
func UpdateStatusById() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		c.JSON(http.StatusCreated, gin.H{
			"message": "Order Status Updated Successfully",
		})
	}
}

// DeleteOrderById godoc
// @Summary Delete Order by Order ID
// @Description This request will delete a particular order
// @Tags Orders Service
// @Schemes
// @Produce json
// @Param orderId path string true "Order Id"
// @Success	200  {string} 	map[string]interface{}
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /orders/{orderId} [DELETE]
func DeleteOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		c.JSON(http.StatusCreated, gin.H{
			"message": "Order Status Updated Successfully",
		})
	}
}

// GetOrdersByCustomerId godoc
// @Summary Get orders of a customer based on customer ID.
// @Description Get order details of a customer based on Customer ID.
// @Tags Orders Service
// @Schemes
// @Param userId path string true "User Id"
// @Produce json
// @Success	200  {object} models.Order
// @Failure	500  {number} http.StatusInternalServerError
// @Security Bearer Token
// @Router /orders/user/{userId} [GET]
func GetOrdersByCustomerId() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		c.JSON(http.StatusCreated, gin.H{
			"message": "All Orders of a particular customer returned",
		})
	}
}



// GenerateInvoiceById godoc
// @Summary Generate invoice for a particular Order by Order ID
// @Description This request will generate an invoice for order
// @Tags Orders Service
// @Schemes
// @Produce json
// @Param orderId path string true "Order Id"
// @Success	201  {string} 	models.OrderInvoice
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /orders/invoice/{orderId} [POST]
func GenerateInvoiceById() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		c.JSON(http.StatusCreated, gin.H{
			"message": "Order Invoice Created Successfully",
		})
	}
}


// GetOrderStatusById-Front Store godoc
// @Summary Get Order Status by Order ID
// @Description This request will fetch details of order status 
// @Tags Orders Service
// @Schemes
// @Produce json
// @Param orderId path string true "Order Id"
// @Success	200  {object} 	models.OrderStatus
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /orders/{orderId}/order_status [GET]
func GetOrderStatusById() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		c.JSON(http.StatusCreated, gin.H{
			"message": "Order Status received Successfully",
		})
	}
}


// HealthCheck godoc
// @Summary To check if the service is running or not.
// @Description This request will return 200 OK if server is up..
// @Tags Health
// @Schemes
// @Accept json
// @Produce json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router / [GET]
func HealthCheck() gin.HandlerFunc {
	//Check to be added for database.
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Order Service is running"})
	}
}