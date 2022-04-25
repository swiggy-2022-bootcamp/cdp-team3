package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/errors"
	kafkaPro "github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/kafka"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/models"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/utils"
	"go.uber.org/zap"
)
var validate = validator.New()

type OrderServiceImpl struct {
	orderRepository repository.OrderRepository
}

func NewOrderServiceImpl(orderRepository repository.OrderRepository) OrderService {
	return &OrderServiceImpl{orderRepository: orderRepository}
}

func (os OrderServiceImpl) GetAllOrders() ([]models.Order, *errors.AppError) {
	zap.L().Info("Inside GetAllOrders Service")
	result,err := os.orderRepository.GetAllOrdersFromDB()
	if err != nil {
		zap.L().Error(err.Message)
		return nil,err
	}
	return result,nil
}

func(os OrderServiceImpl) GetOrdersByStatus(status string) ([]models.Order, *errors.AppError) {
	zap.L().Info("Inside GetOrdersByStatus Service")
	result,err := os.orderRepository.GetOrdersByStatusFromDB(status)
	if err != nil {
		zap.L().Error(err.Message)
		return nil,err
	}
	return result,nil
}

func(os OrderServiceImpl) GetOrderById(orderId string) (*models.Order, *errors.AppError) {
	zap.L().Info("Inside GetOrderById Service")
	result,err := os.orderRepository.GetOrderByIdFromDB(orderId)
	if err != nil {
		zap.L().Error(err.Message)
		return nil,err
	}
	return result,nil
}

func (os OrderServiceImpl) UpdateStatusById(orderId string, orderStatus *models.OrderStatus) (*models.Order, *errors.AppError) {
	zap.L().Info("Inside UpdateStatusById Service")

	//use the validator library to validate required fields
	if validationErr := validate.Struct(orderStatus); validationErr != nil {
		zap.L().Error("Invalid Request")
		return nil, errors.NewBadRequestError("Invalid Request "+validationErr.Error())
	}

	if !(orderStatus.Status == "COMPLETED" || orderStatus.Status == "FAILED") {
		zap.L().Error("Invalid Status")
		return nil, errors.NewBadRequestError("Invalid Status")
	}
	
	result, err := os.orderRepository.UpdateStatusByIdInDB(orderId, orderStatus.Status)
	if err != nil {
		zap.L().Error(err.Message)
		return nil,err
	}

	if orderStatus.Status == "COMPLETED" {
		//Updating 10 persent of Order Amount as Transaction Amount Service
		go kafkaPro.AddTransactionAmountProducer(result.CustomerId, result.TotalAmount * float64(0.10))

		//Calling Checkout Service to Clear the Cart
		go utils.ClearCart(result.CustomerId);
		
	}
	return result,nil
}

func (os OrderServiceImpl) DeleteOrderById(orderId string) (*models.Order, *errors.AppError) {
	zap.L().Info("Inside DeleteOrderById Service")
	result,err := os.orderRepository.DeleteOrderByIdInDB(orderId)
	if err != nil {
		zap.L().Error(err.Message)
		return nil,err
	}
	return result,nil
}

func(os OrderServiceImpl) GetOrdersByCustomerId(customerId string) ([]models.Order, *errors.AppError) {
	zap.L().Info("Inside GetOrdersByCustomerId Service")
	result,err := os.orderRepository.GetOrdersByCustomerIdFromDB(customerId)
	if err != nil {
		zap.L().Error(err.Message)
		return nil,err
	}
	return result,nil
}

func(os OrderServiceImpl) GenerateInvoiceById(orderId string) (*models.Order, *errors.AppError) {
	zap.L().Info("Inside GenerateInvoiceById Service")
	result,err := os.orderRepository.GenerateInvoiceByIdInDB(orderId)
	if err != nil {
		zap.L().Error(err.Message)
		return nil,err
	}
	return result,nil
}