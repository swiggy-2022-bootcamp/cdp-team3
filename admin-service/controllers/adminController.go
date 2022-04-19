package controllers

import (
	"github.com/gin-gonic/gin"
)

// GetAdminDetails godoc
// @Summary Gets the details of the admin that is logged in
// @Description When a request is made, it returns the admin details
// @Tags Admin
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /user [GET]
func GetSelf(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Successfully Fetched Details"})
}
