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
	
	return client
}

func CreateTable(DB *dynamodb.DynamoDB) error {

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("category_id"),
				AttributeType: aws.String("S"),
			},
		
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("category_id"),
				KeyType:       aws.String("HASH"),
			},
		
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("Categories"),
	}
	response, err := DB.CreateTable(input)
	if err != nil {
	
		logger.Error("Got error calling CreateTable: %s", err)
		return err
	}


	logger.Info("Created the table" + response.String())
	return nil
}