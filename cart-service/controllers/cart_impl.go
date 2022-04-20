package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-ipp/cart-service/dto/responses"
	"github.com/swiggy-ipp/cart-service/repositories"
)

// Implementation of the CartController interface
type cartControllerImpl struct {
	cartRepository repositories.CartRepository
}

// Create a new Cart Controller
func NewCartController(cartRepository repositories.CartRepository) CartController {
	return &cartControllerImpl{cartRepository: cartRepository}
}

// @Summary      Create a Cart Item
// @Description  Create a Item in the Cart using the Cart Item data sent.
// @Tags         Cart Items
// @Accept       json
// @Produce      json
// @Param        cartItemDTO  body      requests.CartItemRequest   true  "Item to be added to the Cart"
// @Success      201          {object}  responses.MessageResponse  "Message denoting whether successfully created"
// @Failure      400  {object}  errors.HTTPErrorDTO
// @Failure      404  {object}  errors.HTTPErrorDTO
// @Failure      500  {object}  nil
// @Router       /cart [post]
func (cc *cartControllerImpl) CreateCartItem(c *gin.Context) {
	c.JSON(201, responses.MessageResponse{Message: "Created"})
}

// @Summary      Get all Cart Items
// @Description  Get a list of all the items in the Cart currently.
// @Tags         Cart Items
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.CartItemsResponse  "List of Cart Items"
// @Failure      400          {object}  errors.HTTPErrorDTO
// @Failure      404          {object}  errors.HTTPErrorDTO
// @Failure      500  {object}  nil
// @Router       /cart [get]
func (cc *cartControllerImpl) GetCartItems(c *gin.Context) {
	c.JSON(200, responses.CartItemsResponse{})
}

// @Summary      Update a Cart Item
// @Description  Update a Item in the Cart using the Cart Item data sent.
// @Tags         Cart Items
// @Accept       json
// @Produce      json
// @Param        cartItemDTO  body      requests.CartItemRequest   true  "Item to be updated in the Cart"
// @Success      204          {object}  responses.MessageResponse  "Message denoting whether successfully updated"
// @Failure      400          {object}  errors.HTTPErrorDTO
// @Failure      404          {object}  errors.HTTPErrorDTO
// @Failure      500  {object}  nil
// @Router       /cart [put]
func (cc *cartControllerImpl) UpdateCartItem(c *gin.Context) {
	c.JSON(204, responses.MessageResponse{Message: "Updated"})
}

// @Summary      Delete a Cart Item.
// @Description  Delete an Item in the cart using its key.
// @Tags         Cart Items
// @Accept       json
// @Produce      json
// @Param        key  path      string  true  "Cart Item Key"
// @Success      204  {object}  int64
// @Failure      400  {object}  errors.HTTPErrorDTO
// @Failure      404  {object}  errors.HTTPErrorDTO
// @Failure      500          {object}  nil
// @Router       /cart/{key} [delete]
func (cc *cartControllerImpl) DeleteCartItem(c *gin.Context) {
	c.JSON(204, responses.MessageResponse{Message: "Deleted"})
}

// @Summary      Empty the Cart.
// @Description  Empty the Cart for an user.
// @Tags         Cart Overall
// @Accept       json
// @Produce      json
// @Success      204  {object}  nil
// @Failure      400  {object}  errors.HTTPErrorDTO
// @Failure      404  {object}  errors.HTTPErrorDTO
// @Failure      500          {object}  nil
// @Router       /cart/empty [delete]
func (cc *cartControllerImpl) EmptyCart(c *gin.Context) {
	c.JSON(204, responses.MessageResponse{Message: "Deleted"})
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
func (cc *cartControllerImpl) HealthCheck(c *gin.Context) {
	c.JSON(200, responses.MessageResponse{Message: "up"})
}
