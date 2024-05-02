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

	cons, err := consumer.NewKafkaConsumer(cfg.Topics, cfg.ConsumerConfig, logger)
	if err != nil {
		logger.Fatal("Failed to create Kafka consumer", zap.Error(err))
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		logger.Debug("Shutdown signal received")
		cons.Stop()
	}()

	cons.Consume()
}
