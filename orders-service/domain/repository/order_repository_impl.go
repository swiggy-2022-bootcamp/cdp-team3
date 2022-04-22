package repository

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/models"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/errors"
	"go.uber.org/zap"
)

var ordersCollection = "Orders"
type OrderRepositoryImpl struct {
	orderDB *dynamodb.DynamoDB
}

func NewOrderRepositoryImpl(orderDB *dynamodb.DynamoDB) OrderRepository {
	return &OrderRepositoryImpl{
		orderDB: orderDB,
	}
}

func (ori OrderRepositoryImpl) GetAllOrdersFromDB() ([]models.Order, *errors.AppError) {
	
	var ordersList []models.Order;
	params := &dynamodb.ScanInput{
		TableName: aws.String(ordersCollection),
	}

	err := ori.orderDB.ScanPages(params, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		var orders []models.Order
		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &orders)
		if err != nil {
			zap.L().Error("\nCould not unmarshal AWS data: err ="+err.Error())
			return true
		}
		ordersList = append(ordersList,orders...)
		return true
	})

	if err != nil {
		return nil, errors.NewUnexpectedError("Error fetching data from DB "+err.Error())
	}
	return ordersList, nil
}

func (ori OrderRepositoryImpl) GetOrdersByStatusFromDB(status string) ([]models.Order, *errors.AppError) {

	filt := expression.Name("status").Equal(expression.Value(status))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		zap.L().Error("Error constructing Expression")
		return nil, errors.NewUnexpectedError("Error constructing Expression "+err.Error())
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String(ordersCollection),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}

	res, err := configs.DB.Scan(input)

	if err != nil {
		zap.L().Error("Error Fetching data from DB")
		return nil, errors.NewUnexpectedError("Error Fetching data from DB "+err.Error())
	}

	var orders []models.Order

	if len(res.Items) == 0 {
		zap.L().Info("No orders found with status "+status)
		return nil, errors.NewNotFoundError("No orders found with status "+status)
	}

	if err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &orders); err != nil {
		zap.L().Error("Error unMarshalling Order"+err.Error())
		return nil, errors.NewUnexpectedError("Error unMarshalling Order "+err.Error())
	}
	return orders, nil
}

func(ori OrderRepositoryImpl) GetOrderByIdFromDB(orderId string) (*models.Order, *errors.AppError) {
	order := &models.Order{};

	query := &dynamodb.GetItemInput{
		TableName: aws.String(ordersCollection),
		Key: map[string]*dynamodb.AttributeValue{
			"orderId": {
				S: aws.String(orderId),
			},
		},
	}

	result, err := configs.DB.GetItem(query)

	if err != nil {
		zap.L().Error("Failed to get item from database - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to get item from database - "+err.Error())
	}

	if result.Item == nil {
		zap.L().Error("Order for given ID doesn't exists - ")
		return nil, errors.NewNotFoundError("Order for given ID doesn't exists - "+orderId)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &order)
	if err != nil {
		zap.L().Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to unmarshal document fetched from DB - "+err.Error())
	}

	return order, nil
}

func (ori OrderRepositoryImpl) UpdateStatusByIdInDB(orderId string,status string) (*models.Order, *errors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	order := &models.Order{};

	upd := expression.Set(expression.Name("status"), expression.Value(status))
	expr, err := expression.NewBuilder().WithUpdate(upd).Build()

	if err != nil {
		zap.L().Error("Error while forming expression"+err.Error())
		return nil, errors.NewUnexpectedError("Error while forming expression "+err.Error())
	}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: expr.Values(),
		Key: map[string]*dynamodb.AttributeValue{
			"orderId": {
				S: aws.String(orderId),
			},
		},
		TableName:        aws.String(ordersCollection),
		UpdateExpression: expr.Update(),
		ReturnValues:     aws.String("ALL_NEW"),
		ExpressionAttributeNames: expr.Names(),
	}

	response, err := configs.DB.UpdateItemWithContext(ctx, input)
	if err != nil {
		zap.L().Error("Error while Updating data in dynamoDB"+err.Error())
		return nil, errors.NewUnexpectedError("Error while Updating data in dynamoDB "+err.Error())
	}
	
	if response.Attributes == nil {
		zap.L().Error("Order for given ID doesn't exists - ")
		return nil, errors.NewNotFoundError("Order for given ID doesn't exists - "+orderId)
	}

	err = dynamodbattribute.UnmarshalMap(response.Attributes, &order)
	if err != nil {
		zap.L().Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil,  errors.NewUnexpectedError("Failed to unmarshal document fetched from DB -  "+err.Error())
	}

	return order, nil
}

func (ori OrderRepositoryImpl) DeleteOrderByIdInDB(orderId string) (*models.Order, *errors.AppError) {
	order := &models.Order{};

	response, err := configs.DB.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(ordersCollection),
		Key:      map[string]*dynamodb.AttributeValue{
			"orderId": {
				S: aws.String(orderId),
			},
		},
		ReturnValues: aws.String("ALL_OLD"),
	})
	
	if err != nil {
		zap.L().Error("Error Deleting Order"+err.Error())
		return nil, errors.NewUnexpectedError("Error Deleting Order "+err.Error())
	}
	err = dynamodbattribute.UnmarshalMap(response.Attributes, &order)
	if err != nil {
		zap.L().Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to unmarshal document fetched from DB -  "+err.Error())
	}

	return order, nil
}

func (ori OrderRepositoryImpl) GetOrdersByCustomerIdFromDB(customerId string) ([]models.Order, *errors.AppError) {
	filt := expression.Name("customerId").Equal(expression.Value(customerId))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		zap.L().Error("Error constructing Expression")
		return nil, errors.NewUnexpectedError("Error constructing Expression -  "+err.Error())
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String(ordersCollection),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}

	res, err := configs.DB.Scan(input)

	if err != nil {
		zap.L().Error("Error Fetching data from DB")
		return nil, errors.NewUnexpectedError("Error Fetching data from DB -  "+err.Error())
	}

	var orders []models.Order

	if len(res.Items) == 0 {
		zap.L().Error("No orders found for customer "+customerId)
		return nil, errors.NewNotFoundError("No orders found for customer "+customerId)
	}

	if err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &orders); err != nil {
		zap.L().Error("Error unMarshalling Order"+err.Error())
		return nil, errors.NewUnexpectedError("Error unMarshalling Order "+err.Error())
	}
	return orders, nil
}

