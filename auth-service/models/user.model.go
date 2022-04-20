package models

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/configs"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

var UserTableName = "users"

func CreateTable() {
	tableInput := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("email"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("email"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String(UserTableName),
	}

	_, err := configs.DB.CreateTable(tableInput)
	if err != nil {
		fmt.Println("Error creating table:", err)
	}

}

func CreateUser(user User) (User, error) {
	result, err := configs.DB.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(user.Id),
			},
			"name": {
				S: aws.String(user.Name),
			},
			"email": {
				S: aws.String(user.Email),
			},
			"password": {
				S: aws.String(user.Password),
			},
			"is_admin": {
				BOOL: aws.Bool(user.IsAdmin),
			},
		},
		TableName: aws.String(UserTableName),
	})

	fmt.Println(result)

	if err != nil {
		return User{}, err
	}
	return user, nil
}

func main() {
	CreateTable()
}
