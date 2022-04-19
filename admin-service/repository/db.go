package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var dbInitialized = false
var svc *dynamodb.DynamoDB

func GetDynamoDBInstance() *dynamodb.DynamoDB {
	if !dbInitialized {
		config := &aws.Config{
			Region:   aws.String("us-east-1"),
			Endpoint: aws.String("http://localhost:8000"),
		}

		sess := session.Must(session.NewSession(config))

		svc = dynamodb.New(sess)
		dbInitialized = true
	}
	_, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Println("Connection to dynamoDB failed.")
		return nil
	}
	fmt.Println("Connected to dynamoDB")
	return svc
}
