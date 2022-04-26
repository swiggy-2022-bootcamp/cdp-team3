package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/models"
	"go.uber.org/zap"
)

var transactionCollection = "Transaction"

type TransactionRepositoryImpl struct {
	transactionDB *dynamodb.DynamoDB
}

func NewTransactionRepositoryImpl(transactionDB *dynamodb.DynamoDB) TransactionRepository {
	return &TransactionRepositoryImpl{
		transactionDB: transactionDB,
	}
}

func(tri TransactionRepositoryImpl) GetTransactionByCustomerIdInDB(customerId string) ([]models.Transaction, *errors.AppError) {
	filt := expression.Name("customer_id").Equal(expression.Value(customerId))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		zap.L().Error("Error constructing Expression")
		return nil, errors.NewUnexpectedError("Error constructing Expression -  " + err.Error())
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String(transactionCollection),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}

	res, err := configs.DB.Scan(input)

	if err != nil {
		zap.L().Error("Error Fetching data from DB")
		return nil, errors.NewUnexpectedError("Error Fetching data from DB -  " + err.Error())
	}

	fmt.Print(res)
	var transactions []models.Transaction

	if err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &transactions); err != nil {
		zap.L().Error("Error unMarshalling Transaction" + err.Error())
		return nil, errors.NewUnexpectedError("Error unMarshalling Transaction " + err.Error())
	}
	return transactions, nil
}

func(tri TransactionRepositoryImpl) AddTransactionAmtToCustomerInDB(transaction *models.Transaction) (*models.Transaction, *errors.AppError) {
	
	data, err := dynamodbattribute.MarshalMap(transaction)
	if err != nil {
		zap.L().Error("Marshalling of transaction failed - " + err.Error())
		return nil, errors.NewUnexpectedError("Marshalling of transaction failed - " + err.Error())
	}

	query := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(transactionCollection),
	}

	result, err := configs.DB.PutItem(query)
	if err != nil {
		zap.L().Error("Failed to add transaction - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to add transaction - " + err.Error())
	}
	fmt.Print(result)
	return transaction, nil
}