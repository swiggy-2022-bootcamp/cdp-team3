package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/models"
	"go.uber.org/zap"
)

var rewardsCollection = "Rewards"

type RewardRepositoryImpl struct {
	rewardDB *dynamodb.DynamoDB
}

func NewRewardRepositoryImpl(rewardDB *dynamodb.DynamoDB) RewardRepository {
	return &RewardRepositoryImpl{
		rewardDB: rewardDB,
	}
}

func (rr RewardRepositoryImpl) AddRewardToDB(reward *models.Reward) *errors.AppError {
	data, err := dynamodbattribute.MarshalMap(reward)
	if err != nil {
		zap.L().Error("Marshalling of reward failed - " + err.Error())
		return errors.NewUnexpectedError(err.Error())
	}

	query := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(rewardsCollection),
	}

	result, err := rr.rewardDB.PutItem(query)
	if err != nil {
		zap.L().Error("Failed to insert reward into database - " + err.Error())
		return errors.NewUnexpectedError(err.Error())
	}
	fmt.Println(result)
	return nil
}

func (rr RewardRepositoryImpl) GetAllRewardsFromDB() ([]models.Reward, *errors.AppError) {

	var rewardsList []models.Reward
	params := &dynamodb.ScanInput{
		TableName: aws.String(rewardsCollection),
	}

	err := rr.rewardDB.ScanPages(params, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		var rewards []models.Reward
		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &rewards)
		if err != nil {
			zap.L().Error("\nCould not unmarshal AWS data: err =" + err.Error())
			return true
		}
		rewardsList = append(rewardsList, rewards...)
		return true
	})

	if err != nil {
		return nil, errors.NewUnexpectedError("Error fetching data from DB " + err.Error())
	}
	return rewardsList, nil
}

func (rr RewardRepositoryImpl) GetRewardByIdFromDB(rewardId string) (*models.Reward, *errors.AppError) {
	reward := &models.Reward{}

	query := &dynamodb.GetItemInput{
		TableName: aws.String(rewardsCollection),
		Key: map[string]*dynamodb.AttributeValue{
			"rewardId": {
				S: aws.String(rewardId),
			},
		},
	}

	result, err := configs.DB.GetItem(query)

	if err != nil {
		zap.L().Error("Failed to get item from database - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to get item from database - " + err.Error())
	}

	if result.Item == nil {
		zap.L().Error("Reward for given ID doesn't exist - " + rewardId)
		return nil, errors.NewNotFoundError("Reward for given ID doesn't exist - " + rewardId)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &reward)
	if err != nil {
		zap.L().Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, errors.NewUnexpectedError("Failed to unmarshal document fetched from DB - " + err.Error())
	}

	return reward, nil
}

func (rr RewardRepositoryImpl) GetRewardsByCustomerIdFromDB(customerId string) ([]models.Reward, *errors.AppError) {
	filt := expression.Name("customerId").Equal(expression.Value(customerId))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		zap.L().Error("Error constructing Expression")
		return nil, errors.NewUnexpectedError("Error constructing Expression -  " + err.Error())
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String(rewardsCollection),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}

	res, err := configs.DB.Scan(input)

	if err != nil {
		zap.L().Error("Error Fetching data from DB")
		return nil, errors.NewUnexpectedError("Error Fetching data from DB -  " + err.Error())
	}

	var rewards []models.Reward

	if len(res.Items) == 0 {
		zap.L().Error("No rewards found for customer " + customerId)
		return nil, errors.NewNotFoundError("No rewards found for customer " + customerId)
	}

	if err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &rewards); err != nil {
		zap.L().Error("Error unMarshalling Reward" + err.Error())
		return nil, errors.NewUnexpectedError("Error unMarshalling Reward " + err.Error())
	}
	return rewards, nil
}
