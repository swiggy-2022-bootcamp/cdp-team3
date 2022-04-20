package repository

import (
	"fmt"
	"time"
	"context"
	"strconv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
//	"github.com/google/uuid"
	//"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	apperrors "github.com/cdp-team3/shipping-address-service/app-errors"
	"github.com/cdp-team3/shipping-address-service/domain/models"
	"github.com/cdp-team3/shipping-address-service/utils/logger"
	
)

const shippingCollection = "ShippingAddress"


type ShippingRepositoryImpl struct {
	shippingDB  *dynamodb.DynamoDB
}
func NewShippingRepositoryImpl(shippingDB *dynamodb.DynamoDB) ShippingRepository {
	return &ShippingRepositoryImpl{
		shippingDB: shippingDB,
		
	}
}



func (s ShippingRepositoryImpl) DBHealthCheck() bool {

	_, err := s.shippingDB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		logger.Error("Database connection is down.")
		return false
	}
	return true
}
func (s ShippingRepositoryImpl) InsertShippingAddressToDB(shippingAddress *models.ShippingAddress) (*apperrors.AppError) {

	fmt.Println("inside repo",shippingAddress)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	av, err := dynamodbattribute.MarshalMap(shippingAddress)
	if err != nil {
		return  apperrors.NewUnexpectedError(err.Error())
	}
	fmt.Println(shippingAddress)
    fmt.Println("\n")
	fmt.Println(shippingAddress)
	fmt.Println(av)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(shippingCollection),
	}

	_, err = s.shippingDB.PutItemWithContext(ctx, input)

	if err != nil {
		return apperrors.NewUnexpectedError(err.Error())
	}

	return nil
}

func (s ShippingRepositoryImpl) FindShippingAddressByIdFromDB(ShippingAddressID string) (*models.ShippingAddress,*apperrors.AppError){
	fmt.Println("Inside repo",ShippingAddressID)
	shipping := &models.ShippingAddress{}

	query := &dynamodb.GetItemInput{
		TableName: aws.String(shippingCollection),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(ShippingAddressID),
			},
		},
	}

	result, err := s.shippingDB.GetItem(query)
	if err != nil {
		logger.Info(result)
		logger.Error("Failed to get item from database - " + err.Error())
		return nil ,  apperrors.NewUnexpectedError(err.Error())
	}

	if result.Item == nil {
		logger.Error("Categories for given ID doesn't exists - ")
		err_ := apperrors.NewNotFoundError("Categories for given ID doesn't exists")
		return nil, err_
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, shipping)
	if err != nil {
		logger.Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, apperrors.NewUnexpectedError(err.Error())
	}
	return shipping, nil
}
func (s ShippingRepositoryImpl) UpdateShippingAddressByIdFromDB(id string,shippingAddress *models.ShippingAddress) (bool, *apperrors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//t:=time.Now().Format("2006-01-02 15:04:05")
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				S: aws.String(shippingAddress.FirstName),
			}, ":s1": {
				S: aws.String(shippingAddress.LastName),
			}, ":s2": {
				S: aws.String(shippingAddress.City),
			}, ":s3": {
				S: aws.String(shippingAddress.Address1),
			}, ":s4": {
				S: aws.String(shippingAddress.Address2),
			}, ":s5": {
				N: aws.String(strconv.Itoa(shippingAddress.CountryID)),
			}, ":s6": {
				N: aws.String(strconv.Itoa(shippingAddress.PostCode)),
			},
			//  ":s7": {
			// 	S: aws.String((t)),
			// },
		},
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set firstname =:s, lastname = :s1, city = :s2, address_1 = :s3, address_2 = :s4, country_id = :s5, postcode =:s6"),
		TableName:        aws.String("ShippingAddress"),
	}
fmt.Println("input\n",input)
	_, err := s.shippingDB.UpdateItemWithContext(ctx, input)
	if err != nil {
		return false, &apperrors.AppError{Message: fmt.Sprintf("unable to update - %s", err.Error())}
	}
	return true, nil
}
func (s ShippingRepositoryImpl) DeleteShippingAddressByIdFromDB(id string) (bool, *apperrors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String("ShippingAddress"),
	}

	_, err := s.shippingDB.DeleteItemWithContext(ctx, input)
	if err != nil {
		return false, &apperrors.AppError{Message: fmt.Sprintf("unable to delete- %s", err.Error())}
	}
	return true, nil
}









