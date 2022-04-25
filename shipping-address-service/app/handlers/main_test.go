package handlers

import (
	"github.com/gin-gonic/gin"
	"os"
	"github.com/cdp-team3/shipping-address-service/mocks"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

// Creating test server with ShippingService mock.
func NewServer(shippingService *mocks.MockShippingService) *gin.Engine {
	shippingHandler := NewShippingHandler(shippingService)

	server := gin.Default()
	router := server.Group("shipping-service/api")
    
	router.GET("/shippingaddress/:id", shippingHandler.GetShippingAddress())
	router.POST("/shippingaddress", shippingHandler.AddNewShippingAddress())
	router.PUT("/shippingaddress/:id",  shippingHandler.HandleUpdateShippingAddressByID())
    router.DELETE("/shippingaddress/:id",  shippingHandler.HandleDeleteShippingAddressById())

	return server
}