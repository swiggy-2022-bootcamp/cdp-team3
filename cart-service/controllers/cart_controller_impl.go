package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-ipp/cart-service/dto/errors"
	"github.com/swiggy-ipp/cart-service/dto/requests"
	"github.com/swiggy-ipp/cart-service/dto/responses"
	"github.com/swiggy-ipp/cart-service/grpcs/auth/proto"
	"github.com/swiggy-ipp/cart-service/services"
)

// Implementation of the CartController interface
type cartControllerImpl struct {
	cartService services.CartService
}

// Create a new Cart Controller
func NewCartController(cartService services.CartService) CartController {
	return &cartControllerImpl{cartService: cartService}
}

// @Summary      Create a Cart Item
// @Description  Create a Item in the Cart using the Cart Item data sent.
// @Tags         Cart Items
// @Accept       json
// @Produce      json
// @Param        cartItemDTO  body      requests.CartItemRequest   true  "Item to be added to the Cart"
// @Success      201          {object}  responses.MessageResponse  "Message denoting whether successfully created"
// @Failure      400     {object}  errors.HTTPErrorDTO
// @Failure      404     {object}  errors.HTTPErrorDTO
// @Failure      500        {object}  nil
// @Router       /cart [post]
func (cc *cartControllerImpl) CreateCartItem(c *gin.Context) {
	// Get User Claims
	claims := c.MustGet("claims").(*proto.VerifyTokenResponse)
	// Get Cart Item Request DTO Object
	cartItemDTO := requests.CartItemRequest{}
	if err := c.ShouldBindJSON(&cartItemDTO); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewHTTPErrorDTO(http.StatusBadRequest, err))
	} else if claims.GetUserId() == "" {
		c.JSON(http.StatusForbidden, errors.NewHTTPErrorDTO(http.StatusForbidden, nil, "You are not authorized to perform this action."))
	} else {
		// Create Cart Item
		err := cc.cartService.CreateCartItem(c.Request.Context(), &cartItemDTO, claims.GetUserId())
		if err != nil {
			c.JSON(http.StatusInternalServerError, errors.NewHTTPErrorDTO(http.StatusInternalServerError, err))
		} else {
			c.JSON(http.StatusCreated, responses.MessageResponse{Message: "Created"})
		}
	}
}

// @Summary      Get all Cart Items
// @Description  Get a list of all the items in the Cart currently. Only one of Cart ID or User ID must be provided
// @Tags         Cart Items
// @Accept       json
// @Produce      json
// @Param        cartID  query     string                       false  "ID of the Cart to get the items for."
// @Param        userID  query     string                       false  "ID of the User to get the items for."
// @Success      200     {object}  responses.CartItemsResponse  "List of Cart Items"
// @Failure      400          {object}  errors.HTTPErrorDTO
// @Failure      404          {object}  errors.HTTPErrorDTO
// @Failure      500     {object}  nil
// @Router       /cart [get]
func (cc *cartControllerImpl) GetCartItems(c *gin.Context) {
	cartID, userID := c.Query("cartID"), c.Query("userID")
	if cartID == "" && userID == "" {
		c.JSON(
			http.StatusBadRequest,
			errors.NewHTTPErrorDTO(http.StatusBadRequest, nil, "Either Cart ID or User ID must be provided."),
		)
	} else if cartID != "" && userID != "" {
		c.JSON(
			http.StatusBadRequest,
			errors.NewHTTPErrorDTO(http.StatusBadRequest, nil, "Ambiguous Request. Both Cart ID and User ID are provided."),
		)
	} else {
		// Get User Claims
		claims := c.MustGet("claims").(*proto.VerifyTokenResponse)
		if (cartID != "" && claims.GetIsAdmin()) || (userID != "" && userID == claims.GetUserId()) {
			res, err := cc.cartService.GetCartItems(c.Request.Context(), cartID, userID)
			if err != nil {
				c.JSON(
					http.StatusInternalServerError,
					errors.NewHTTPErrorDTO(http.StatusInternalServerError, err, "Error while getting Cart Items."),
				)
			}
			c.JSON(http.StatusOK, res)
		} else {
			c.JSON(
				http.StatusForbidden,
				errors.NewHTTPErrorDTO(http.StatusForbidden, nil, "You are not authorized to perform this action."),
			)
		}
	}
}

// @Summary      Update a Cart Item
// @Description  Update Quantity of a Item in the Cart using the Cart Item data sent.
// @Tags         Cart Items
// @Accept       json
// @Produce      json
// @Param        cartItemDTO  body      requests.CartItemRequest   true  "Item to be updated in the Cart"
// @Success      204          {object}  responses.MessageResponse  "Message denoting whether successfully updated"
// @Failure      400          {object}  errors.HTTPErrorDTO
// @Failure      404          {object}  errors.HTTPErrorDTO
// @Failure      500           {object}  nil
// @Router       /cart [put]
func (cc *cartControllerImpl) UpdateCartItem(c *gin.Context) {
	// Get User Claims
	claims := c.MustGet("claims").(*proto.VerifyTokenResponse)
	// Get Cart Item Request DTO Object
	cartItemDTO := requests.CartItemRequest{}
	if err := c.ShouldBindJSON(&cartItemDTO); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewHTTPErrorDTO(http.StatusBadRequest, err))
	} else if claims.GetUserId() == "" {
		c.JSON(http.StatusForbidden, errors.NewHTTPErrorDTO(http.StatusForbidden, nil, "You are not authorized to perform this action."))
	} else {
		// Create Cart Item
		err := cc.cartService.UpdateCartItem(c.Request.Context(), &cartItemDTO, claims.GetUserId())
		if err != nil {
			c.JSON(http.StatusInternalServerError, errors.NewHTTPErrorDTO(http.StatusInternalServerError, err))
		} else {
			c.JSON(http.StatusNoContent, responses.MessageResponse{Message: "Updated"})
		}
	}
}

// @Summary      Delete a Cart Item.
// @Description  Delete an Item in the cart using its Product ID.
// @Tags         Cart Items
// @Accept       json
// @Produce      json
// @Param        productID  path      string  true  "Cart Item Product ID"
// @Success      204        {object}  int64
// @Failure      400        {object}  errors.HTTPErrorDTO
// @Failure      404        {object}  errors.HTTPErrorDTO
// @Failure      500          {object}  nil
// @Router       /cart/{productID} [delete]
func (cc *cartControllerImpl) DeleteCartItem(c *gin.Context) {
	// Get User Claims
	claims := c.MustGet("claims").(*proto.VerifyTokenResponse)
	// Get Cart Item Request DTO Object
	cartItemDTO := requests.CartItemRequest{}
	if err := c.ShouldBindJSON(&cartItemDTO); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewHTTPErrorDTO(http.StatusBadRequest, err))
	} else if claims.GetUserId() == "" {
		c.JSON(http.StatusForbidden, errors.NewHTTPErrorDTO(http.StatusForbidden, nil, "You are not authorized to perform this action."))
	} else {
		// Get Cart Item Product ID from URL
		productID := c.Param("productID")
		// Create Cart Item
		err := cc.cartService.DeleteCartItem(c.Request.Context(), productID, claims.GetUserId())
		if err != nil {
			c.JSON(http.StatusInternalServerError, errors.NewHTTPErrorDTO(http.StatusInternalServerError, err))
		} else {
			c.JSON(http.StatusNoContent, responses.MessageResponse{Message: "Deleted"})
		}
	}
}

// @Summary      Empty the Cart.
// @Description  Empty the Cart for an user.
// @Tags         Cart Overall
// @Accept       json
// @Produce      json
// @Param        emptyCartDTO  body      requests.CartIDRequest  true  "Cart ID Request DTO. Must Either provide User ID (user request)  or  Cart  ID  (Admin Request),  but  not  both."
// @Success      204           {object}  nil
// @Failure      400           {object}  errors.HTTPErrorDTO
// @Failure      404           {object}  errors.HTTPErrorDTO
// @Failure      500          {object}  nil
// @Router       /cart/empty [delete]
func (cc *cartControllerImpl) EmptyCart(c *gin.Context) {
	// Extract the request
	var emptyCartRequest requests.CartIDRequest
	if err := c.ShouldBindJSON(&emptyCartRequest); err != nil {
		c.JSON(http.StatusBadRequest, errors.HTTPErrorDTO{Code: http.StatusBadRequest, Message: err.Error()})
	} else if emptyCartRequest.CartID == "" && emptyCartRequest.UserID == "" {
		c.JSON(
			http.StatusBadRequest,
			errors.NewHTTPErrorDTO(http.StatusBadRequest, nil, "Either Cart ID or User ID must be provided."),
		)
	} else if emptyCartRequest.CartID != "" && emptyCartRequest.UserID != "" {
		c.JSON(
			http.StatusBadRequest,
			errors.NewHTTPErrorDTO(http.StatusBadRequest, nil, "Ambiguous Request. Both Cart ID and User ID are provided."),
		)
	} else {
		// Get User Claims
		claims := c.MustGet("claims").(*proto.VerifyTokenResponse)
		if (emptyCartRequest.CartID != "" && claims.GetIsAdmin()) || (emptyCartRequest.UserID != "" && emptyCartRequest.UserID == claims.GetUserId()) {
			err := cc.cartService.EmptyCart(c.Request.Context(), emptyCartRequest)
			if err != nil {
				c.JSON(
					http.StatusInternalServerError,
					errors.NewHTTPErrorDTO(http.StatusInternalServerError, err, "Error while emptying Cart."),
				)
			}
			c.JSON(204, responses.MessageResponse{Message: "Cart Emptied"})
		} else {
			c.JSON(
				http.StatusForbidden,
				errors.NewHTTPErrorDTO(http.StatusForbidden, nil, "You are not authorized to perform this action."),
			)
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
func (cc *cartControllerImpl) HealthCheck(c *gin.Context) {
	// Check DB health by simple Read
	err := cc.cartService.DBHealthCheck(c.Request.Context())
	var dbHealth string = "OK"
	if err != nil {
		dbHealth = "FAIL"
	}

	// Generate DTO
	c.JSON(200, responses.HealthCheckResponse{
		ServiceHealth:     "OK",
		KafkaServerHealth: "OK",
		DBHealth:          dbHealth,
	})
}

func isAdmin() bool {
	return true
}
