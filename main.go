package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
)

func main() {
	// Initialize Zap logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // Flushes buffer, if any

	// Create Consumer instance
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "my-consumer-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		logger.Fatal("Failed to create consumer", zap.Error(err))
	}

	// Subscribe to topic
	err = c.SubscribeTopics([]string{"my-topic"}, nil)
	if err != nil {
		logger.Fatal("Failed to subscribe to topics", zap.Error(err))
	}

	// Process messages
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			logger.Info("Received message",
				zap.String("topic", *msg.TopicPartition.Topic),
				zap.Binary("value", msg.Value),
				zap.Int32("partition", msg.TopicPartition.Partition),
				zap.Int64("offset", int64(msg.TopicPartition.Offset)))
		} else {
			// Log the error with structured logging
			logger.Error("Consumer error",
				zap.Error(err))
		}
	}

	// c.Close()
}
