package order

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/configs"
	order "github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/grpc/order/proto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/models"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const orderCollection = "Orders"
var validate = validator.New()

type server struct {
	order.OrderServiceServer
}

func (s *server) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	zap.L().Info("Inside Create Order Protobuf")
	orderFromClient := req.Order
	var orderedProducts []models.OrderedProduct

	for _,v := range orderFromClient.OrderedProducts {
		orderedProducts = append(orderedProducts, models.OrderedProduct{ProductId: v.ProductId, Quantity: v.Quantity})
	}

	newOrder := models.Order{
		OrderId: uuid.New().String(),
		DateTime: time.Now(),
		Status: "PENDING",
		CustomerId: orderFromClient.CustomerId,
		TotalAmount: float64(orderFromClient.TotalAmount),
		ShippingAddressId: orderFromClient.ShippingAddressId,
		OrderedProducts: orderedProducts,
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&newOrder); validationErr != nil {
			zap.L().Error("Required fields not present"+validationErr.Error())
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Required fields not present: %v", validationErr.Error()),
			)
	}

	data, err := dynamodbattribute.MarshalMap(newOrder)
	if err != nil {
		zap.L().Error("Marshalling of order failed - " + err.Error())
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Marshalling of order failed: %v", err.Error()),
		)
	}
	query := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(orderCollection),
	}
	result, err := configs.DB.PutItem(query)
	if err != nil {
		zap.L().Error("Failed to add order - " + err.Error())
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed to add order: %v", err.Error()),
		)
	}
	zap.L().Info("Successfully created Order"+result.GoString())
	return &order.CreateOrderResponse{
		Order: &order.ResponseOrder{
			OrderId: newOrder.OrderId,
			DateTime: newOrder.DateTime.String(),
			Status: newOrder.Status,
			CustomerId: newOrder.CustomerId,
			TotalAmount: float32(newOrder.TotalAmount),
			OrderedProducts: orderFromClient.OrderedProducts,
		},
	}, nil
}

func InitializeGRPCServer(port string) {
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		zap.L().Error("Failed to listen"+err.Error())
	}else {
		s := grpc.NewServer()
		order.RegisterOrderServiceServer(s, &server{})

		if err := s.Serve(lis); err != nil {
			zap.L().Error("Error in starting grpc server"+err.Error())
			fmt.Println("Error in starting grpc server",err)
		}
	}
}