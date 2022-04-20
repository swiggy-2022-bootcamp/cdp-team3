package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/dto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/models"
	"go.uber.org/zap"
)

const ordersCollection = "Orders"
var validate = validator.New()
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
func GetAllOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetAllOrders Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var ordersList []models.Order;
		params := &dynamodb.ScanInput{
			TableName: aws.String(ordersCollection),
		}

		err := configs.DB.ScanPages(params, func(page *dynamodb.ScanOutput, lastPage bool) bool {
			var orders []models.Order
			err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &orders)
			if err != nil {
				zap.L().Error("\nCould not unmarshal AWS data: err ="+err.Error())
				c.JSON(http.StatusInternalServerError,  dto.ResponseDTO{
					Status: http.StatusInternalServerError, 
					Message: "UnMarshalling of order failed", 
					Data: map[string]interface{}{"data": err.Error()},
				})
				return true
			}
			ordersList = append(ordersList,orders...)
			return true
		})

		if err != nil {
			zap.L().Error(err.Error())
			c.JSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Internal Error", 
				Data: map[string]interface{}{"data": err.Error()},
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
func GetOrdersByStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetOrdersByStatus Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		status := c.Param("status")

		filt := expression.Name("status").Equal(expression.Value(status))

		expr, err := expression.NewBuilder().WithFilter(filt).Build()
		if err != nil {
			zap.L().Error("Error constructing Expression")
			c.JSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Error Fetching data from DB", 
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		input := &dynamodb.ScanInput{
			TableName:                 aws.String(ordersCollection),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			FilterExpression:          expr.Filter(),
		}

		res, err := configs.DB.Scan(input)

		if err != nil {
			zap.L().Error("Error Fetching data from DB")
			c.JSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Error Fetching data from DB", 
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		var orders []models.Order

		if len(res.Items) == 0 {
			zap.L().Info("No orders found with status "+status)
			c.JSON(http.StatusNotFound,  dto.ResponseDTO{
				Status: http.StatusNotFound, 
				Message: "No Orders Found", 
				Data: map[string]interface{}{"data": nil},
			})
			return
		}

		if err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &orders); err != nil {
			zap.L().Error("Error unMarshalling Order"+err.Error())
			c.JSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Error unMarshalling Order", 
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		zap.L().Info("Fetched all orders with status"+status+"successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status: http.StatusOK, 
			Message: "success", 
			Data: map[string]interface{}{"orders": orders},
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
		zap.L().Info("Inside GetOrderById Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		orderId := c.Param("orderId")
		var order models.Order;

		query := &dynamodb.GetItemInput{
			TableName: aws.String(ordersCollection),
			Key: map[string]*dynamodb.AttributeValue{
				"orderId": {
					S: aws.String(orderId),
				},
			},
		}

		result, err := configs.DB.GetItem(query)

		if err != nil {
			zap.L().Error("Failed to get item from database - " + err.Error())
			c.JSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Internal Error", 
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}
	
		if result.Item == nil {
			zap.L().Error("Order for given ID doesn't exists - ")
			c.JSON(http.StatusNotFound,  dto.ResponseDTO{
				Status: http.StatusNotFound, 
				Message: "Order for given ID doesn't exists", 
				Data: map[string]interface{}{"data": nil},
			})
			return
		}
	
		err = dynamodbattribute.UnmarshalMap(result.Item, order)
		if err != nil {
			zap.L().Error("Failed to unmarshal document fetched from DB - " + err.Error())
			c.JSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Failed to unmarshal document fetched from DB", 
				Data: map[string]interface{}{"data": err.Error()},
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
func UpdateStatusById() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside UpdateStatusById Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		orderId := c.Param("orderId")

		var status string
		if err := c.BindJSON(&status); err != nil {
			zap.L().Error("Invalid Request")
			c.JSON(http.StatusBadRequest,  dto.ResponseDTO{
				Status: http.StatusBadRequest, 
				Message: "Invalid Request", 
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&status); validationErr != nil {
			zap.L().Error("Invalid Request")
			c.JSON(http.StatusBadRequest,  dto.ResponseDTO{
				Status: http.StatusBadRequest, 
				Message: "Invalid Request", 
				Data: map[string]interface{}{"data": validationErr.Error()},
			})
			return
		}

		input := &dynamodb.UpdateItemInput{
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				"status": {
					N: aws.String(status),
				},
			},
			Key: map[string]*dynamodb.AttributeValue{
				"orderId": {
					S: aws.String(orderId),
				},
			},
			TableName:        aws.String(ordersCollection),
			UpdateExpression: aws.String("set status = :status"),
			ReturnValues:     aws.String("UPDATED_NEW"),
		}

		response, err := configs.DB.UpdateItem(input)
		if err != nil {
			zap.L().Error("Error while Updating data in dynamoDB"+err.Error())
			c.JSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Error while Updating data in dynamoDB", 
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		zap.L().Info("Updated order status successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status: http.StatusOK, 
			Message: "success", 
			Data: map[string]interface{}{"order": response},
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
func DeleteOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside DeleteOrderById Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		orderId := c.Param("orderId")
	
		response, err := configs.DB.DeleteItem(&dynamodb.DeleteItemInput{
			TableName: aws.String(ordersCollection),
			Key:      map[string]*dynamodb.AttributeValue{
				"orderId": {
					N: aws.String(orderId),
				},
			},
		})
		
		if err != nil {
			zap.L().Error("Error Deleting Order"+err.Error())
			c.JSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Error Deleting Order", 
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		zap.L().Info("Order "+orderId+" Successfully Deleted")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status: http.StatusOK, 
			Message: "success", 
			Data: map[string]interface{}{"order": response},
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
		zap.L().Info("Inside GetOrdersByStatus Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		userId := c.Param("userId")

		filt := expression.Name("customerId").Equal(expression.Value(userId))

		expr, err := expression.NewBuilder().WithFilter(filt).Build()
		if err != nil {
			zap.L().Error("Error constructing Expression")
			c.JSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Error Fetching data from DB", 
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		input := &dynamodb.ScanInput{
			TableName:                 aws.String(ordersCollection),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			FilterExpression:          expr.Filter(),
		}

		res, err := configs.DB.Scan(input)

		if err != nil {
			zap.L().Error("Error Fetching data from DB")
			c.JSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Error Fetching data from DB", 
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		var orders []models.Order

		if len(res.Items) == 0 {
			zap.L().Info("No orders found for customer "+userId)
			c.JSON(http.StatusNotFound,  dto.ResponseDTO{
				Status: http.StatusNotFound, 
				Message: "No Orders Found", 
				Data: map[string]interface{}{"data": nil},
			})
			return
		}

		if err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &orders); err != nil {
			zap.L().Error("Error unMarshalling Order"+err.Error())
			c.JSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Error unMarshalling Order", 
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}

		zap.L().Info("Fetched all orders for customer "+userId+"successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status: http.StatusOK, 
			Message: "success", 
			Data: map[string]interface{}{"orders": orders},
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
		zap.L().Info("Inside GetOrderStatusById Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		orderId := c.Param("orderId")
		var order models.Order;

		query := &dynamodb.GetItemInput{
			TableName: aws.String(ordersCollection),
			Key: map[string]*dynamodb.AttributeValue{
				"orderId": {
					S: aws.String(orderId),
				},
			},
		}

		result, err := configs.DB.GetItem(query)

		if err != nil {
			zap.L().Error("Failed to get item from database - " + err.Error())
			c.JSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Internal Error", 
				Data: map[string]interface{}{"data": err.Error()},
			})
			return
		}
	
		if result.Item == nil {
			zap.L().Error("Order for given ID doesn't exists - ")
			c.JSON(http.StatusNotFound,  dto.ResponseDTO{
				Status: http.StatusNotFound, 
				Message: "Order for given ID doesn't exists", 
				Data: map[string]interface{}{"data": nil},
			})
			return
		}
	
		err = dynamodbattribute.UnmarshalMap(result.Item, order)
		if err != nil {
			zap.L().Error("Failed to unmarshal document fetched from DB - " + err.Error())
			c.JSON(http.StatusInternalServerError,  dto.ResponseDTO{
				Status: http.StatusInternalServerError, 
				Message: "Failed to unmarshal document fetched from DB", 
				Data: map[string]interface{}{"data": err.Error()},
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