package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/models"
	"go.uber.org/zap"
)
var validate = validator.New()

type OrderServiceImpl struct {
	orderRepository repository.OrderRepository
}

func NewOrderServiceImpl(orderRepository repository.OrderRepository) OrderService {
	return &OrderServiceImpl{orderRepository: orderRepository}
}

func (os OrderServiceImpl) GetAllOrders() ([]models.Order, error) {
	zap.L().Info("Inside GetAllOrders Service")
	result,err := os.orderRepository.GetAllOrdersFromDB()
	if err != nil {
		zap.L().Error(err.Error())
		return nil,err
	}
	return result,nil
}

func(os OrderServiceImpl) GetOrdersByStatus(status string) ([]models.Order, error) {
	zap.L().Info("Inside GetOrdersByStatus Service")
	result,err := os.orderRepository.GetOrdersByStatusFromDB(status)
	if err != nil {
		zap.L().Error(err.Error())
		return nil,err
	}
	return result,nil
}

func(os OrderServiceImpl) GetOrderById(orderId string) (*models.Order, error) {
	zap.L().Info("Inside GetOrderById Service")
	result,err := os.orderRepository.GetOrderByIdFromDB(orderId)
	if err != nil {
		zap.L().Error(err.Error())
		return nil,err
	}
	return result,nil
}

func (os OrderServiceImpl) UpdateStatusById(orderId string, orderStatus models.OrderStatus) (*models.Order, error) {
	zap.L().Info("Inside UpdateStatusById Service")

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&orderStatus); validationErr != nil {
		zap.L().Error("Invalid Request")
		return nil, validationErr
	}

	result, err := os.orderRepository.UpdateStatusByIdInDB(orderId, orderStatus.Status)
	if err != nil {
		zap.L().Error(err.Error())
		return nil,err
	}
	return result,nil
}

func (os OrderServiceImpl) DeleteOrderById(orderId string) (*models.Order, error) {
	zap.L().Info("Inside DeleteOrderById Service")
	result,err := os.orderRepository.DeleteOrderByIdInDB(orderId)
	if err != nil {
		zap.L().Error(err.Error())
		return nil,err
	}
	return result,nil
}

func(os OrderServiceImpl) GetOrdersByCustomerId(customerId string) ([]models.Order, error) {
	zap.L().Info("Inside GetOrdersByCustomerId Service")
	result,err := os.orderRepository.GetOrdersByCustomerIdFromDB(customerId)
	if err != nil {
		zap.L().Error(err.Error())
		return nil,err
	}
	return result,nil
}