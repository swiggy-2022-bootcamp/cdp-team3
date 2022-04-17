package configs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go.uber.org/zap"
)

var DB *dynamodb.DynamoDB;
func ConnectDB() *dynamodb.DynamoDB {
	//initialize client
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(EnvRegion()),
		Credentials: credentials.NewStaticCredentials(EnvAccessKey(), EnvSecretKey(), ""),
	})
	DB = dynamodb.New(sess)

	//ping the database
	_, err := DB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		zap.L().Error("Connection to dynamoDB failed.")
	}
	zap.L().Info("Connected to dynamoDB")
	return DB
}

func CreateTable(DB *dynamodb.DynamoDB) error {

	zap.L().Info("Inside create table")
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("transaction_id"),
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
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("Transaction"),
	}
	response, err := DB.CreateTable(input)
	if err != nil {
		fmt.Println("err",err)
		zap.L().Error("Error Creating Table:"+err.Error())
		return err
	}
	fmt.Println("response",response)

	zap.L().Info("Created the table" + response.String())
	return nil
}