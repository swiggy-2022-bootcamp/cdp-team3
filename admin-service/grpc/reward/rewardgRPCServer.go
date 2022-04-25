package reward

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	rewards "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/grpc/reward/proto"
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

type RewardPoints struct {
	rewards.UnimplementedRewardPointsServer
}

func (s *RewardPoints) SendRewardPoints(ctx context.Context, req *rewards.RewardDetails) (*rewards.SuccessMessage, error) {
	fetchedCustomer, _ := GetCustomerByIdFromDB(req.UserId)
	if fetchedCustomer == nil {
		zap.L().Error("User with Id Doesn't Exist")
		return &rewards.SuccessMessage{
			IsAdded: "Failure",
		}, nil
	}
	_, err := UpdateRewardPointsByIdFromDB(req.UserId, req.Reward)
	if err != nil {
		zap.L().Error("Failed To Update Rewards")
		return &rewards.SuccessMessage{
			IsAdded: "Failure",
		}, nil
	}

	return &rewards.SuccessMessage{
		IsAdded: "Success",
	}, nil

}

func InitialiseRewardsServer() {
	port := configs.EnvGrpcRewardServerPORT()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("Started Reward GRPC server")

	s := grpc.NewServer()
	rewards.RegisterRewardPointsServer(s, &RewardPoints{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func GetCustomerByIdFromDB(customerId string) (*models.Customer, *errors.AppError) {
	customer := &models.Customer{}

	query := &dynamodb.GetItemInput{
		TableName: aws.String("Customers"),
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

func UpdateRewardPointsByIdFromDB(customerId string, rewards int32) (*models.Customer, *errors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	customer := &models.Customer{}
	upd := expression.Add(expression.Name("rewards"), expression.Value(rewards))
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
