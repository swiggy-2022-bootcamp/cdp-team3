package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func CheckAuthorized(usergroup string) {

// }

// AddReward godoc
// @Summary Adds Reward Point To The Customer
// @Description Adds Reward Point To The Customer based on the given ID
// @Tags Reward
// @Schemes
// @Accept json
// @Produce json
// @Param        CustomerID path string  true "customer id"
// @Param        Reward Details body models.Reward true "reward details"
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router /reward/{id} [POST]
func AddReward(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Successfully Added Reward To The User"})
}

// Healthcheck godoc
// @Summary Checks whether the service is up & running
// @Description When a request is made to the / endpoint, if the service is running, it returns "Okay"
// @Tags Health
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {String} 	success
// @Failure	400  {number} 	400
// @Failure	500  {number} 	500
// @Router / [GET]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Okay",
	})
}
