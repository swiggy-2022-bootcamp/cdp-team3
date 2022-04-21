package services

import "github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/models"

type OrderService interface {
	GetAllOrders() ([]models.Order, error)
	GetOrdersByStatus(status string) ([]models.Order, error)
	GetOrderById(orderId string) (*models.Order, error)
	UpdateStatusById(orderId string, orderStatus models.OrderStatus) (*models.Order, error)
	DeleteOrderById(orderId string) (*models.Order, error)

	GetOrdersByCustomerId(customerId string) ([]models.Order, error)
}