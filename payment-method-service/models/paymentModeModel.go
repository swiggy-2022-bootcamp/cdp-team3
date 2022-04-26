package models

type ModeOfPaymentDetails struct {
	PaymentMode string `json:"paymentMode"`
	Agree       bool   `json:"agree"`
	Message     string `json:"message"`
}
