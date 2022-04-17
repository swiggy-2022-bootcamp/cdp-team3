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
func (tr ShippingRoutes) InitRoutes(newRouter *gin.RouterGroup) {
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", tr.healthCheckhandler.HealthCheck)

	
	newRouter.GET("/shippingaddress/:id", tr.shippingHandler.GetShippingAddress)
	newRouter.POST("/shippingaddress", tr.shippingHandler.AddNewShippingAddress)
	newRouter.PUT("/shippingaddress/:id",  tr.shippingHandler.HandleUpdateShippingAddressByID())
    newRouter.DELETE("/shippingaddress/:id",  tr.shippingHandler.HandleDeleteShippingAddressById())

}