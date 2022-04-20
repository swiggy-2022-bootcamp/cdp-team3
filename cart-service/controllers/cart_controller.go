package controllers

import "github.com/gin-gonic/gin"

// CartController is the controller for the cart REST service
type CartController interface {
	// Create a Cart Item
	CreateCartItem(c *gin.Context)
	// Get all Cart Items
	GetCartItems(c *gin.Context)
	// Update a Cart Item
	UpdateCartItem(c *gin.Context)
	// Delete a Cart Item
	DeleteCartItem(c *gin.Context)
	// Empty the Cart
	EmptyCart(c *gin.Context)
	// Health Check Endpoint
	HealthCheck(c *gin.Context)
}
