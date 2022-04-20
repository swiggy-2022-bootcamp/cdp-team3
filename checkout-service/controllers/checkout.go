package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/swiggy-ipp/checkout-service/dto/responses"
	"github.com/swiggy-ipp/checkout-service/grpcs"
	"github.com/swiggy-ipp/checkout-service/grpcs/cart_checkout"
	"github.com/swiggy-ipp/checkout-service/grpcs/shipping_checkout"
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
	// Set up context
	ctx := c.Request.Context()
	// Make GRPC Call
	out, err := grpcs.ShippingCheckoutGRPCClient.GetShippingAddress(ctx, &shipping_checkout.ShippingAddressRequest{})
	if err != nil {
		log.Error("Error emptying cart: ", err)
		c.JSON(500, err)
	} else {
		c.JSON(200, out)
	}
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
// @Router       /confirm/success [post]
func OrderCompleteWebhook(c *gin.Context) {
	// Set up context
	ctx := c.Request.Context()
	// Make GRPC Call
	out, err := grpcs.CartCheckoutGRPCClient.EmptyCart(ctx, &cart_checkout.CartEmptySignal{CartID: "hello123"})
	if err != nil {
		log.Error("Error emptying cart: ", err)
		c.JSON(500, err)
	} else if !out.Result {
		c.JSON(401, "Error")
	} else {
		c.JSON(200, responses.MessageResponse{Message: "Order Complete."})
	}
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
// @Router       / [get]
func HealthCheck(c *gin.Context) {
	c.JSON(200, responses.MessageResponse{Message: "up"})
}
