package repository

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/models"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/errors"
)

type OrderRepository interface {
	GetAllOrdersFromDB() ([]models.Order, *errors.AppError)
	GetOrdersByStatusFromDB(status string) ([]models.Order, *errors.AppError)
	GetOrderByIdFromDB(orderId string) (*models.Order, *errors.AppError)
	UpdateStatusByIdInDB(orderId string, status string) (*models.Order, *errors.AppError)
	DeleteOrderByIdInDB(orderId string) (*models.Order, *errors.AppError)

	GenerateInvoiceByIdInDB(orderId string) (*models.Order, *errors.AppError)
	GetOrdersByCustomerIdFromDB(customerId string) ([]models.Order, *errors.AppError)
}