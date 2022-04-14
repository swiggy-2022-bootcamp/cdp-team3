package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-ipp/checkout-service/dto/responses"
)

// @Summary      Get an overview of the order
// @Description  Get an overview of the order from the database.
// @Tags         Checkout API
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.MessageResponse  "Order Overview Data"
// @Failure      400  {object}  errors.HTTPErrorDTO
// @Failure      404  {object}  errors.HTTPErrorDTO
// @Failure      500  {object}  nil
// @Router       /confirm [post]
func GetOrderOverview(c *gin.Context) {
	c.JSON(200, responses.MessageResponse{Message: "Here's your order."})
}

// @Summary      Order Successful Webhook
// @Description  Webhook hit when Order is successful to clear Cart and Unset Session Data.
// @Tags         Checkout API
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.MessageResponse  "Cart Cleared message DTO."
// @Failure      400  {object}  errors.HTTPErrorDTO
// @Failure      404  {object}  errors.HTTPErrorDTO
// @Failure      500  {object}  nil
// @Router       /confirm/success [put]
func OrderCompleteWebhook(c *gin.Context) {
	c.JSON(200, responses.MessageResponse{Message: "Order Complete."})
}

// @Summary      Health Check Endpoint
// @Description  Health Check Endpoint
// @Tags         Health Check
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.MessageResponse  "Health Check Message."
// @Failure      400  {object}  errors.HTTPErrorDTO
// @Failure      404  {object}  errors.HTTPErrorDTO
// @Failure      500  {object}  nil
// @Router       /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(200, responses.MessageResponse{Message: "up"})
}
