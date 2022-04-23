package utils

import (
	"context"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"github.com/swiggy-ipp/cart-service/configs"
	"github.com/swiggy-ipp/cart-service/dto/requests"
	"github.com/swiggy-ipp/cart-service/services"
)

// KafkaCartConsumeService is a Kafka Consumer for Cart Service
type KafkaCartConsumeService struct {
	brokerAddress string
	topic         string
	cartService   services.CartService
}

// Create a new Kafka Cart Consumer
func NewKafkaCartConsumeService(
	brokerAddress string,
	topic string,
	cartService services.CartService,
) *KafkaCartConsumeService {
	return &KafkaCartConsumeService{
		brokerAddress: brokerAddress,
		topic:         topic,
		cartService:   cartService,
	}
}

// Consume the Kafka User IDs from Profile service.
// Create the cart of the user when user ID received on Created topic.
// Delete the cart of the user when user ID received on Deleted topic.
func (kc *KafkaCartConsumeService) KafkaUserIDConsume() {
	// Set up Background context for Kafka listener
	ctx := context.Background()
	l := log.New()
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kc.brokerAddress},
		Topic:   kc.topic,
		GroupID: "my-group",
		Logger:  l,
	})

	for {
		// The `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Error("Could not read Kafka Message: " + err.Error())
		}

		// After receiving the message, log its value and apply the callback
		log.Info(string(msg.Key) + ": " + string(msg.Value))

		if kc.topic == configs.EnvKafkaUserCreatedTopic() {
			// Create Cart by User ID
			err = kc.cartService.CreateCart(ctx, string(msg.Value))
			if err != nil {
				log.Error("Could not create cart: " + err.Error())
			}
		} else if kc.topic == configs.EnvKafkaUserDeletedTopic() {
			// Delete Cart by User ID
			err = kc.cartService.DeleteCart(ctx, requests.CartIDRequest{UserID: string(msg.Key)})
			if err != nil {
				log.Error("Could not delete cart: " + err.Error())
			}
		}
	}
}
