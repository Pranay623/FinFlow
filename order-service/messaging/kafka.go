package messaging

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

var KafkaWriter *kafka.Writer

func InitKafka() {
	broker := os.Getenv("KAFKA_BROKERS")
	if broker == "" {
		broker = "localhost:9092"
	}

	KafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    "order_events",
		Balancer: &kafka.LeastBytes{},
	}
}

func PublishOrderEvent(event interface{}) error {
	msgBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = KafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Value: msgBytes,
		},
	)
	if err != nil {
		log.Printf("Could not write message to Kafka: %v", err)
		return err
	}

	return nil
}
