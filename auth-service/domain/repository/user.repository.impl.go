package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/models"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/utils"
)

type UserRepositoryImpl struct {
	userDB *dynamodb.DynamoDB
}

func NewUserRepository(userDB *dynamodb.DynamoDB) UserRepository {
	return &UserRepositoryImpl{
		userDB: userDB,
	}
}

var logger = utils.NewLoggerService("user-repository")

func (uri UserRepositoryImpl) GetUserByEmail(email string) (*models.User, *errors.AppError) {
	var user models.User

	customer, err := getCustomerByEmailFromDB(email)

	if err == nil {
		email := customer.Email
		id := customer.CustomerId
		name := customer.Firstname + " " + customer.Lastname
		password := customer.Password
		isAdmin := customer.IsAdmin

		user.Email = email
		user.Id = id
		user.Name = name
		user.Password = password
		user.IsAdmin = isAdmin

		return &user, nil
	} else {
		admin, err := getAdminByEmailFromDB(email)

		if err == nil {
			email := admin.Email
			id := admin.AdminId
			name := admin.Firstname + " " + admin.Lastname
			password := admin.Password
			isAdmin := admin.IsAdmin

			user.Email = email
			user.Id = id
			user.Name = name
			user.Password = password
			user.IsAdmin = isAdmin

			return &user, nil
		}
	}
	return nil, errors.NewNotFoundError("User not found for given email - " + email)
}

func getCustomerByEmailFromDB(emailId string) (*models.Customer, *errors.AppError) {

	customer := &models.Customer{}

	emailIndex := "email-index"
	params := &dynamodb.QueryInput{
		TableName:              aws.String(models.CustomerTableName),
		IndexName:              &emailIndex,
		KeyConditionExpression: aws.String("#email = :email"),
		ExpressionAttributeNames: map[string]*string{
			"#email": aws.String("email"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":email": {
				S: aws.String(emailId),
			},
		},
	}

	result, err := configs.DB.Query(params)

	if err != nil {
		logger.Log("Failed to get item from database - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to get item from database - " + err.Error())
	}

	if len(result.Items) == 0 {
		logger.Log("Customer for given email doesn't exist - " + emailId)
		return nil, errors.NewNotFoundError("Customer for given email doesn't exist - " + emailId)
	}

	err = dynamodbattribute.UnmarshalMap(result.Items[0], &customer)
	if err != nil {
		logger.Log("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to unmarshal document fetched from DB - " + err.Error())
	}

	return customer, nil
}

func getAdminByEmailFromDB(emailId string) (*models.Admin, *errors.AppError) {

	admin := &models.Admin{}

	emailIndex := "email-index"
	params := &dynamodb.QueryInput{
		TableName:              aws.String(models.AdminTableName),
		IndexName:              &emailIndex,
		KeyConditionExpression: aws.String("#email = :email"),
		ExpressionAttributeNames: map[string]*string{
			"#email": aws.String("email"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":email": {
				S: aws.String(emailId),
			},
		},
	}

	result, err := configs.DB.Query(params)

	if err != nil {
		logger.Log("Failed to get item from database - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to get item from database - " + err.Error())
	}

	if len(result.Items) == 0 {
		logger.Log("admin for given email doesn't exist - " + emailId)
		return nil, errors.NewNotFoundError("admin for given email doesn't exist - " + emailId)
	}

	err = dynamodbattribute.UnmarshalMap(result.Items[0], &admin)
	if err != nil {
		logger.Log("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to unmarshal document fetched from DB - " + err.Error())
	}

	return admin, nil
}
