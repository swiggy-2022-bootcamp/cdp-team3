package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/configs"
	authproto "github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/grpc/auth/proto"
	"go.uber.org/zap"
)

func CheckLoggedInUserWithOrderCustomerId(c *gin.Context, orderCudtomerId string) bool {
	var userDetails = c.MustGet("user_details").(*authproto.VerifyTokenResponse)
	return userDetails.UserId == orderCudtomerId
}

func IsAdmin(c *gin.Context) bool {
	var userDetails = c.MustGet("user_details").(*authproto.VerifyTokenResponse)
	return userDetails.IsAdmin
}

func ClearCart(customerId string) {
	checkoutHost := configs.EnvCheckoutHost()
	checkoutURL := "http://" + checkoutHost + ":" + configs.EnvCheckOutPORT() + "/checkout_service/success"
	requestBody := strings.NewReader(`
		{
			"userId": "` + customerId + `" 
		}
	`)
	res, err := http.Post(checkoutURL, "application/json; charset=UTF-8", requestBody)

	if err != nil {
		zap.L().Error("Error Calling Rest call to checkout Service" + err.Error())
		return
	}
	res.Body.Close()
	if res.StatusCode == 200 {
		zap.L().Info("Cart cleared Successfully")
	} else {
		zap.L().Error("Error in clearing Cart")
	}
}
