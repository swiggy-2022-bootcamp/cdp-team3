package kafka

import (
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/configs"
)

func UserCreationProducer(customerId string) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{configs.EnvAddUserBrokerAddress()},
		Topic:    configs.EnvAddUserAdditionTopic(),
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	msg := customerId + " " + "User Added"

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("customerId"),
			Value: []byte(msg),
		},
	)

	if err != nil {
		fmt.Print("Error Writing Kafka Msg to 'user.created.topic' topic", err)
	}
}

func UserDeletionProducer(customerId string) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:" + configs.EnvAddUserBrokerAddress()},
		Topic:    configs.EnvAddUserDeletionTopic(),
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	msg := customerId + " " + "User Deleted"

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("customerId"),
			Value: []byte(msg),
		},
	)

	if err != nil {
		fmt.Print("Error Writing Kafka Msg to 'user.deleted.topic' topic")
	}
}
