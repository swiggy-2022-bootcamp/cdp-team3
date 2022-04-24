package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/swiggy-ipp/checkout-service/dto/errors"
	"github.com/swiggy-ipp/checkout-service/dto/requests"
	"github.com/swiggy-ipp/checkout-service/dto/responses"
	authProto "github.com/swiggy-ipp/checkout-service/grpcs/auth/proto"
	"github.com/swiggy-ipp/checkout-service/grpcs/cart_checkout"
	orderProto "github.com/swiggy-ipp/checkout-service/grpcs/order/proto"
	"github.com/swiggy-ipp/checkout-service/grpcs/shipping_checkout"
)

// Implementation of the CheckoutController interface
type checkoutControllerImpl struct {
	cartCheckoutGRPCClient     cart_checkout.CartCheckoutServiceClient
	shippingCheckoutGRPCClient shipping_checkout.ShippingClient
	orderGRPCClient            orderProto.OrderServiceClient
}

// Create a new Cart Controller
func NewCheckoutController(
	cartCheckoutGRPCClient cart_checkout.CartCheckoutServiceClient,
	shippingCheckoutGRPCClient shipping_checkout.ShippingClient,
	orderGRPCClient orderProto.OrderServiceClient,
) CheckoutController {
	return &checkoutControllerImpl{
		cartCheckoutGRPCClient:     cartCheckoutGRPCClient,
		shippingCheckoutGRPCClient: shippingCheckoutGRPCClient,
		orderGRPCClient:            orderGRPCClient,
	}
}

// @Summary      Creates a pending order if it is valid and gets an overview of it.
// @Description  Get an overview of the order from the database.
// @Tags         Checkout API
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.MessageResponse  "Order Overview Data"
// @Failure      400            {object}  errors.HTTPErrorDTO
// @Failure      404            {object}  errors.HTTPErrorDTO
// @Failure      500            {object}  nil
// @Router       /confirm [post]
func (cc *checkoutControllerImpl) GetOrderOverview(c *gin.Context) {
	// Get User Claims
	claims := c.MustGet("claims").(*authProto.VerifyTokenResponse)
	if claims.GetUserId() == "" {
		c.JSON(http.StatusForbidden, errors.NewHTTPErrorDTO(http.StatusForbidden, nil, "User ID needs to be sent as Authorization."))
	} else {
		// Validate Shipping Address
		shippingOut, err := cc.shippingCheckoutGRPCClient.GetShippingAddress(c.Request.Context(), &shipping_checkout.ShippingAddressRequest{})
		if err != nil {
			log.Error("Error Getting Shipping Address: ", err)
			c.JSON(http.StatusInternalServerError, errors.NewHTTPErrorDTO(http.StatusInternalServerError, err, "Error Getting Shipping Address"))
			return
		} else if validateShippingAddress(shippingOut) == false {
			c.JSON(http.StatusBadRequest, errors.NewHTTPErrorDTO(http.StatusBadRequest, nil, "Shipping Address is not valid."))
			return
		}

		// Validate Cart Items Exist
		cartOut, err := cc.cartCheckoutGRPCClient.GetCartItems(c.Request.Context(), &cart_checkout.CartIDSignal{UserID: claims.UserId})
		if err != nil {
			log.Error("Error Getting Cart Items: ", err)
			c.JSON(http.StatusInternalServerError, errors.NewHTTPErrorDTO(http.StatusInternalServerError, err, "Error Getting Cart Items"))
			return
		} else if len(cartOut.CartItems) == 0 {
			c.JSON(http.StatusBadRequest, errors.NewHTTPErrorDTO(http.StatusBadRequest, nil, "Cart is empty. Nothing to Checkout."))
			return
		}

		// Make GRPC Call to Order Service
		orderedProducts := mapCartItemsToOrderedProducts(cartOut.CartItems)
		out, err := cc.orderGRPCClient.CreateOrder(c.Request.Context(), &orderProto.CreateOrderRequest{
			Order: &orderProto.RequestOrder{
				CustomerId:      claims.UserId,
				OrderedProducts: orderedProducts,
			},
		})
		if err != nil {
			log.Error("Error Creating Order: ", err)
			c.JSON(http.StatusInternalServerError, errors.NewHTTPErrorDTO(http.StatusInternalServerError, err, "Error Creating Order"))
			return
		}
		c.JSON(http.StatusOK, out)
	}
}

// @Summary      Order Successful Webhook
// @Description  Webhook hit when Order is successful to clear Cart and Unset Session Data.
// @Tags         Checkout API
// @Accept       json
// @Produce      json
// @Param        userIDRequest  body      requests.UserIDRequest     true  "User ID Request DTO."
// @Success      200            {object}  responses.MessageResponse  "Cart Cleared message DTO."
// @Failure      400  {object}  errors.HTTPErrorDTO
// @Failure      404  {object}  errors.HTTPErrorDTO
// @Failure      500  {object}  nil
// @Router       /confirm/success [post]
func (cc *checkoutControllerImpl) OrderCompleteWebhook(c *gin.Context) {
	// Get User ID from Request
	userIDRequest := &requests.UserIDRequest{}
	if err := c.ShouldBindJSON(userIDRequest); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewHTTPErrorDTO(http.StatusBadRequest, err, "User ID Request DTO is not valid."))
	} else {
		// Make GRPC Call
		out, err := cc.cartCheckoutGRPCClient.EmptyCart(c.Request.Context(), &cart_checkout.CartIDSignal{UserID: userIDRequest.UserID})
		if err != nil {
			log.Error("Error emptying cart: ", err)
			c.JSON(http.StatusInternalServerError, errors.NewHTTPErrorDTO(http.StatusInternalServerError, err, "Error emptying cart"))
		} else if !out.Result {
			c.JSON(http.StatusUnauthorized, "Failed to update Cart in DB.")
		} else {
			c.JSON(http.StatusOK, responses.MessageResponse{Message: "Order Completed."})
		}
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

// Utility function to validate Shipping Address and return boolean if valid
func validateShippingAddress(address *shipping_checkout.ShippingAddressResponse) bool {
	return address.GetAddress1() != "" && address.GetPostcode() != 0 && address.GetCity() != "" && address.GetFirstname() != ""
}

// Utility function to map Cart Items to Order Products Data Type for cross service compatibility
func mapCartItemsToOrderedProducts(cartItem []*cart_checkout.CartItem) []*orderProto.OrderedProduct {
	orderedProducts := make([]*orderProto.OrderedProduct, len(cartItem))
	for i, item := range cartItem {
		orderedProducts[i] = &orderProto.OrderedProduct{
			ProductId: item.ProductID,
			Quantity:  item.Quantity,
		}
	}
	return orderedProducts
}
