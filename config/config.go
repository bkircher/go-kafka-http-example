package config

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
)

type Config struct {
	ConsumerConfig *kafka.ConfigMap
	Topics         []string
	LogLevel       zap.AtomicLevel
}

func New() (*Config, error) {
	consumerCfg := &kafka.ConfigMap{
		"bootstrap.servers":  "localhost:9092",
		"group.id":           "my-consumer-group",
		"auto.offset.reset":  "latest",
		"enable.auto.commit": false,
	}

	logLevel := zap.NewAtomicLevel()
	if err := logLevel.UnmarshalText([]byte(getEnv("LOG_LEVEL", "info"))); err != nil {
		return nil, err
	}

	return &Config{
		ConsumerConfig: consumerCfg,
		Topics:         []string{"my-topic"},
		LogLevel:       logLevel,
	}, nil
}

// getEnv fetches environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
