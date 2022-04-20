package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/cdp-team3/shipping-address-service/app/handlers"
	_ "github.com/cdp-team3/shipping-address-service/docs"
)
type ShippingRoutes struct {
    shippingHandler handlers.ShippingHandler
	healthCheckhandler handlers.HealthCheckHandler
}
func NewShippingRoutes(shippingHandler handlers.ShippingHandler, healthCheckhandler handlers.HealthCheckHandler) ShippingRoutes {
	return ShippingRoutes{shippingHandler: shippingHandler, healthCheckhandler: healthCheckhandler}
}
func (sr ShippingRoutes) InitRoutes(newRouter *gin.RouterGroup) {
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", sr.healthCheckhandler.HealthCheck)
    newRouter.GET("/shippingaddress/:id", sr.shippingHandler.GetShippingAddress())
	newRouter.POST("/shippingaddress", sr.shippingHandler.AddNewShippingAddress())
	newRouter.PUT("/shippingaddress/:id",  sr.shippingHandler.HandleUpdateShippingAddressByID())
    newRouter.DELETE("/shippingaddress/:id",  sr.shippingHandler.HandleDeleteShippingAddressById())

}