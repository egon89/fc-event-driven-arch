package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func Consumer(broker, topic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "balance-group",
	})

	log.Printf("kafka consumer listening on topic %s\n", topic)

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error reading message: %v", err)
			continue
		}

		log.Printf("received: %s = %s\n", string(message.Key), string(message.Value))
	}
}
