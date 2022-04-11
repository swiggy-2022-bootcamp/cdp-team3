package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/utils"
)

func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := utils.HealthCheck()
		c.JSON(200, response)
	}
}

func DeepHealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := utils.DeepHealthCheck()
		c.JSON(200, response)
	}
}
