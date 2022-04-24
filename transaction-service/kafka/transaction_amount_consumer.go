package kafka

import (
	"context"
	"strconv"
	"strings"

	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/domain/repository"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/domain/services"
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/models"
	"go.uber.org/zap"
)

func AddTransactionAmountConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:" + configs.EnvAddTransactionAmountBrokerAddress()},
		Topic:   configs.EnvAddTransactionAmountTopic(),
	})

	for {

		m, err := r.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}

		strArr := strings.Split(string(m.Value), " ")

		if len(strArr) == 2 {
			customerId := strArr[0]
			transactionAmount := strArr[1]

			if s, err := strconv.ParseFloat(transactionAmount, 64); err == nil {
				newTransaction := &models.Transaction{
					TransactionId: uuid.New().String(),
					Amount:        s,
					Description:   "Transaction Amount Added from Orders Service through KafKa topic(add_transaction_amount)",
					CustomerID:    customerId,
				}

				transactionRepository := repository.NewTransactionRepositoryImpl(configs.DB)
				transactionService := services.NewTransactionServiceImpl(transactionRepository)
				_, err := transactionService.AddTransactionAmtToCustomer(newTransaction)

				if err == nil {
					zap.L().Info("Successfully added transaction to customer from Orders Service" + customerId)
				} else {
					zap.L().Error("Error adding transaction points through kafka" + err.Message)
				}
			} else {
				zap.L().Error("Error adding transaction amount through Kafka topic(add_transaction_amount) due to invalid transaction amount")
			}
		} else {
			zap.L().Error("Error consuming msgs from Kafka topic(add_transaction_amount) due to invalid msg format - " + string(m.Value))
		}
	}
}
