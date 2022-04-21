package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/swiggy-2022-bootcamp/cdp-team3/products-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/products-service/dto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/products-service/models"
	"go.uber.org/zap"
)

var tableName string = "Products"

// GetAllProducts godoc
// @Summary Fetches and displays all the products
// @Description Fetches and displays all the products
// @Tags Product Service
// @Schemes
// @Produce json
// @Success	200  {object} map[string]interface{}
// @Failure	500  {number} 	500
// @Security Bearer Token
// @Router /products [GET]
func GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside GetAllProducts Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var productsList []models.Product
		params := &dynamodb.ScanInput{
			TableName: aws.String("Products"),
		}

		err := configs.DB.ScanPages(params, func(page *dynamodb.ScanOutput, lastPage bool) bool {
			var products []models.Product
			err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &products)
			if err != nil {
				zap.L().Error("\nCould not unmarshal AWS data: err =" + err.Error())
				c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
					Status:  http.StatusInternalServerError,
					Message: "UnMarshalling of order failed",
					Data:    map[string]interface{}{"data": err.Error()},
				})
				return true
			}
			productsList = append(productsList, products...)
			return true
		})

		if err != nil {
			zap.L().Error(err.Error())
			c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: "Internal Error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		zap.L().Info("Fetched all products successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"products": productsList},
		})

		// c.JSON(http.StatusCreated, gin.H{
		// 	"message": "All Products loaded succesfully",
		// })
	}
}

// AddProduct godoc
// @Summary Adds Product to the DB
// @Description Adds Product to the DB
// @Tags Product Service
// @Schemes
// @Produce json
// @Success	200 {object}  map[string]interface{}
// @Failure	500  {number} 	500
// @Security Bearer Token
// @Router /products [POST]
func AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside AddProduct Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		newProduct := models.Product{}
		json.NewDecoder(c.Request.Body).Decode(&newProduct)
		newProduct.Id = uuid.New().String()

		// fmt.Printf("newProduct: %v\n", newProduct)

		putItem, err := dynamodbattribute.MarshalMap(newProduct)
		if err != nil {
			zap.L().Error(err.Error())
			c.JSON(http.StatusCreated, dto.ResponseDTO{
				Status:  http.StatusBadRequest,
				Data:    map[string]interface{}{"data": err.Error()},
				Message: "Check the Payload Data",
			})
			return
		}

		input := &dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item:      putItem,
		}

		// _, err := configs.DB.PutItemWithContext(c, input)

		if _, err := configs.DB.PutItemWithContext(c, input); err != nil {
			zap.L().Error(err.Error())
			c.JSON(http.StatusCreated, dto.ResponseDTO{
				Status:  http.StatusInternalServerError,
				Data:    map[string]interface{}{"data": err.Error()},
				Message: "Unable to Add Data",
			})
			return
		}

		zap.L().Info("Added product successfully")
		c.JSON(http.StatusCreated, dto.ResponseDTO{
			Status:  http.StatusCreated,
			Data:    map[string]interface{}{"product": newProduct},
			Message: "Product added successfully",
		})

		// c.JSON(http.StatusCreated, gin.H{
		// 	"message": "Product added successfully",
		// })
	}
}

// UpdateProducts godoc
// @Summary Updates the Product details in DB
// @Description Updates the Product details in DB
// @Tags Product Service
// @Schemes
// @Produce json
// @Success	200  {object} models.Product
// @Failure	500  {number} 	500
// @Security Bearer Token
// @Router /products/:productId [PUT]
func UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside UpdateProduct Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		productId := c.Param("productId")

		inpProduct := models.Product{}
		json.NewDecoder(c.Request.Body).Decode(&inpProduct)

		// fmt.Println(productId)
		// fmt.Println(inpProduct)

		input := &dynamodb.UpdateItemInput{
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":model": {
					S: aws.String(inpProduct.Model),
				},
				":quantity": {
					N: aws.String(strconv.Itoa(int(inpProduct.Quantity))),
				},
				":max_quantity": {
					N: aws.String(strconv.Itoa(int(inpProduct.MaxQuantity))),
				},
				":can_place_order": {
					BOOL: &inpProduct.CanPlaceOrder,
				},
				":price": {
					N: aws.String(strconv.FormatFloat(inpProduct.Price, 'E', -1, 64)),
				},
				":manufacturer_id": {
					N: aws.String(strconv.Itoa(int(inpProduct.ManufacturerId))),
				},
				":sku": {
					S: aws.String(inpProduct.Sku),
				},
				":points": {
					N: aws.String(strconv.Itoa(int(inpProduct.Points))),
				},
				":rewards": {
					N: aws.String(strconv.Itoa(int(inpProduct.Rewards))),
				},
				":image": {
					S: aws.String(inpProduct.Image),
				},
				":weight": {
					N: aws.String(strconv.FormatFloat(inpProduct.Weight, 'E', -1, 64)),
				},
				":length": {
					N: aws.String(strconv.FormatFloat(inpProduct.Length, 'E', -1, 64)),
				},
				":width": {
					N: aws.String(strconv.FormatFloat(inpProduct.Width, 'E', -1, 64)),
				},
				":height": {
					N: aws.String(strconv.FormatFloat(inpProduct.Height, 'E', -1, 64)),
				},
				":minimum": {
					N: aws.String(strconv.Itoa(int(inpProduct.Minimun))),
				},
				// ":product_category": {
				// 	L: aws.String(),
				// },
				// ":product_related": {
				// 	N: aws.String(),
				// },
				// ":product_seo_url": {
				// 	N: aws.String(),
				// },
				// ":product_description": {
				// 	N: aws.String(),
				// },
			},
			Key: map[string]*dynamodb.AttributeValue{
				"productId": {
					S: aws.String(productId),
				},
			},
			TableName:        aws.String("Products"),
			UpdateExpression: aws.String("set model = :model, quantity = :quantity, max_quantity = :max_quantity, can_place_order = :can_place_order, price = :price, manufacturer_id = :manufacturer_id, sku = :sku, points = :points, rewards = :rewards, image = :image, weight = :weight, len = :length, width = :width, height = :height, minimum = :minimum"),
			ReturnValues:     aws.String("UPDATED_NEW"),
			// ConditionExpression: aws.String("productId = :productId"),
		}

		response, err := configs.DB.UpdateItem(input)
		if err != nil {
			zap.L().Error("Error while Updating data in dynamoDB" + err.Error())
			c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: "Error while Updating data in dynamoDB",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		zap.L().Info("Updated Product successfully")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"order": response},
		})

		// c.JSON(http.StatusCreated, gin.H{
		// 	"message": "Product updated successfully",
		// })
	}
}

// DeleteProducts godoc
// @Summary Fetches and displays all the products
// @Description Fetches and displays all the products
// @Tags Product Service
// @Schemes
// @Produce json
// @Success	200  {string} Product Deleted Succesfully
// @Failure	500  {number} 	500
// @Security Bearer Token
// @Router /products/:productId [DELETE]
func DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().Info("Inside DeleteOrderById Controller")
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		productId := c.Param("productId")

		response, err := configs.DB.DeleteItem(&dynamodb.DeleteItemInput{
			TableName: aws.String("Products"),
			Key: map[string]*dynamodb.AttributeValue{
				"productId": {
					S: aws.String(productId),
				},
			},
		})

		if err != nil {
			zap.L().Error("Error Deleting Product" + err.Error())
			c.JSON(http.StatusInternalServerError, dto.ResponseDTO{
				Status:  http.StatusInternalServerError,
				Message: "Error Deleting Product",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		zap.L().Info("Order " + productId + " Successfully Deleted")
		c.JSON(http.StatusOK, dto.ResponseDTO{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"order": response},
		})

		// c.JSON(http.StatusCreated, gin.H{
		// 	"message": "Product deleted successfully",
		// })
	}
}

// SearchProduct godoc
// @Summary Searches for Products in DB with given search term and returns list of Products
// @Description Searches for Products in DB with given search term and returns list of Products
// @Tags Product Service
// @Schemes
// @Produce json
// @Success	200  {object} models.Product
// @Failure	500  {number} 	500
// @Security Bearer Token
// @Router /products/search/:search [GET]
func SearchProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		c.JSON(http.StatusCreated, gin.H{
			"message": "Product Search Results",
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
