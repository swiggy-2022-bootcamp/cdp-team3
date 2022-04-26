package admin

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	admin "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/grpc/admin/proto"
	"go.uber.org/zap"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"

	// "github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/utils"s
	"google.golang.org/grpc"
)

var customerCollection string = "Customers"

type TransactionAmount struct {
	admin.UnimplementedTransactionAmountServer
}

func SendTransactionAmount(transaction *admin.TransactionDetails) (*admin.SuccessMessage, error) {
	fetchedCustomer, _ := GetCustomerByIdFromDB(transaction.UserId)
	if fetchedCustomer == nil {
		zap.L().Error("User with Id Doesn't Exist")
		return &admin.SuccessMessage{
			IsAdded: "Failure",
		}, nil
	}
	_, err := UpdateTransactionPointsByIdFromDB(transaction.UserId, transaction.TransactionAmount)
	if err != nil {
		zap.L().Error("Failed To Update Transaction Points")
		return &admin.SuccessMessage{
			IsAdded: "Failure",
		}, nil
	}

	return &admin.SuccessMessage{
		IsAdded: "Success",
	}, nil
}

func InitialiseTransactionsServer() {
	port := configs.EnvGrpcTransactionServerPORT()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("Started Transaction GRPC server")

	s := grpc.NewServer()
	admin.RegisterTransactionAmountServer(s, &TransactionAmount{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func GetCustomerByIdFromDB(customerId string) (*models.Customer, *errors.AppError) {
	customer := &models.Customer{}

	query := &dynamodb.GetItemInput{
		TableName: aws.String(customerCollection),
		Key: map[string]*dynamodb.AttributeValue{
			"customerId": {
				S: aws.String(customerId),
			},
		},
	}

	result, err := configs.DB.GetItem(query)

	if err != nil {
		zap.L().Error("Failed to get item from database - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to get item from database - " + err.Error())
	}

	if result.Item == nil {
		zap.L().Error("Customer for given ID doesn't exist - " + customerId)
		return nil, errors.NewNotFoundError("Customer for given ID doesn't exist - " + customerId)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &customer)
	if err != nil {
		zap.L().Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to unmarshal document fetched from DB - " + err.Error())
	}

	return customer, nil
}

func UpdateTransactionPointsByIdFromDB(customerId string, transaction float32) (*models.Customer, *errors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	customer := &models.Customer{}
	upd := expression.Add(expression.Name("transaction_points"), expression.Value(transaction))
	expr, err := expression.NewBuilder().WithUpdate(upd).Build()

	if err != nil {
		zap.L().Error("Error while forming expression" + err.Error())
		return nil, errors.NewUnexpectedError("Error while forming expression " + err.Error())
	}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: expr.Values(),
		Key: map[string]*dynamodb.AttributeValue{
			"customerId": {
				S: aws.String(customerId),
			},
		},
		TableName:                aws.String(customerCollection),
		UpdateExpression:         expr.Update(),
		ReturnValues:             aws.String("ALL_NEW"),
		ExpressionAttributeNames: expr.Names(),
	}

	response, err := configs.DB.UpdateItemWithContext(ctx, input)
	if err != nil {
		zap.L().Error("Error while Updating data in dynamoDB" + err.Error())
		return nil, errors.NewUnexpectedError("Error while Updating data in dynamoDB " + err.Error())
	}

	if response.Attributes == nil {
		zap.L().Error("User Doesn't Exist")
		return nil, errors.NewNotFoundError("User Doesn't Exist For ID -" + customerId)
	}

	err = dynamodbattribute.UnmarshalMap(response.Attributes, &customer)
	if err != nil {
		zap.L().Error("Error in Unmarshalling" + err.Error())
		return nil, errors.NewUnexpectedError("Error in Unmarshalling" + err.Error())
	}

	return customer, nil
}
