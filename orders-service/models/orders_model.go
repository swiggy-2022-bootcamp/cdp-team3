package models

import (
	"time"
)

type Order struct {
	OrderId         string						 `json:"orderId" dynamodbav:"orderId" validate:"required"`
	DateTime        time.Time          `json:"dateTime" dynamodbav:"dateTime" validate:"required"`
	Status          string             `json:"status" dynamodbav:"status" validate:"required"`
	CustomerId      string             `json:"customerId" dynamodbav:"customerId" validate:"required" `
	TotalAmount			float64						 `json:"totalAmount" dynamodbav:"totalAmount" validate:"required" `
	OrderedProducts []OrderedProduct   `json:"orderedProducts" dynamodbav:"orderedProducts" validate:"required" `
}

type OrderedProduct struct {
	ProductId string `json:"productId" dynamodbav:"productId" validate:"required"`
	Quantity  int64  `json:"quantity" dynamodbav:"quantity" validate:"required"`
}

type OrderStatus struct {
	Status string `json:"status" dynamodbav:"status"`
}

type OrderInvoice struct {
	InvoiceId  string `json:"invoiceId" dynamodbav:"invoiceId" validate:"required"`
	OrderId string `json:"orderId" dynamodbav:"orderId" validate:"required"`
	TotalAmount float64  `json:"totalAmount" dynamodbav:"totalAmount" validate:"required"`
}