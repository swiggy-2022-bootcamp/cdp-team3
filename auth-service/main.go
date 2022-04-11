package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/routes"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/utils"
)

var logger = utils.NewLoggerService("main")

func main() {
	router := gin.Default()

	configs.ConnectDB() 
	logger.Log("Connected to DB")

	routes.AuthRoutes(router)
	routes.HealthCheckRoutes(router)

	router.Run("localhost:" + configs.EnvPORT())
}
