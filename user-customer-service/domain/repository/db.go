package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/swiggy-2022-bootcamp/cdp-team3/user-customer-service/configs"
)

var dbInitialized = false
var svc *dynamodb.DynamoDB

func GetDynamoDBInstance() *dynamodb.DynamoDB {
	if !dbInitialized {
		config := &aws.Config{
			Region: aws.String("us-east-1"),
			// Endpoint: aws.String("http://localhost:8000"),
			Credentials: credentials.NewStaticCredentials(configs.EnvAccessKey(), configs.EnvSecretKey(), ""),
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
