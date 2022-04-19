package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/errors"
	model "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"
	"golang.org/x/crypto/bcrypt"
)

type CustomerCollectionInterface interface {
	Create(customer model.Customer) (*model.Customer, error)
	GetById(customerId string) (*model.Customer, error)
	GetByEmail(customerEmail string) (*model.Customer, error)
	Update(customer model.Customer) (*model.Customer, error)
	Delete(customerId string) (*string, error)
}

type CustomerCollection struct {
}

var db *dynamodb.DynamoDB

func init() {
	db = GetDynamoDBInstance()
}

func (customerCollection *CustomerCollection) Create(customer model.Customer) (*model.Customer, error) {
	fetchedCustomer, _ := customerCollection.GetByEmail(customer.Email)
	if fetchedCustomer != nil {
		return nil, errors.NewEmailAlreadyRegisteredError()
	}

	customer.CustomerId = uuid.New().String()
	hash, err := HashPassword(customer.Password)
	if err != nil {
		return nil, err
	}
	customer.Password = hash
	info, err := dynamodbattribute.MarshalMap(customer)
	if err != nil {
		return nil, errors.NewMarshallError()
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String("Team-3-Customers"),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (customerCollection *CustomerCollection) GetById(customerId string) (*model.Customer, error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String("Team-3-Customers"),
		Key: map[string]*dynamodb.AttributeValue{
			"customerId": {
				S: aws.String(customerId),
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

	var fetchedCustomer model.Customer
	dynamodbattribute.UnmarshalMap(resp.Item, &fetchedCustomer)
	return &fetchedCustomer, nil
}

func (customerCollection *CustomerCollection) GetByEmail(customerEmail string) (*model.Customer, error) {
	emailIndex := "email-index"
	params := &dynamodb.QueryInput{
		TableName:              aws.String("Team-3-Customers"),
		IndexName:              &emailIndex,
		KeyConditionExpression: aws.String("#email = :customersEmail"),
		ExpressionAttributeNames: map[string]*string{
			"#email": aws.String("email"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":customersEmail": {
				S: aws.String(customerEmail),
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

	var fetchedCustomer []model.Customer
	dynamodbattribute.UnmarshalListOfMaps(resp.Items, &fetchedCustomer)
	return &fetchedCustomer[0], nil
}

func (customerCollection *CustomerCollection) Update(customer model.Customer) (*model.Customer, error) {
	fetchedCustomer, err := customerCollection.GetById(customer.CustomerId)
	if err != nil {
		return nil, err
	}

	customer.DateAdded = fetchedCustomer.DateAdded
	if customer.Password != "" && !CheckPasswordHash(customer.Password, fetchedCustomer.Password) {
		hash, err := HashPassword(customer.Password)
		if err != nil {
			return nil, err
		}
		customer.Password = hash
	} else {
		customer.Password = fetchedCustomer.Password
	}

	if customer.Email == "" || customer.Email == fetchedCustomer.Email {
		customer.Email = fetchedCustomer.Email
	}

	if customer.Firstname == "" || customer.Firstname == fetchedCustomer.Firstname {
		customer.Firstname = fetchedCustomer.Firstname
	}

	if customer.Lastname == "" || customer.Lastname == fetchedCustomer.Lastname {
		customer.Lastname = fetchedCustomer.Lastname
	}

	if customer.Telephone == "" || customer.Telephone == fetchedCustomer.Telephone {
		customer.Telephone = fetchedCustomer.Telephone
	}

	info, err := dynamodbattribute.MarshalMap(customer)
	if err != nil {
		return nil, errors.NewMarshallError()
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String("Team-3-Customers"),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (customerCollection *CustomerCollection) Delete(customerId string) (*string, error) {
	allOld := "ALL_OLD"
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String("Team-3-Customers"),
		Key: map[string]*dynamodb.AttributeValue{
			"customerId": {
				S: aws.String(customerId),
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
