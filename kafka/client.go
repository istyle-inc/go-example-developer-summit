package kafka

import (
	"go-example-developers-summit/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaClient struct {
	BootstrapServers *string
	Logger           logger.Loggable
	KafkaProducer    ProducerInterface
	ClientConnector  *kafka.Producer
}

type ProducerInterface interface {
	ProduceMessages(msgs []string)
}
