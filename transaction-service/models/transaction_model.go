package models

type Transaction struct {
	TransactionId  string  `json:"transaction_id" dynamodbav:"transaction_id" `
	Amount         float64  `json:"amount" dynamodbav:"amount" validate:"required"`
	Description     string   `json:"description" dynamodbav:"description"`
	CustomerID  string `json:"customer_id" dynamodbav:"customer_id" validate:"required"`
}
