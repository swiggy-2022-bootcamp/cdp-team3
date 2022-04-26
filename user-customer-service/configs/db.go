package configs

// import (
// 	"fmt"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/credentials"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/dynamodb"
// 	"go.uber.org/zap"
// )

// var DB *dynamodb.DynamoDB

// func ConnectDB() *dynamodb.DynamoDB {
// 	//initialize client
// 	sess, _ := session.NewSession(&aws.Config{
// 		Region: aws.String(EnvRegion()),
// 		// Endpoint:    aws.String("http://localhost:8000"),
// 		Credentials: credentials.NewStaticCredentials(EnvAccessKey(), EnvSecretKey(), ""),
// 	})
// 	DB = dynamodb.New(sess)

// 	//ping the database
// 	_, err := DB.ListTables(&dynamodb.ListTablesInput{})
// 	if err != nil {
// 		zap.L().Error("Connection to dynamoDB failed.")
// 	}
// 	zap.L().Info("Connected to dynamoDB")
// 	return DB
// }

// func CreateTable(DB *dynamodb.DynamoDB) error {

// 	zap.L().Info("Inside create table")
// 	input := &dynamodb.CreateTableInput{
// 		AttributeDefinitions: []*dynamodb.AttributeDefinition{
// 			{
// 				AttributeName: aws.String("rewardId"),
// 				AttributeType: aws.String("S"),
// 			},
// 		},
// 		KeySchema: []*dynamodb.KeySchemaElement{
// 			{
// 				AttributeName: aws.String("rewardId"),
// 				KeyType:       aws.String("HASH"),
// 			},
// 		},
// 		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
// 			ReadCapacityUnits:  aws.Int64(10),
// 			WriteCapacityUnits: aws.Int64(10),
// 		},
// 		TableName: aws.String("Customers"),
// 	}
// 	response, err := DB.CreateTable(input)
// 	if err != nil {
// 		fmt.Println("err", err)
// 		zap.L().Error("Error Creating Table:" + err.Error())
// 		return err
// 	}
// 	fmt.Println("response", response)

// 	zap.L().Info("Created the table" + response.String())
// 	return nil
// }
