package config

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/cdp-team3/categories-service/utils/logger"
)

func ConnectDB() *dynamodb.DynamoDB {
	//initialize client
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(EnvRegion()),
		Credentials: credentials.NewStaticCredentials(EnvAccessKey(), EnvSecretKey(), ""),
	})
	client := dynamodb.New(sess)

	//ping the database
	_, err := client.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		logger.Error("Connection to dynamoDB failed.")
	}
	fmt.Println("Connected to DynamoDB")
	return client
}

func CreateTable(DB *dynamodb.DynamoDB) error {
fmt.Println("Inside create table")
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("category_id"),
				AttributeType: aws.String("S"),
			},
			// {
			// 	AttributeName: aws.String("category_description"),
			// 	AttributeType: aws.String("L"),
			// },
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("category_id"),
				KeyType:       aws.String("HASH"),
			},
			// {
			// 	AttributeName:aws.String("category_description"),
			// 	KeyType:       aws.String("RANGE"),
			// },
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("Categories"),
	}
	response, err := DB.CreateTable(input)
	if err != nil {
		fmt.Println("err while creating",err)
		logger.Error("Got error calling CreateTable: %s", err)
		return err
	}
	fmt.Println("response",response)

	logger.Info("Created the table" + response.String())
	return nil
}