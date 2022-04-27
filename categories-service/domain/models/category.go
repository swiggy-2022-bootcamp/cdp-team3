package models


type Category struct {
	CategoryId             string                    `json:"category_id" dynamodbav:"category_id" `
	CategoryDescription     []CategoryDescription                    `json:"category_description" dynamodbav:"category_description"validate:"required"`
}