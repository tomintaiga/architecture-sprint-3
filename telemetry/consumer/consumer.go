package consumer

import (
	"context"
	"encoding/json"
	"log"
	swagger "telemetry/go"
	"telemetry/repository"

	"github.com/IBM/sarama"
)

type Consumer struct {
	consumer sarama.Consumer
	repo     repository.Repository
}

const (
	KAFKA_TOPIC = "temperature"
)

func NewConsumer(kafkaAddress string, repo repository.Repository) (*Consumer, error) {
	c, err := sarama.NewConsumer([]string{kafkaAddress}, nil)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		consumer: c,
		repo:     repo,
	}, nil
}

func (c *Consumer) Run(ctx context.Context) {
	partConsumer, err := c.consumer.ConsumePartition(KAFKA_TOPIC, 0, sarama.OffsetNewest)
	if err != nil {
		log.Panicf("Can't create cosumer: %v", err)
	}

	defer partConsumer.Close()

	for {
		select {
		case msg, ok := <-partConsumer.Messages():
			if !ok {
				log.Printf("Channel closed, exiting")
				return
			}

			var payload swagger.TelemetryRecord
			err := json.Unmarshal(msg.Value, &payload)
			if err != nil {
				log.Printf("Can't unmarshal message: %v", err)
				continue
			}

			err = c.repo.SetValue(ctx, payload)
			if err != nil {
				log.Printf("Can't store record: %v", err)
				continue
			}
		case <-ctx.Done():
			log.Printf("Got done from context")
			return
		}
	}
}
