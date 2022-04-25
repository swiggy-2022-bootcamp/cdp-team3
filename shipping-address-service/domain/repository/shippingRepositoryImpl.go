package repository

import (
	"fmt"
	"time"
	"context"
	"strconv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
//	"github.com/google/uuid"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
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
func (s ShippingRepositoryImpl) InsertShippingAddressToDB(shippingAddress *models.ShippingAddress) ( string,*apperrors.AppError) {

	fmt.Println("inside repo",shippingAddress)
	Id:=shippingAddress.Id
	fmt.Println(Id)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	av, err := dynamodbattribute.MarshalMap(shippingAddress)
	if err != nil {
		return  "",apperrors.NewUnexpectedError(err.Error())
	}
	fmt.Println("Inside repo",shippingAddress)
    
	fmt.Println(shippingAddress)
	fmt.Println(av)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(shippingCollection),
	}

	_, err = s.shippingDB.PutItemWithContext(ctx, input)

	if err != nil {
		return "",apperrors.NewUnexpectedError(err.Error())
	}

	return Id,nil
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
		logger.Error("Shipping Address for given ID doesn't exists - ")
		err_ := apperrors.NewNotFoundError("Shipping Address for given ID doesn't exists")
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
				N: aws.String(strconv.FormatUint(uint64(shippingAddress.CountryID), 10)),
			}, ":s6": {
				N: aws.String(strconv.FormatUint(uint64(shippingAddress.PostCode), 10)),
			},
			// ":s7": {
			// 	S: aws.String(shippingAddress.UserID),
			// },
			// ":s8": {
			// 	S: aws.String(shippingAddress.DefaultAddress),
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
func (s ShippingRepositoryImpl) HandleSetExistingShippingAddressToDefaultByIdToDB(id string) (bool, *apperrors.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filt := expression.Name("id").Equal(expression.Value(id))
	proj := expression.NamesList(expression.Name("id"), expression.Name("default_address"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
if err != nil {
	fmt.Println(err)
  //  log.Fatalf("Got error building expression: %s", err)
}
// Build the query input parameters
params := &dynamodb.ScanInput{
    ExpressionAttributeNames:  expr.Names(),
    ExpressionAttributeValues: expr.Values(),
    FilterExpression:          expr.Filter(),
    ProjectionExpression:      expr.Projection(),
    TableName:                 aws.String("ShippingAddress"),
}

// Make the DynamoDB Query API call
result, err := s.shippingDB.Scan(params)
if err != nil {
	fmt.Println("err",err)
   // log.Fatalf("Query API call failed: %s", err)
}
fmt.Println("result",result)
numItems := 0

for _, i := range result.Items {
    item := models.ShippingAddress{}

    err = dynamodbattribute.UnmarshalMap(i, &item)

    if err != nil {
		fmt.Println("err",err)
        //log.Fatalf("Got error unmarshalling: %s", err)
    }

  
    if item.DefaultAddress == "1" {
       
        numItems++

        fmt.Println("Title: ", item.DefaultAddress)
     
        fmt.Println()
    }
}
fmt.Println("numItemss",numItems)
if(numItems==0){


	query := &dynamodb.UpdateItemInput{
		TableName: aws.String("ShippingAddress"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":points": {
				S:  aws.String("1"),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set default_address = :points"),
	}
fmt.Println("input\n",query)
     res, err := s.shippingDB.UpdateItemWithContext(ctx, query)
	
	if err != nil {
		fmt.Println("err",err)
		return false, &apperrors.AppError{Message: fmt.Sprintf("unable to update - %s", err.Error())}
	}
	fmt.Println("res",res)
	return true, nil
}
return false,&apperrors.AppError{Message:"Default Address is already set"}
}
func (s ShippingRepositoryImpl) GetDefaultShippingAddressOfUserByIdFromDB(id string) (*models.ShippingAddress, *apperrors.AppError) {


	filt := expression.Name("user_id").Equal(expression.Value(id))
	proj := expression.NamesList(expression.Name("id"),expression.Name("firstname"),expression.Name("city"),expression.Name("user_id"),expression.Name("lastname"),expression.Name("address_1"),expression.Name("address_2"),expression.Name("country_id"),expression.Name("postcode"), expression.Name("default_address"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
if err != nil {
	fmt.Println(err)

}

params := &dynamodb.ScanInput{
    ExpressionAttributeNames:  expr.Names(),
    ExpressionAttributeValues: expr.Values(),
    FilterExpression:          expr.Filter(),
    ProjectionExpression:      expr.Projection(),
    TableName:                 aws.String("ShippingAddress"),
}

// Make the DynamoDB Query API call
result, err := s.shippingDB.Scan(params)
if err != nil {
	fmt.Println("err",err)
   // log.Fatalf("Query API call failed: %s", err)
}
fmt.Println("result",result)
numItems := 0

for _, i := range result.Items {
    item := models.ShippingAddress{}

    err = dynamodbattribute.UnmarshalMap(i, &item)

    if err != nil {
		fmt.Println("err",err)
        //log.Fatalf("Got error unmarshalling: %s", err)
    }

  
    if item.DefaultAddress == "1" && item.UserID == id {
       
		fmt.Println("Item found",item)
		shipping := &models.ShippingAddress{
	     Id:item.Id,
		FirstName: item.FirstName,
		LastName:  item.LastName,
		City:      item.City,
		Address1:  item.Address1,
		Address2:  item.Address2,
		CountryID: uint32(item.CountryID),
		PostCode:  uint32(item.PostCode),
		UserID:    item.UserID,
		DefaultAddress: item.DefaultAddress,
	
		}
		return shipping,nil
        
        numItems++

        fmt.Println("Title: ", item.DefaultAddress)
     
        fmt.Println()
    }
}
fmt.Println("numItemss",numItems)

return nil,&apperrors.AppError{Message:"No Default Address,Please set one"}
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









