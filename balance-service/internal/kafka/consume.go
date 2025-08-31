package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/egon89/fc-event-driven-arch/internal/usecase"
	"github.com/segmentio/kafka-go"
)

type balanceUpdatedEvent struct {
	Name    string
	Payload saveBalanceConsumerInputDto
}

type saveBalanceConsumerInputDto struct {
	AccountIdFrom      string  `json:"account_id_from"`
	BalanceAccountFrom float64 `json:"balance_account_id_from"`
	AccountIdTo        string  `json:"account_id_to"`
	BalanceAccountTo   float64 `json:"balance_account_id_to"`
}

type balanceConsumer struct {
	Reader  *kafka.Reader
	UseCase *usecase.SaveBalanceUseCase
}

func NewBalanceConsumer(broker, topic string, useCase *usecase.SaveBalanceUseCase) *balanceConsumer {
	log.Printf("initializing kafka at %s listening on topic %s\n", broker, topic)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "balance-group",
	})

	return &balanceConsumer{
		Reader:  reader,
		UseCase: useCase,
	}
}

func (bc *balanceConsumer) Consume(ctx context.Context) {
	log.Printf("kafka consumer listening on topic %s\n", bc.Reader.Config().Topic)

	for {
		message, err := bc.Reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("error reading message: %v\n", err)
			continue
		}

		log.Printf("received: %s = %s\n", string(message.Key), string(message.Value))

		var input balanceUpdatedEvent
		err = json.Unmarshal(message.Value, &input)
		if err != nil {
			log.Printf("error unmarshalling message: %v", err)
			continue
		}

		log.Printf("processing input: %+v\n", input)

		err = bc.UseCase.Execute(ctx, usecase.SaveBalanceInputDto{
			AccountIdFrom:      input.Payload.AccountIdFrom,
			BalanceAccountFrom: input.Payload.BalanceAccountFrom,
			AccountIdTo:        input.Payload.AccountIdTo,
			BalanceAccountTo:   input.Payload.BalanceAccountTo,
		})
		if err != nil {
			log.Printf("error executing use case: %v\n", err)
			continue
		}
	}
}
