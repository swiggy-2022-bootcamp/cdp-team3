package models

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/configs"
)

type Payment struct {
	TransactionId string               `json:"transactionId"`
	UserId        string               `json:"userId"`
	Amount        float64              `json:"amount"`
	OrderId       string               `json:"orderId"`
	Status        string               `json:"status"`
	Details       ModeOfPaymentDetails `json:"details"`
	CreatedAt     time.Time            `json:"createdAt"`
}

type ModeOfPaymentDetails struct {
	PaymentMode string `json:"paymentMode"`
	Agree       bool   `json:"agree"`
	Message     string `json:"message"`
}

const PaymentTableName = "Payments"

func CreatePaymentTable() {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("transaction_id"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("user_id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("transaction_id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String(PaymentTableName),
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("userid-index"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("user_id"),
						KeyType:       aws.String("HASH"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(1),
					WriteCapacityUnits: aws.Int64(1),
				},
			},
		},
	}

	_, err := configs.DB.CreateTable(input)
	if err != nil {
		fmt.Println(err.Error())
	}
}
