package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/models"
	"go.uber.org/zap"
)

var usersCollection = "Users"

type UserRepositoryImpl struct {
	userDB *dynamodb.DynamoDB
}

func NewUserRepositoryImpl(UserDB *dynamodb.DynamoDB) UserRepository {
	return &UserRepositoryImpl{
		userDB: UserDB,
	}
}

func (rr UserRepositoryImpl) AddUserToDB(user *models.User) *errors.AppError {
	data, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		zap.L().Error("Marshalling of user failed - " + err.Error())
		return errors.NewUnexpectedError(err.Error())
	}

	query := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(usersCollection),
	}

	result, err := rr.userDB.PutItem(query)
	if err != nil {
		zap.L().Error("Failed to insert user into database - " + err.Error())
		return errors.NewUnexpectedError(err.Error())
	}
	fmt.Println(result)
	return nil
}

func (rr UserRepositoryImpl) GetAllUsersFromDB() ([]models.User, *errors.AppError) {

	var usersList []models.User
	params := &dynamodb.ScanInput{
		TableName: aws.String(usersCollection),
	}

	err := rr.userDB.ScanPages(params, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		var users []models.User
		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &users)
		if err != nil {
			zap.L().Error("\nCould not unmarshal AWS data: err =" + err.Error())
			return true
		}
		usersList = append(usersList, users...)
		return true
	})

	if err != nil {
		return nil, errors.NewUnexpectedError("Error fetching data from DB " + err.Error())
	}
	return usersList, nil
}

func (rr UserRepositoryImpl) GeUserByIdFromDB(userId string) (*models.User, *errors.AppError) {
	user := &models.User{}

	query := &dynamodb.GetItemInput{
		TableName: aws.String(usersCollection),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(userId),
			},
		},
	}

	result, err := configs.DB.GetItem(query)

	if err != nil {
		zap.L().Error("Failed to get item from database - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to get item from database - " + err.Error())
	}

	if result.Item == nil {
		zap.L().Error("User for given ID doesn't exist - " + userId)
		return nil, errors.NewNotFoundError("User for given ID doesn't exist - " + userId)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		zap.L().Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to unmarshal document fetched from DB - " + err.Error())
	}

	return user, nil
}

func (rr UserRepositoryImpl) GetUsersByCustomerIdFromDB(customerId string) ([]models.User, *errors.AppError) {
	filt := expression.Name("customerId").Equal(expression.Value(customerId))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		zap.L().Error("Error constructing Expression")
		return nil, errors.NewUnexpectedError("Error constructing Expression -  " + err.Error())
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String(usersCollection),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}

	res, err := configs.DB.Scan(input)

	if err != nil {
		zap.L().Error("Error Fetching data from DB")
		return nil, errors.NewUnexpectedError("Error Fetching data from DB -  " + err.Error())
	}

	var users []models.User

	if len(res.Items) == 0 {
		zap.L().Error("No users found for customer " + customerId)
		return nil, errors.NewNotFoundError("No users found for customer " + customerId)
	}

	if err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &users); err != nil {
		zap.L().Error("Error unMarshalling User" + err.Error())
		return nil, errors.NewUnexpectedError("Error unMarshalling User" + err.Error())
	}
	return users, nil
}
