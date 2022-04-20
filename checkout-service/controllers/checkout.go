package controllers

import "github.com/gin-gonic/gin"

// CheckoutController is the controller for the Checkout REST service
type CheckoutController interface {
	// Get an overview of the order
	GetOrderOverview(c *gin.Context)
	// Order Successful Webhook
	OrderCompleteWebhook(c *gin.Context)
	// Health Check Endpoint
	HealthCheck(c *gin.Context)
}
