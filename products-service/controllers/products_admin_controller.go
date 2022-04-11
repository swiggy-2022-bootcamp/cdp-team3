package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAllProducts godoc
// @Summary Fetches and displays all the products
// @Description Fetches and displays all the products
// @Tags Product Service
// @Schemes
// @Produce json
// @Success	200  {object} models.Product
// @Failure	500  {number} 	500
// @Security Bearer Token
// @Router /products [GET]
func GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		c.JSON(http.StatusCreated, gin.H{
			"message": "All Products loaded succesfully",
		})
	}
}

// AddProduct godoc
// @Summary Adds Product to the DB
// @Description Adds Product to the DB
// @Tags Product Service
// @Schemes
// @Produce json
// @Success	200  {string} Product Added Succesfully
// @Failure	500  {number} 	500
// @Security Bearer Token
// @Router /products [POST]
func AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		c.JSON(http.StatusCreated, gin.H{
			"message": "Product added successfully",
		})
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
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		c.JSON(http.StatusCreated, gin.H{
			"message": "Product updated successfully",
		})
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
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		c.JSON(http.StatusCreated, gin.H{
			"message": "Product deleted successfully",
		})
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
