package kafka

import (
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/configs"
)

func  AddTransactionAmountProducer(customerId string, transactionAmount float64 ) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:"+configs.EnvAddTransactionAmountBrokerAddress()},
		Topic:   configs.EnvAddTransactionAmountTopic(),
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	msg := customerId+" "+fmt.Sprintf("%f", transactionAmount)
	
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("customerId"),
			Value: []byte(msg),
		},
	)

	if err != nil {
		fmt.Print("Error Writing Kafka Msg to add_transaction_amount topic")
	}
}