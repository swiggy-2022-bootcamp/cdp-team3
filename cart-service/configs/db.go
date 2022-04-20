package configs

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// GetDynamoDBClient returns a DynamoDB client
func GetDynamoDBClient() *dynamodb.Client {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(EnvRegion()),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	return dynamodb.NewFromConfig(cfg)
}

// waitForTable waits for a table to be created
func waitForTable(ctx context.Context, tableName string, db *dynamodb.Client) error {
    w := dynamodb.NewTableExistsWaiter(db)
    err := w.Wait(ctx,
        &dynamodb.DescribeTableInput{
            TableName: aws.String(tableName),
        },
        2*time.Minute,
        func(o *dynamodb.TableExistsWaiterOptions) {
            o.MaxDelay = 5 * time.Second
            o.MinDelay = 5 * time.Second
        })
    if err != nil {
        return errors.Wrap(err, "Timed out while waiting for table to become active")
    }
    return err
}

// CreateDynamoDBTable creates a DynamoDB table
func CreateDynamoDBTable(tableName string, db *dynamodb.Client) error {
	ctx := context.TODO()
	// Create a table
	_, err := db.CreateTable(
		ctx,
		&dynamodb.CreateTableInput{
			TableName: &tableName,
			AttributeDefinitions: []types.AttributeDefinition{
				{
					AttributeName: aws.String("id"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("user_id"),
					AttributeType: types.ScalarAttributeTypeS,
				},
			},
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("id"),
					KeyType:      types.KeyTypeHash,
				},
				{
					AttributeName: aws.String("user_id"),
					KeyType:      types.KeyTypeRange,
				},
			},
			BillingMode: types.BillingModePayPerRequest,
		},
	)
	if err != nil {
		log.Fatalf("Unable to create table, %v", err)
		return err
	}
	log.Info("Waiting for created table to be active: ", tableName)
	return waitForTable(ctx, tableName, db)
}

// DeleteDynamoDBTable deletes a DynamoDB table
func DeleteDynamoDBTable(tableName string, db *dynamodb.Client) error {
	ctx := context.TODO()
	// Delete the table
	_, err := db.DeleteTable(
		ctx,
		&dynamodb.DeleteTableInput{
			TableName: &tableName,
		},
	)
	if err != nil {
		log.Fatalf("Unable to delete table, %v", err)
		return err
	}
	log.Info("Deleted Table: ", tableName)
	return nil
}
