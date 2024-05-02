package consumer

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
	ctx      context.Context
	cancel   context.CancelFunc
	logger   *zap.Logger
}

func NewKafkaConsumer(topics []string, cfg *kafka.ConfigMap, logger *zap.Logger) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(cfg)
	if err != nil {
		logger.Fatal("Failed to create consumer", zap.Error(err))
	}

	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		logger.Fatal("Failed to subscribe to topics", zap.Error(err))
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &KafkaConsumer{
		consumer: c,
		ctx:      ctx,
		cancel:   cancel,
		logger:   logger,
	}, nil
}

func (kc *KafkaConsumer) Consume() {
	for {
		select {
		case <-kc.ctx.Done():
			kc.logger.Debug("Stopping consumer…")
			if err := kc.consumer.Close(); err != nil {
				kc.logger.Error("Failed to close consumer", zap.Error(err))
			}
			return
		default:
			ev := kc.consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				kc.logger.Debug("Received message",
					zap.String("topic", *e.TopicPartition.Topic),
					zap.ByteString("key", e.Key),
					zap.ByteString("value", e.Value),
					zap.Int32("partition", e.TopicPartition.Partition),
					zap.Int64("offset", int64(e.TopicPartition.Offset)))

			case kafka.Error:
				kc.logger.Error("Consumer error",
					zap.Error(e))
			}
		}
	}
}

func (kc *KafkaConsumer) Stop() {
	kc.cancel()
}
