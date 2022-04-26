package controllers

import (
	"github.com/gin-gonic/gin"
)

// GetAvailablePaymentMethods godoc
// @Summary fetches all the available payment methods
// @Description fetches the available payment methods on the portal
// @Tags Payment Method
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /paymentmethods [GET]
func GetAllPaymentMethods() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"COD": "Cash on Delivery",
			"Net": "Internet Banking",
			"CC":  "Credit Card",
			"DC":  "Debit Card",
			"UPI": "Unified Payments Interface",
		})
	}
}
