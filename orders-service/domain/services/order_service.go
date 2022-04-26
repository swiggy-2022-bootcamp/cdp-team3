package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/models"
)

type OrderService interface {
	GetAllOrders() ([]models.Order, *errors.AppError)
	GetOrdersByStatus(status string) ([]models.Order, *errors.AppError)
	GetOrderById(orderId string) (*models.Order, *errors.AppError)
	UpdateStatusById(orderId string, orderStatus *models.OrderStatus) (*models.Order, *errors.AppError)
	DeleteOrderById(orderId string) (*models.Order, *errors.AppError)

	GenerateInvoiceById(orderId string) (*models.Order, *errors.AppError)
	GetOrdersByCustomerId(customerId string) ([]models.Order, *errors.AppError)
}