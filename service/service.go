package service

import (
	"kafka-http-example/config"
	"kafka-http-example/consumer"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func Run(cfg *config.Config) {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level = cfg.LogLevel
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	c, err := consumer.NewKafkaConsumer(cfg.Topics, cfg.ConsumerConfig, logger)
	if err != nil {
		logger.Fatal("Failed to create Kafka consumer", zap.Error(err))
	}
	defer c.Stop()
	c.Consume()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}
