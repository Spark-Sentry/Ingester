package kafka

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

// Consumer represents a Kafka consumer using the Confluent Kafka Go library.
type Consumer struct {
	consumer *kafka.Consumer
	topic    string
}

// NewConsumer creates and initializes a new Kafka consumer.
func NewConsumer(brokerAddress, topic, groupID string) (*Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokerAddress,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %s", err)
	}

	return &Consumer{
		consumer: consumer,
		topic:    topic,
	}, nil
}

// Consume starts consuming messages on the configured topic.
func (c *Consumer) Consume(ctx context.Context, handleMessage func([]byte)) error {
	err := c.consumer.Subscribe(c.topic, nil)
	if err != nil {
		return fmt.Errorf("failed to subscribe to topic: %s", err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping consumer...")
			return nil
		default:
			msg, err := c.consumer.ReadMessage(-1)
			if err != nil {
				log.Printf("Error reading message: %s\n", err)
				continue
			}
			handleMessage(msg.Value)
		}
	}
}
