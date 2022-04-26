package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/domain/services"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/dto"
	order "github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/grpc/order"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/kafka"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/utils"
)

const requestTimeout = time.Second * 5

var logger = utils.NewLoggerService("payment-controller")

type PaymentController struct {
	paymentService services.PaymentService
}

func NewPaymentController(ps services.PaymentService) *PaymentController {
	return &PaymentController{paymentService: ps}
}

// Pay godoc
// @Summary Payment
// @Description This request is used for Payment
// @Tags Payment Service
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {string} 	message
// @Failure	400  {number} 	http.StatusBadRequest
// @Failure	401  {number} 	http.StatusUnauthorized
// @Failure	404  {number} 	http.StatusNotFound
// @Failure	500  {number} 	http.StatusInternalServerError
// @Router /pay [POST]
func (pc PaymentController) Pay() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pr dto.PaymentRequestDto
		if err := c.ShouldBindJSON(&pr); err != nil {
			logger.Log(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseDTO{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
			return
		}

		// TODO: Add Order Validation GRPC call

		status, err := order.GetOrderStatus(pr.OrderId)

		if status == "COMPLETED" {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseDTO{
				Message: "Order is already Paid",
				Status:  http.StatusBadRequest,
			})
			return
		}

		if err != nil {
			logger.Log(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.ResponseDTO{
				Message: err.Message,
				Status:  http.StatusBadRequest,
			})
			return
		}

		res, err := pc.paymentService.Pay(pr)
		if err != nil {
			logger.Log(err)
			c.AbortWithStatusJSON(err.Code, dto.ResponseDTO{
				Message: err.Message,
				Status:  err.Code,
			})
			kafka.UpdateOrderStatusProducer(pr.OrderId, "FAILED")
			return
		}

		kafka.UpdateOrderStatusProducer(pr.OrderId, "COMPLETED")

		c.JSON(http.StatusOK, dto.ResponseDTO{
			Message: res,
			Status:  http.StatusOK,
		})
	}
}
