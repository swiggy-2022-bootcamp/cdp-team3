package repository

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/models"
)

var userCollection = "users"

type UserRepositoryImpl struct {
	userDB *dynamodb.DynamoDB
}

func NewUserRepository(userDB *dynamodb.DynamoDB) UserRepository {
	return &UserRepositoryImpl{
		userDB: userDB,
	}
}

func (uri UserRepositoryImpl) GetUserByEmail(email string) (*models.User, *errors.AppError) {
	var user *models.User
	input := &dynamodb.GetItemInput{
		TableName: &userCollection,
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: &email,
			},
		},
	}
	result, err := uri.userDB.GetItem(input)
	if err != nil {
		return user, errors.NewInternalServerError(err.Error())
	}
	if result.Item == nil {
		return user, errors.NewNotFoundError(errors.MsgUserDoesNotExists)
	}
	user.Email = *result.Item["email"].S
	user.Password = *result.Item["password"].S
	return user, nil
}
