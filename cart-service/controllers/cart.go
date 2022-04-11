package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-ipp/cart-service/dto/responses"
)

// @Summary      Create a Cart Item
// @Description  Create a Item in the Cart using the Cart Item data sent.
// @Tags         Cart Items
// @Accept       json
// @Produce      json
// @Param        cartItemDTO  body      requests.CartItemRequest  true  "Item to be added to the Cart"
// @Success      201          {object}  interface{}
// @Failure      400  {object}  errors.HTTPErrorDTO
// @Failure      404  {object}  errors.HTTPErrorDTO
// @Failure      500  {object}  nil
// @Router       /cart [post]
func CreateCartItem(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "Created",
	})
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
func GetCartItems(c *gin.Context) {
	c.JSON(200, responses.CartItemsResponse{})
}

// @Summary      Update a Cart Item
// @Description  Update a Item in the Cart using the Cart Item data sent.
// @Tags         Cart Items
// @Accept       json
// @Produce      json
// @Param        cartItemDTO  body      requests.CartItemRequest  true  "Item to be updated in the Cart"
// @Success      204          {object}  interface{}
// @Failure      400          {object}  errors.HTTPErrorDTO
// @Failure      404          {object}  errors.HTTPErrorDTO
// @Failure      500  {object}  nil
// @Router       /cart [put]
func UpdateCartItem(c *gin.Context) {
	c.JSON(204, gin.H{"message": "Updated"})
}

// @Summary      Deletes a Cart Item.
// @Description  Deletes an Item in the cart using its key.
// @Tags         Cart Items
// @Accept       json
// @Produce      json
// @Param        key  path      string  true  "Cart Item Key"
// @Success      204  {object}  int64
// @Failure      400  {object}  errors.HTTPErrorDTO
// @Failure      404  {object}  errors.HTTPErrorDTO
// @Failure      500          {object}  nil
// @Router       /cart/{key} [delete]
func DeleteCartItem(c *gin.Context) {
	c.JSON(204, gin.H{"message": "Deleted"})
}

// @Summary      Empty Cart.
// @Description  Empty the Cart for an user.
// @Tags         Cart Overall
// @Accept       json
// @Produce      json
// @Success      204  {object}  nil
// @Failure      400  {object}  errors.HTTPErrorDTO
// @Failure      404  {object}  errors.HTTPErrorDTO
// @Failure      500          {object}  nil
// @Router       /cart/empty [delete]
func EmptyCart(c *gin.Context) {
	c.JSON(204, gin.H{"message": "Deleted"})
}
