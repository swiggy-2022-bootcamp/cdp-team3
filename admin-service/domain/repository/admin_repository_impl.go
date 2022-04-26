package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"
	"go.uber.org/zap"
)

var adminCollection = "Admins"

type AdminRepositoryImpl struct {
	adminDB *dynamodb.DynamoDB
}

func NewAdminRepositoryImpl(adminDB *dynamodb.DynamoDB) AdminRepository {
	return &AdminRepositoryImpl{
		adminDB: adminDB,
	}
}

func (ar AdminRepositoryImpl) AddAdminToDB(admin *models.Admin) *errors.AppError {
	data, err := dynamodbattribute.MarshalMap(admin)
	if err != nil {
		zap.L().Error("Marshalling of admin failed - " + err.Error())
		return errors.NewUnexpectedError(err.Error())
	}

	query := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(adminCollection),
	}

	_, err = ar.adminDB.PutItem(query)
	if err != nil {
		zap.L().Error("Failed to insert admin into database - " + err.Error())
		return errors.NewUnexpectedError(err.Error())
	}
	return nil
}

func (ar AdminRepositoryImpl) GetSelfFromDB(adminId string) (*models.Admin, *errors.AppError) {
	admin := &models.Admin{}

	query := &dynamodb.GetItemInput{
		TableName: aws.String(customerCollection),
		Key: map[string]*dynamodb.AttributeValue{
			"adminId": {
				S: aws.String(adminId),
			},
		},
	}

	result, err := configs.DB.GetItem(query)

	if err != nil {
		zap.L().Error("Failed to get item from database - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to get item from database - " + err.Error())
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &admin)
	if err != nil {
		zap.L().Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to unmarshal document fetched from DB - " + err.Error())
	}

	return admin, nil
}
