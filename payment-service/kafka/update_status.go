package kafka

import (
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
	"github.com/swiggy-2022-bootcamp/cdp-team3/payment-service/configs"
)

var (
	topic         = "update_status"
	brokerAddress = configs.EnvBrokerAddress()
)

func UpdateOrderStatusProducer(orderId string, orderStatus string) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	msg := orderId + " " + orderStatus

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("orderId"),
			Value: []byte(msg),
		},
	)

	if err != nil {
		fmt.Print("Error Writing Kafka Msg to update_status topic")
	}
}
