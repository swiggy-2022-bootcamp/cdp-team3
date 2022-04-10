package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthcheck godoc
// @Summary Checks whether the service is up & running
// @Description When a request is made to the / endpoint, if the service is running, it returns "Okay"
// @Tags Health
// @Schemes
// @Accept json
// @Produce json
// @Success	201  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router / [GET]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Okay",
	})
}
