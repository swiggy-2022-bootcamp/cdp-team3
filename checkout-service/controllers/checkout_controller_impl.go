package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/swiggy-ipp/checkout-service/dto/errors"
	"github.com/swiggy-ipp/checkout-service/dto/responses"
	"github.com/swiggy-ipp/checkout-service/grpcs/cart_checkout"
	"github.com/swiggy-ipp/checkout-service/grpcs/shipping_checkout"
)

// Implementation of the CheckoutController interface
type checkoutControllerImpl struct {
	cartCheckoutGRPCClient     cart_checkout.CartCheckoutServiceClient
	shippingCheckoutGRPCClient shipping_checkout.ShippingClient
}

// Create a new Cart Controller
func NewCheckoutController(
	cartCheckoutGRPCClient cart_checkout.CartCheckoutServiceClient,
	shippingCheckoutGRPCClient shipping_checkout.ShippingClient,
) CheckoutController {
	return &checkoutControllerImpl{
		cartCheckoutGRPCClient:     cartCheckoutGRPCClient,
		shippingCheckoutGRPCClient: shippingCheckoutGRPCClient,
	}
}

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
func (cc *checkoutControllerImpl) GetOrderOverview(c *gin.Context) {
	// Set up context
	ctx := c.Request.Context()
	// Make GRPC Call
	out, err := cc.shippingCheckoutGRPCClient.GetShippingAddress(ctx, &shipping_checkout.ShippingAddressRequest{})
	if err != nil {
		log.Error("Error Getting Order Overview: ", err)
		c.JSON(http.StatusInternalServerError, errors.NewHTTPErrorDTO(http.StatusInternalServerError, err, "Error Getting Order Overview"))
	} else {
		c.JSON(http.StatusOK, out)
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
func (cc *checkoutControllerImpl) OrderCompleteWebhook(c *gin.Context) {
	// Set up context
	ctx := c.Request.Context()
	// Make GRPC Call
	out, err := cc.cartCheckoutGRPCClient.EmptyCart(ctx, &cart_checkout.CartEmptySignal{CartID: "hello123"})
	if err != nil {
		log.Error("Error emptying cart: ", err)
		c.JSON(http.StatusInternalServerError, errors.NewHTTPErrorDTO(http.StatusInternalServerError, err, "Error emptying cart"))
	} else if !out.Result {
		c.JSON(http.StatusUnauthorized, "Error")
	} else {
		c.JSON(http.StatusOK, responses.MessageResponse{Message: "Order Completed."})
	}
}

// @Summary      Health Check Endpoint
// @Description  Health Check Endpoint
// @Tags         Health Check
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.HealthCheckResponse  "Health Check Response."
// @Failure      400  {object}  errors.HTTPErrorDTO
// @Failure      404  {object}  errors.HTTPErrorDTO
// @Failure      500  {object}  nil
// @Router       / [get]
func (cc *checkoutControllerImpl) HealthCheck(c *gin.Context) {
	// Generate DTO
	c.JSON(200, responses.HealthCheckResponse{
		ServiceHealth:     "OK",
		KafkaServerHealth: "OK",
	})
}
