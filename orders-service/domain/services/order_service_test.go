package services

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/models"
)

func TestOrderServiceImpl_GetOrderById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	order_id := "423fec6b-82c-4f99-8b2b-6eeef7605a37"

	successOrder := &models.Order{
		OrderId: order_id,
		DateTime: time.Now(),
		Status: "COMPLETED",
		CustomerId: "5243c6b-8a0c-4f99-8b2b-6eeef7605a37",
		TotalAmount: 500,
		OrderedProducts: []models.OrderedProduct{
			{
				ProductId: "3456ffc6b-8a0c-4f99-8b2b-6eeef7605a37",
				Quantity: 5,
			},
			{
				ProductId: "123ffc6b-8a0c-46599-8b2b-6eeef7605a37",
				Quantity: 4,
			},
		},
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockOrderRepository)
		checkResponse func(t *testing.T, order interface{}, err interface{})
	}{
		{
			name: "successOrderFound",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
				  GetOrderByIdFromDB(order_id).
					Times(1).
					Return(successOrder, nil)
			},
			checkResponse: func(t *testing.T, order interface{}, err interface{}) {
				assert.NotNil(t,order)
				assert.Nil(t,err)
			},
		},
		{
			name: "FailureOrderNotFound",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
					GetOrderByIdFromDB(order_id).
					Times(1).
					Return(nil, errors.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T, order interface{}, err interface{}) {
			    assert.Nil(t,order)
					assert.NotNil(t,err)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
				  GetOrderByIdFromDB(order_id).
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, order interface{}, err interface{}) {

				assert.Nil(t,order)
				assert.NotNil(t,err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			orderRepository := mocks.NewMockOrderRepository(ctrl)
			tc.buildStubs(orderRepository)

			orderServiceImpl := NewOrderServiceImpl(orderRepository)
			category, err := orderServiceImpl.GetOrderById(order_id)
			tc.checkResponse(t,category, err)
		})
	}
}

func TestOrderServiceImpl_GetAllOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)
	successOrders := []models.Order{
		{
			OrderId: "423fec6b-8a0c-4f99-8b2b-6eeef7605a37",
			DateTime: time.Now(),
			Status: "PENDING",
			CustomerId: "523fec6b-8a0c-4f99-8b2b-6eeef7605a37",
			TotalAmount: 200,
			OrderedProducts: []models.OrderedProduct{
				{
					ProductId: "333ffc6b-8a0c-4f99-8b2b-6eeef7605a37",
					Quantity: 5,
				},
				{
					ProductId: "123ffc6b-8a0c-4f99-8b2b-6eeef7605a37",
					Quantity: 3,
				},
			},
		},
		{
			OrderId: "423fec6b-82c-4f99-8b2b-6eeef7605a37",
			DateTime: time.Now(),
			Status: "COMPLETED",
			CustomerId: "5243c6b-8a0c-4f99-8b2b-6eeef7605a37",
			TotalAmount: 500,
			OrderedProducts: []models.OrderedProduct{
				{
					ProductId: "3456ffc6b-8a0c-4f99-8b2b-6eeef7605a37",
					Quantity: 5,
				},
				{
					ProductId: "123ffc6b-8a0c-46599-8b2b-6eeef7605a37",
					Quantity: 4,
				},
			},
		},
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockOrderRepository)
		checkResponse func(t *testing.T,res []models.Order, err interface{})
	}{
		{
			name: "SuccessOrdersFound",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
					GetAllOrdersFromDB().
					Times(1).
					Return(successOrders, nil)
			},
			checkResponse: func(t *testing.T,res []models.Order, err interface{}) {
				assert.NotNil(t,res)
				assert.Nil(t,err)
			},
		},
		{
			name: "FailureOrderNotFound",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
					GetAllOrdersFromDB().
					Times(1).
					Return(nil, errors.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T,res []models.Order, err interface{}) {
				assert.Nil(t,res)
				assert.NotNil(t,err)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
					GetAllOrdersFromDB().
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T,res []models.Order, err interface{}) {
				assert.Nil(t,res)
				assert.NotNil(t,err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			orderRepository := mocks.NewMockOrderRepository(ctrl)
			tc.buildStubs(orderRepository)

			orderServiceImpl := NewOrderServiceImpl(orderRepository)
			res, err := orderServiceImpl.GetAllOrders()
			tc.checkResponse(t,  res, err)
		})
	}
}

func TestOrderServiceImpl_GetOrdersByStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)

	status := "PENDING"
	successOrders := []models.Order{
		{
			OrderId: "423fec6b-8a0c-4f99-8b2b-6eeef7605a37",
			DateTime: time.Now(),
			Status: "PENDING",
			CustomerId: "523fec6b-8a0c-4f99-8b2b-6eeef7605a37",
			TotalAmount: 200,
			OrderedProducts: []models.OrderedProduct{
				{
					ProductId: "333ffc6b-8a0c-4f99-8b2b-6eeef7605a37",
					Quantity: 5,
				},
				{
					ProductId: "123ffc6b-8a0c-4f99-8b2b-6eeef7605a37",
					Quantity: 3,
				},
			},
		},
		{
			OrderId: "423fec6b-82c-4f99-8b2b-6eeef7605a37",
			DateTime: time.Now(),
			Status: "COMPLETED",
			CustomerId: "5243c6b-8a0c-4f99-8b2b-6eeef7605a37",
			TotalAmount: 500,
			OrderedProducts: []models.OrderedProduct{
				{
					ProductId: "3456ffc6b-8a0c-4f99-8b2b-6eeef7605a37",
					Quantity: 5,
				},
				{
					ProductId: "123ffc6b-8a0c-46599-8b2b-6eeef7605a37",
					Quantity: 4,
				},
			},
		},
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockOrderRepository)
		checkResponse func(t *testing.T,res []models.Order, err interface{})
	}{
		{
			name: "SuccessOrdersByStatusFound",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
					GetOrdersByStatusFromDB(status).
					Times(1).
					Return(successOrders, nil)
			},
			checkResponse: func(t *testing.T,res []models.Order, err interface{}) {
				assert.NotNil(t,res)
				assert.Nil(t,err)
			},
		},
		{
			name: "FailureOrdersByStatusNotFound",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
					GetOrdersByStatusFromDB(status).
					Times(1).
					Return(nil, errors.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T,res []models.Order, err interface{}) {
				assert.Nil(t,res)
				assert.NotNil(t,err)
			},
		},
		{
			name: "FailureOrdersByStatusUnexpectedError",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
					GetOrdersByStatusFromDB(status).
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T,res []models.Order, err interface{}) {
				assert.Nil(t,res)
				assert.NotNil(t,err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			orderRepository := mocks.NewMockOrderRepository(ctrl)
			tc.buildStubs(orderRepository)

			orderServiceImpl := NewOrderServiceImpl(orderRepository)
			res, err := orderServiceImpl.GetOrdersByStatus(status)
			tc.checkResponse(t,  res, err)
		})
	}
}

func TestOrderServiceImpl_GetOrdersByCustomerId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	customerId := "523fec6b-8a0c-4f99-8b2b-6eeef7605a37"
	successOrders := []models.Order{
		{
			OrderId: "423fec6b-8a0c-4f99-8b2b-6eeef7605a37",
			DateTime: time.Now(),
			Status: "PENDING",
			CustomerId: "523fec6b-8a0c-4f99-8b2b-6eeef7605a37",
			TotalAmount: 200,
			OrderedProducts: []models.OrderedProduct{
				{
					ProductId: "333ffc6b-8a0c-4f99-8b2b-6eeef7605a37",
					Quantity: 5,
				},
				{
					ProductId: "123ffc6b-8a0c-4f99-8b2b-6eeef7605a37",
					Quantity: 3,
				},
			},
		},
		{
			OrderId: "423fec6b-82c-4f99-8b2b-6eeef7605a37",
			DateTime: time.Now(),
			Status: "COMPLETED",
			CustomerId: "5243c6b-8a0c-4f99-8b2b-6eeef7605a37",
			TotalAmount: 500,
			OrderedProducts: []models.OrderedProduct{
				{
					ProductId: "3456ffc6b-8a0c-4f99-8b2b-6eeef7605a37",
					Quantity: 5,
				},
				{
					ProductId: "123ffc6b-8a0c-46599-8b2b-6eeef7605a37",
					Quantity: 4,
				},
			},
		},
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockOrderRepository)
		checkResponse func(t *testing.T,res []models.Order, err interface{})
	}{
		{
			name: "SuccessOrdersByCustomerIdFound",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
				GetOrdersByCustomerIdFromDB(customerId).
					Times(1).
					Return(successOrders, nil)
			},
			checkResponse: func(t *testing.T,res []models.Order, err interface{}) {
				assert.NotNil(t,res)
				assert.Nil(t,err)
			},
		},
		{
			name: "FailureOrdersByCustomerIdNotFound",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
				GetOrdersByCustomerIdFromDB(customerId).
					Times(1).
					Return(nil, errors.NewNotFoundError(""))
			},
			checkResponse: func(t *testing.T,res []models.Order, err interface{}) {
				assert.Nil(t,res)
				assert.NotNil(t,err)
			},
		},
		{
			name: "FailureOrdersByCustomerIdUnexpectedError",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
					GetOrdersByCustomerIdFromDB(customerId).
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T,res []models.Order, err interface{}) {
				assert.Nil(t,res)
				assert.NotNil(t,err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			orderRepository := mocks.NewMockOrderRepository(ctrl)
			tc.buildStubs(orderRepository)

			orderServiceImpl := NewOrderServiceImpl(orderRepository)
			res, err := orderServiceImpl.GetOrdersByCustomerId(customerId)
			tc.checkResponse(t,  res, err)
		})
	}
}

func TestOrderServiceImpl_DeleteOrderById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	order_id := "bb912edc-50d9-42d7-b7a1-9ce66d459thj"
	
	deletedOrder := &models.Order{
		OrderId: order_id,
		DateTime: time.Now(),
		Status: "COMPLETED",
		CustomerId: "5243c6b-8a0c-4f99-8b2b-6eeef7605a37",
		TotalAmount: 500,
		OrderedProducts: []models.OrderedProduct{
			{
				ProductId: "3456ffc6b-8a0c-4f99-8b2b-6eeef7605a37",
				Quantity: 5,
			},
			{
				ProductId: "123ffc6b-8a0c-46599-8b2b-6eeef7605a37",
				Quantity: 4,
			},
		},
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockOrderRepository)
		checkResponse func(t *testing.T, order interface{} , err interface{})
	}{
		{
			name: "SuccessOrderDeleted",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
					DeleteOrderByIdInDB(order_id).
					Times(1).
					Return(deletedOrder, nil)
			},
			checkResponse: func(t *testing.T, order interface{}, err interface{}) {
				assert.NotNil(t,order)
				assert.Nil(t,err)
			},
		},
		{
			name: "FailureOrderNotDeleted",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
				  DeleteOrderByIdInDB(order_id).
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, order interface{}, err interface{}) {
				assert.Nil(t,order)
				assert.NotNil(t,err)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
				  DeleteOrderByIdInDB(order_id).
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, order interface{}, err interface{}) {
				assert.Nil(t,order)
				assert.NotNil(t,err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			orderRepository := mocks.NewMockOrderRepository(ctrl)
			tc.buildStubs(orderRepository)

			orderServiceImpl := NewOrderServiceImpl(orderRepository)
			value, err := orderServiceImpl.DeleteOrderById(order_id)
			tc.checkResponse(t, value, err)
		})
	}
}

func TestOrderServiceImpl_UpdateStatusById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	order_id := "bb912edc-50d9-42d7-b7a1-9ce66d459thj"
	
	updateOrderStatus := &models.OrderStatus{
		Status: "COMPLETED",
	}

	updatedOrder := &models.Order{
		OrderId: order_id,
		DateTime: time.Now(),
		Status: "COMPLETED",
		CustomerId: "5243c6b-8a0c-4f99-8b2b-6eeef7605a37",
		TotalAmount: 500,
		OrderedProducts: []models.OrderedProduct{
			{
				ProductId: "3456ffc6b-8a0c-4f99-8b2b-6eeef7605a37",
				Quantity: 5,
			},
			{
				ProductId: "123ffc6b-8a0c-46599-8b2b-6eeef7605a37",
				Quantity: 4,
			},
		},
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockOrderRepository)
		checkResponse func(t *testing.T, order interface{} , err interface{})
	}{
		{
			name: "SuccessOrderStatusUpdated",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
				  UpdateStatusByIdInDB(order_id,updateOrderStatus.Status).
					Times(1).
					Return(updatedOrder, nil)
			},
			checkResponse: func(t *testing.T,order interface{}, err interface{}) {
				assert.NotNil(t,order)
				assert.Equal(t, order, updatedOrder)
				assert.Nil(t,err)
			},
		},
		{
			name: "FailureOrderNotUpdated",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
					UpdateStatusByIdInDB(order_id,updateOrderStatus.Status).
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, order interface{}, err interface{}) {
				assert.Nil(t,order)
				assert.NotNil(t,err)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
					UpdateStatusByIdInDB(order_id,updateOrderStatus.Status).
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, order interface{}, err interface{}) {
				assert.Nil(t,order)
				assert.NotNil(t,err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

		    orderRepository := mocks.NewMockOrderRepository(ctrl)
			tc.buildStubs(orderRepository)

			orderServiceImpl := NewOrderServiceImpl(orderRepository)
			value, err := orderServiceImpl.UpdateStatusById(order_id, updateOrderStatus)
			tc.checkResponse(t,value, err)
		})
	}
}

func TestOrderServiceImpl_GenerateInvoiceById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	order_id := "bb912edc-50d9-42d7-b7a1-9ce66d459thj"

	updatedInvoiceOrder := &models.Order{
		OrderId: order_id,
		DateTime: time.Now(),
		Status: "COMPLETED",
		CustomerId: "5243c6b-8a0c-4f99-8b2b-6eeef7605a37",
		TotalAmount: 500,
		InvoiceId: "524dvf66b-8a0c-4f99-8b2b-6eeef7605a37",
		OrderedProducts: []models.OrderedProduct{
			{
				ProductId: "3456ffc6b-8a0c-4f99-8b2b-6eeef7605a37",
				Quantity: 5,
			},
			{
				ProductId: "123ffc6b-8a0c-46599-8b2b-6eeef7605a37",
				Quantity: 4,
			},
		},
	}

	testCases := []struct {
		name          string
		buildStubs    func(repository *mocks.MockOrderRepository)
		checkResponse func(t *testing.T, order interface{} , err interface{})
	}{
		{
			name: "SuccessGeneratedInvoiceById",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
					GenerateInvoiceByIdInDB(order_id).
					Times(1).
					Return(updatedInvoiceOrder, nil)
			},
			checkResponse: func(t *testing.T,order interface{}, err interface{}) {
				assert.NotNil(t,order)
				assert.Equal(t, order, updatedInvoiceOrder)
				assert.Nil(t,err)
			},
		},
		{
			name: "FailureInvoiceNotUpdated",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
					GenerateInvoiceByIdInDB(order_id).
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, order interface{}, err interface{}) {
				assert.Nil(t,order)
				assert.NotNil(t,err)
			},
		},
		{
			name: "FailureUnexpectedError",
			buildStubs: func(repository *mocks.MockOrderRepository) {
				repository.EXPECT().
				GenerateInvoiceByIdInDB(order_id).
					Times(1).
					Return(nil, errors.NewUnexpectedError(""))
			},
			checkResponse: func(t *testing.T, order interface{}, err interface{}) {
				assert.Nil(t,order)
				assert.NotNil(t,err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

		    orderRepository := mocks.NewMockOrderRepository(ctrl)
			tc.buildStubs(orderRepository)

			orderServiceImpl := NewOrderServiceImpl(orderRepository)
			value, err := orderServiceImpl.GenerateInvoiceById(order_id)
			tc.checkResponse(t,value, err)
		})
	}
}

// func TestOrderServiceImpl_AddCategory(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	categoryDesc := models.OrderDescription{
	
// 		Name    :"TestCategoryDescName",        
	
// 		Description     : "TestCategoryDescDescription",
// 		MetaDescription :"TestCategoryDescMetaDescription",
// 		MetaKeyword     :"TestCategoryDescMetaKeyword",
// 		MetaTitle       :"TestCategoryDescMetaTitle",

// }
// category := &models.Order{
	
// 	CategoryId: "bb912edc-50d9-42d7-b7a1-9ce66d459thj",
// 	CategoryDescription : []models.OrderDescription{categoryDesc},
	
	
// }

// 	testCases := []struct {
// 		name          string
// 		buildStubs    func(repository *mocks.MockOrderRepository)
// 		checkResponse func(t *testing.T,err interface{})
// 	}{
	
// 		{
// 			name: "SuccessAddCategory",
// 			buildStubs: func(repository *mocks.MockOrderRepository) {
// 				repository.EXPECT().
// 				AddCategoryToDB(category).
// 					Times(1).
// 					Return(nil)
// 			},
// 			checkResponse: func(t *testing.T,err interface{}) {
// 				assert.Nil(t,err)
// 			},
// 		},
// 		{
// 			name: "FailureAddCategory",
// 			buildStubs: func(repository *mocks.MockOrderRepository) {
// 				repository.EXPECT().
// 				    AddCategoryToDB(category).
// 					Times(1).
// 					Return(errors.NewUnexpectedError(""))

// 			},
// 			checkResponse: func(t *testing.T,err interface{}) {
// 				assert.NotNil(t,err)
// 			},
// 		},
// 		{
// 			name: "FailureUnexpectedError",
// 			buildStubs: func(repository *mocks.MockOrderRepository) {
// 				repository.EXPECT().
// 				    AddCategoryToDB(category).
// 					Times(1).
// 					Return(errors.NewUnexpectedError(""))
// 			},
// 			checkResponse: func(t *testing.T, err interface{}) {
// 				assert.NotNil(t,err)
			
// 			},
// 		},
	
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			orderRepository := mocks.NewMockOrderRepository(ctrl)
// 			tc.buildStubs(orderRepository)

// 			orderServiceImpl := NewOrderServiceImpl(orderRepository)
// 			err := orderServiceImpl.AddCategory(category)
// 			tc.checkResponse(t, err)
// 		})
// 	}
// }