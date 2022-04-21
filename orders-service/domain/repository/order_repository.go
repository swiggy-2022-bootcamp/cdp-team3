package repository

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/models"
)

type OrderRepository interface {
	GetAllOrdersFromDB() ([]models.Order, error)
	GetOrdersByStatusFromDB(status string) ([]models.Order, error)
	GetOrderByIdFromDB(orderId string) (*models.Order, error)
	UpdateStatusByIdInDB(orderId string, status string) (*models.Order, error)
	DeleteOrderByIdInDB(orderId string) (*models.Order, error)

	GetOrdersByCustomerIdFromDB(customerId string) ([]models.Order, error)
}