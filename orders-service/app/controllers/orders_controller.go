package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/domain/services"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/dto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/models"
	"go.uber.org/zap"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return OrderController{orderService: orderService}
}

// GetAllOrders godoc
// @Summary Fetch all the orders
// @Description This request will fetch all the orders
// @Tags Orders Service
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {array} 	models.Order
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /orders [GET]
func (oc OrderController) GetAllOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetAllOrders Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		ordersList,err := oc.orderService.GetAllOrders()

		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Internal Error", 
				Data: map[string]interface{}{"data": err.Message},
			})
			return
		}

		zap.L().Info("Fetched all orders successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status: http.StatusOK, 
			Message: "success", 
			Data: map[string]interface{}{"orders": ordersList},
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
func (oc OrderController) GetOrdersByStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetOrdersByStatus Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		status := c.Param("status")

		ordersList,err := oc.orderService.GetOrdersByStatus(status)

		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Internal Error", 
				Data: map[string]interface{}{"data": err.Message},
			})
			return
		}

		zap.L().Info("Fetched all orders with status"+status+"successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status: http.StatusOK, 
			Message: "success", 
			Data: map[string]interface{}{"orders": ordersList},
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
func (oc OrderController) GetOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetOrderById Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		orderId := c.Param("orderId")
		order, err := oc.orderService.GetOrderById(orderId)

		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Internal Error", 
				Data: map[string]interface{}{"data": err.Message},
			})
			return
		}

		zap.L().Info("Fetched order successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status: http.StatusOK, 
			Message: "success", 
			Data: map[string]interface{}{"order": order},
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
func  (oc OrderController) UpdateStatusById() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside UpdateStatusById Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		orderId := c.Param("orderId")

		var orderStatus models.OrderStatus
		if err := c.BindJSON(&orderStatus); err != nil {
			zap.L().Error("Invalid Request")
			c.JSON(http.StatusBadRequest,  dto.ResponseDTO{
				Status: http.StatusBadRequest, 
				Message: "Invalid Request", 
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		updatedOrder, err := oc.orderService.UpdateStatusById(orderId, orderStatus)
		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Internal Error", 
				Data: map[string]interface{}{"data": err.Message},
			})
			return
		}

		zap.L().Info("Updated order status successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status: http.StatusOK, 
			Message: "success", 
			Data: map[string]interface{}{"order": updatedOrder},
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
// @Success	200  {string} 	Deleted Successfully
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	500  {number} 	http.StatusInternalServerError
// @Security Bearer Token
// @Router /orders/{orderId} [DELETE]
func  (oc OrderController) DeleteOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside DeleteOrderById Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		orderId := c.Param("orderId")
	
		deletedOrder, err := oc.orderService.DeleteOrderById(orderId)
		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Internal Error", 
				Data: map[string]interface{}{"data": err.Message},
			})
			return
		}


		zap.L().Info("Order "+orderId+" Successfully Deleted")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status: http.StatusOK, 
			Message: "success", 
			Data: map[string]interface{}{"order": deletedOrder},
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
func  (oc OrderController) GetOrdersByCustomerId() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetOrdersByStatus Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		userId := c.Param("userId")

		ordersList,err := oc.orderService.GetOrdersByCustomerId(userId)

		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Internal Error", 
				Data: map[string]interface{}{"data": err.Message},
			})
			return
		}

		zap.L().Info("Fetched all orders for customer "+userId+"successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status: http.StatusOK, 
			Message: "success", 
			Data: map[string]interface{}{"orders": ordersList},
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
func (oc OrderController) GenerateInvoiceById() gin.HandlerFunc {
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
func (oc OrderController) GetOrderStatusById() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetOrderStatusById Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		orderId := c.Param("orderId")
		order, err := oc.orderService.GetOrderById(orderId)

		if err != nil {
			zap.L().Error(err.Message)
			c.AbortWithStatusJSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Internal Error", 
				Data: map[string]interface{}{"data": err.Message},
			})
			return
		}

		zap.L().Info("Fetched order status successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status: http.StatusOK, 
			Message: "success", 
			Data: map[string]interface{}{"status": order.Status},
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
	
	//Ping DB
	_, err := configs.DB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		zap.L().Error("Database connection is down.")
		return func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, dto.HealthCheckResponse{Server: "Server is up", Database: "Database is down"})
		}
	}
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, dto.HealthCheckResponse{Server: "Server is up", Database: "Database is up"})
	}
}