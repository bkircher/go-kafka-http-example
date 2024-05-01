package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

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

	sigs := make(chan os.Signal, 1)
	stop := make(chan struct{}, 1) // Channel to stop reading messages

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		logger.Info("Received signal", zap.String("signal", sig.String()))
		close(stop)
	}()

	config := kafka.ConfigMap{
		"bootstrap.servers":  "localhost:9092",
		"group.id":           "my-consumer-group",
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	}

	// Create Consumer instance
	c, err := kafka.NewConsumer(&config)
	if err != nil {
		logger.Fatal("Failed to create consumer", zap.Error(err))
	}
	defer c.Close()

	// Subscribe to topic
	err = c.SubscribeTopics([]string{"my-topic"}, nil)
	if err != nil {
		logger.Fatal("Failed to subscribe to topics", zap.Error(err))
	}

	// Process messages
	for {
		select {
		case <-stop:
			logger.Info("Stopping consumer")
			return
		default:
			msg, err := c.ReadMessage(time.Second)
			if err == nil {
				logger.Debug("Received message",
					zap.String("topic", *msg.TopicPartition.Topic),
					zap.ByteString("key", msg.Key),
					zap.ByteString("value", msg.Value),
					zap.Int32("partition", msg.TopicPartition.Partition),
					zap.Int64("offset", int64(msg.TopicPartition.Offset)))
			} else {
				logger.Error("Consumer error",
					zap.Error(err))
			}
		}
	}
}
