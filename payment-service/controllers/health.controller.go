package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/utils"
)

// HealthCheck godoc
// @Summary HealthCheck
// @Description This request is used to check the health of the entire service at once
// @Tags Auth Service
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {object} 	utils.HealthCheckResponse
// @Failure 500  {number} 	http.StatusInternalServerError
// @Router / [GET]
func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := utils.HealthCheck()
		c.JSON(200, response)
	}
}

// HealthCheck godoc
// @Summary Deep HealthCheck
// @Description This request is used to check the health of the every single service at once
// @Tags Auth Service
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {object} 	utils.HealthCheckResponse
// @Failure 500  {number} 	http.StatusInternalServerError
// @Router /deep [GET]
func DeepHealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := utils.DeepHealthCheck()
		c.JSON(200, response)
	}
}
