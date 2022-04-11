package app

import (
	"github.com/gin-gonic/gin"
	"io"
	"fmt"
	"os"
	"github.com/cdp-team3/categories-service/app/routes"
    logger	"github.com/cdp-team3/categories-service/utils/logger"
)

func Start() {
	file, err := os.OpenFile("categories-service-logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}

	router := gin.New()
   

	router.Use(logger.UseLogger(logger.DefaultLoggerFormatter), gin.Recovery())

	routes.InitRoutes(router)
	fmt.Println("Categories service started")

	router.Run(":3002")
}