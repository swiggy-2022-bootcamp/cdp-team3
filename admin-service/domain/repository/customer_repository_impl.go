package repository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/errors"
	shipping "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/grpc/shipping/proto"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/models"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GrpcFunction interface {
	DeleteShippingAddress(ShippingAddressDeleteRequest) string
}

func DeleteShippingProtoConv(customer *models.Customer) *ShippingAddressDeleteRequest {
	return &ShippingAddressDeleteRequest{
		ShippingAddressId: customer.Address.ShippingAddressId,
	}
}

type ShippingAddressDeleteRequest struct {
	ShippingAddressId string `json:shippingaddressid`
}

var customerCollection = "Customers"

type CustomerRepositoryImpl struct {
	customerDB *dynamodb.DynamoDB
}

func NewCustomerRepositoryImpl(customerDB *dynamodb.DynamoDB) CustomerRepository {
	return &CustomerRepositoryImpl{
		customerDB: customerDB,
	}
}

func (cr CustomerRepositoryImpl) AddCustomerToDB(customer *models.Customer) *errors.AppError {
	fetchedCustomer, _ := cr.GetCustomerByEmailFromDB(customer.Email)
	if fetchedCustomer != nil {
		zap.L().Error("User with the email Already Exists")
		return errors.UserAlreadyExists("User Already Exists")
	}
	data, err := dynamodbattribute.MarshalMap(customer)
	if err != nil {
		zap.L().Error("Marshalling of customer failed - " + err.Error())
		return errors.NewUnexpectedError(err.Error())
	}

	query := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(customerCollection),
	}

	_, err = cr.customerDB.PutItem(query)
	if err != nil {
		zap.L().Error("Failed to insert customer into database - " + err.Error())
		return errors.NewUnexpectedError(err.Error())
	}
	return nil
}

func (cr CustomerRepositoryImpl) GetAllCustomersFromDB() ([]models.Customer, *errors.AppError) {

	var customersList []models.Customer
	params := &dynamodb.ScanInput{
		TableName: aws.String(customerCollection),
	}

	err := cr.customerDB.ScanPages(params, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		var customers []models.Customer
		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &customers)
		if err != nil {
			zap.L().Error("\nCould not unmarshal AWS data: err =" + err.Error())
			return true
		}
		customersList = append(customersList, customers...)
		return true
	})

	if err != nil {
		return nil, errors.NewUnexpectedError("Error fetching data from DB " + err.Error())
	}
	return customersList, nil
}

func (cr CustomerRepositoryImpl) GetCustomerByIdFromDB(customerId string) (*models.Customer, *errors.AppError) {
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

func (cr CustomerRepositoryImpl) GetCustomerByEmailFromDB(emailId string) (*models.Customer, *errors.AppError) {

	customer := &models.Customer{}

	emailIndex := "email-index"
	params := &dynamodb.QueryInput{
		TableName:              aws.String(customerCollection),
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
		zap.L().Error("Failed to get item from database - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to get item from database - " + err.Error())
	}

	if len(result.Items) == 0 {
		zap.L().Error("Customer for given email doesn't exist - " + emailId)
		return nil, errors.NewNotFoundError("Customer for given email doesn't exist - " + emailId)
	}

	err = dynamodbattribute.UnmarshalMap(result.Items[0], &customer)
	if err != nil {
		zap.L().Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to unmarshal document fetched from DB - " + err.Error())
	}

	return customer, nil
}

func (cr CustomerRepositoryImpl) DeleteCustomerByIdFromDB(customerId string) (bool, *errors.AppError) {
	fetchedCustomer, _ := cr.GetCustomerByIdFromDB(customerId)
	if fetchedCustomer == nil {
		zap.L().Error("User with Id Doesn't Exist")
		return false, errors.UserAlreadyExists("User with Id Doesn't Exist")
	}
	deleteId := DeleteShippingProtoConv(fetchedCustomer)
	fmt.Println(deleteId)
	Deleted := DeleteShippingAddress(deleteId)
	if !Deleted {
		zap.L().Error("GRPC Call With Shipping Service Failed")
		return false, errors.NewUnexpectedError("GRPC Call With Shipping Service Failed")
	}
	_, err := configs.DB.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(customerCollection),
		Key: map[string]*dynamodb.AttributeValue{
			"customerId": {
				S: aws.String(customerId),
			},
		},
		ReturnValues: aws.String("ALL_OLD"),
	})
	if err != nil {
		zap.L().Error("Error Deleting Customer" + err.Error())
		return false, errors.NewUnexpectedError("Error Deleting Customer " + err.Error())
	}

	return true, nil
}

func (cr CustomerRepositoryImpl) UpdateCustomerByIdFromDB(customerId string, customer *models.Customer) (*models.Customer, *errors.AppError) {
	fetchedCustomer, _ := cr.GetCustomerByIdFromDB(customerId)
	if fetchedCustomer == nil {
		zap.L().Error("User with Id Doesn't Exist")
		return nil, errors.UserAlreadyExists("User with Id Doesn't Exist")
	}
	if fetchedCustomer.Email != customer.Email {
		fetchedCustomerWithEmail, _ := cr.GetCustomerByEmailFromDB(customer.Email)
		if fetchedCustomerWithEmail != nil {
			return nil, errors.UserAlreadyExists("User with Email Id Already Exists")
		}
	}
	customer.DateAdded = fetchedCustomer.DateAdded

	info, err := dynamodbattribute.MarshalMap(customer)
	if err != nil {
		return nil, errors.ParseFail("Marshalling Failed")
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String(customerCollection),
	}

	_, err = configs.DB.PutItem(input)
	if err != nil {
		return nil, errors.NewExpectationFailed(err.Error())
	}

	return customer, nil
}

func DeleteShippingAddress(request *ShippingAddressDeleteRequest) bool {

	conn, err := grpc.Dial(":"+configs.EnvShippingAddressPort(), grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error while making connection, %v\n", err)
		return false
	}

	c := shipping.NewShippingClient(conn)

	_, err1 := c.DeleteShippingAddress(
		context.Background(),
		&shipping.ShippingAddressDeleteRequest{
			ShippingAddressID: request.ShippingAddressId,
		},
	)

	return err1 == nil
}
