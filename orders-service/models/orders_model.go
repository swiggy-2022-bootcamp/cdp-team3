package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderId         primitive.ObjectID `json:"orderId,omitempty"`
	DateTime        time.Time          `json:"dateTime,omitempty" validate:"required"`
	Status          string             `json:"status,omitempty"`
	CustomerId      string             `json:"customerId,omitempty"`
	OrderedProducts []OrderedProduct   `json:"orderedProducts,omitempty"`
}

type OrderedProduct struct {
	ProductId string `json:"productId,omitempty"`
	Quantity  int64 `json:"quantity,omitempty"`
}

type OrderStatus struct {
	Status string `json:"status,omitempty"`
}

type OrderInvoice struct {
	InvoiceId  primitive.ObjectID `json:"invoiceId,omitempty"`
	OrderId string `json:"orderId,omitempty"`
	TotalAmount float64 `json:"totalAmount,omitempty"`
}