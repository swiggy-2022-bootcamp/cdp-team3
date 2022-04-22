package models

type Reward struct {
	RewardId   string `json:"rewardId" dynamodbav:"rewardId" validate:"required"`
	CustomerId string `json:"customerId" dynamodbav:"customerId" validate:"required"`
	Rewards    int32  `json:"rewards" dynamodbav:"rewards" validate:"required"`
}

type SwaggerReward struct {
	CustomerId string `json:"customerId" dynamodbav:"customerId" validate:"required"`
	Rewards    int32  `json:"rewards" dynamodbav:"rewards" validate:"required"`
}
