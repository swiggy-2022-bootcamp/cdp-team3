package controllers

import (
	"github.com/gin-gonic/gin"
)

// @Summary      Get an overview of the order
// @Description  Get an overview of the order from the database.
// @Tags         Checkout API
// @Accept       json
// @Produce      json
// @Success      200  {object}  interface{}  "Order Overview Data"
// @Failure      400  {object}  errors.HTTPErrorDTO
// @Failure      404  {object}  errors.HTTPErrorDTO
// @Failure      500  {object}  nil
// @Router       /confirm [post]
func GetOrderOverview(c *gin.Context) {
	c.JSON(201, gin.H{"message": "Here's your order."})
}

// @Summary      Order Successful Webhook
// @Description  Webhook hit when Order is successful to clear Cart and Unset Session Data.
// @Tags         Checkout API
// @Accept       json
// @Produce      json
// @Success      200  {object}  interface{}  "Cart Cleared message DTO."
// @Failure      400  {object}  errors.HTTPErrorDTO
// @Failure      404  {object}  errors.HTTPErrorDTO
// @Failure      500  {object}  nil
// @Router       /confirm/success [put]
func OrderCompleteWebhook(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Order Complete."})
}
