package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/models"
)

type PaymentRepositoryImpl struct {
	paymentDB *dynamodb.DynamoDB
}

func NewPaymentRepository(paymentDB *dynamodb.DynamoDB) PaymentRepositoryImpl {
	return PaymentRepositoryImpl{paymentDB: paymentDB}
}

func (pr PaymentRepositoryImpl) CreatePayment(payment models.Payment) (string, *errors.AppError) {
	_, err := pr.paymentDB.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(models.PaymentTableName),
		Item: map[string]*dynamodb.AttributeValue{
			"transaction_id": {
				S: aws.String(payment.TransactionId),
			},
			"order_id": {
				S: aws.String(payment.OrderId),
			},
			"user_id": {
				S: aws.String(payment.UserId),
			},
			"amount": {
				N: aws.String(fmt.Sprint(payment.Amount)),
			},
			"details": {
				M: map[string]*dynamodb.AttributeValue{
					"payment_mode": {
						S: &payment.Details.PaymentMode,
					},
					"agree": {
						BOOL: &payment.Details.Agree,
					},
					"message": {
						S: &payment.Details.Message,
					},
				},
			},
			"status": {
				S: aws.String(payment.Status),
			},
			"created_at": {
				S: aws.String(payment.CreatedAt.String()),
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return "", errors.NewInternalServerError(err.Error())
	}
	return payment.TransactionId, nil
}
