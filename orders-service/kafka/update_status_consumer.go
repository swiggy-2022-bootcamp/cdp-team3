package kafka

import (
	"context"
	"strings"

	kafka "github.com/segmentio/kafka-go"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/orders-service/domain/repository"
	"go.uber.org/zap"
)

func UpdateOrderStatusConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:"+configs.EnvUpdateStatusBrokerAddress()},
		Topic:   configs.EnvUpdateStatusTopic(),
	})

	for {

		m, err := r.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}

		strArr := strings.Split(string(m.Value), " ")
		
		if len(strArr) == 2 {
			orderId := strArr[0]
			orderStatus := strArr[1]

			if orderStatus == "COMPLETED" || orderStatus == "FAILED" {

				orderRepository := repository.NewOrderRepositoryImpl(configs.DB)
				_, err := orderRepository.UpdateStatusByIdInDB(orderId, orderStatus)
				if err == nil {
					zap.L().Info("Updated order status to "+orderStatus+" for order- "+orderId+" successfully through kafka topic(update_status)")	
				} else {
					zap.L().Error("Error updating status through kafka in DB for order - "+orderId+" "+err.Message)
				}

			} else {
				zap.L().Error("Error updating order status through Kafka topic topic due to invalid status")
			}
		} else {
			zap.L().Error("Error consuming msgs from Kafka topic(update_status) due to invalid msg format - "+string(m.Value))
		}
	}
}