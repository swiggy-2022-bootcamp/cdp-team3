package kafka

import (
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
)

const (
	topic         = "update_status"
	brokerAddress = "localhost:9092"
)

func UpdateOrderStatusProducer(orderId string, orderStatus string ) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:   topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	msg := orderId+" "+orderStatus
	
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