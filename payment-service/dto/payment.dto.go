package dto

import "github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/models"

type PaymentRequestDto struct {
	UserId  string                      `json:"userId"`
	Amount  float64                     `json:"amount"`
	OrderId string                      `json:"orderId"`
	Details models.ModeOfPaymentDetails `json:"details"`
}

type ResponseDTO struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}
