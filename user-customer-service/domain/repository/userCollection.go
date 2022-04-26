package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/errors"
	model "github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/models"
	"golang.org/x/crypto/bcrypt"
)

type UserCollectionInterface interface {
	Create(user model.User) (*model.User, error)
	GetById(userId string) (*model.User, error)
	GetByEmail(userEmail string) (*model.User, error)
	Update(user model.User) (*model.User, error)
	Delete(userId string) (*string, error)
}

type UserCollection struct {
}

var db *dynamodb.DynamoDB

func init() {
	db = GetDynamoDBInstance()
}

func (userCollection *UserCollection) Create(user model.User) (*model.User, error) {
	fetchedUser, _ := userCollection.GetByEmail(user.Email)
	if fetchedUser != nil {
		return nil, errors.NewEmailAlreadyRegisteredError()
	}

	user.UserId = uuid.New().String()
	hash, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hash
	info, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return nil, errors.NewMarshallError()
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String("Team-3-Users"),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userCollection *UserCollection) GetById(userId string) (*model.User, error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String("Team-3-Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(userId),
			},
		},
	}

	resp, err := db.GetItem(params)
	if err != nil {
		return nil, err
	}

	if len(resp.Item) == 0 {
		return nil, errors.NewUserNotFoundError()
	}

	var fetchedUser model.User
	dynamodbattribute.UnmarshalMap(resp.Item, &fetchedUser)
	return &fetchedUser, nil
}

func (userCollection *UserCollection) GetByEmail(userEmail string) (*model.User, error) {
	emailIndex := "email-index"
	params := &dynamodb.QueryInput{
		TableName:              aws.String("Team-3-Users"),
		IndexName:              &emailIndex,
		KeyConditionExpression: aws.String("#email = :usersEmail"),
		ExpressionAttributeNames: map[string]*string{
			"#email": aws.String("email"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":usersEmail": {
				S: aws.String(userEmail),
			},
		},
	}

	resp, err := db.Query(params)
	if err != nil {
		return nil, err
	}

	if len(resp.Items) == 0 {
		return nil, errors.NewUserNotFoundError()
	}

	var fetchedUser []model.User
	dynamodbattribute.UnmarshalListOfMaps(resp.Items, &fetchedUser)
	return &fetchedUser[0], nil
}

func (userCollection *UserCollection) Update(user model.User) (*model.User, error) {
	fetchedUser, err := userCollection.GetById(user.UserId)
	if err != nil {
		return nil, err
	}

	user.DateAdded = fetchedUser.DateAdded
	if user.Password != "" && !CheckPasswordHash(user.Password, fetchedUser.Password) {
		hash, err := HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hash
	} else {
		user.Password = fetchedUser.Password
	}

	if user.Email == "" || user.Email == fetchedUser.Email {
		user.Email = fetchedUser.Email
	}

	if user.Firstname == "" || user.Firstname == fetchedUser.Firstname {
		user.Firstname = fetchedUser.Firstname
	}

	if user.Lastname == "" || user.Lastname == fetchedUser.Lastname {
		user.Lastname = fetchedUser.Lastname
	}

	if user.Telephone == "" || user.Telephone == fetchedUser.Telephone {
		user.Telephone = fetchedUser.Telephone
	}

	info, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return nil, errors.NewMarshallError()
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String("Team-3-Users"),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userCollection *UserCollection) Delete(userId string) (*string, error) {
	allOld := "ALL_OLD"
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String("Team-3-Users"),
		Key: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(userId),
			},
		},
		ReturnValues: &allOld,
	}

	deletedItem, err := db.DeleteItem(params)
	if err != nil {
		return nil, err
	}

	if len(deletedItem.Attributes) == 0 {
		return nil, errors.NewUserNotFoundError()
	}

	str := "deletion successful"
	return &str, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
