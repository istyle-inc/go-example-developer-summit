package kafka

import (
	. "go-example-developers-summit/config"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (kc *KafkaClient) PrepareProducer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":            *kc.BootstrapServers,
		"broker.address.family":        "v4",
		"queue.buffering.max.messages": "1000",
		"client.id":                    "testingClient",
	})
	if err != nil {
		panic(err)
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					kc.Logger.Info(fmt.Sprintf("Delivery failed: %v\n", ev.TopicPartition))
				} else {
					kc.Logger.Info(fmt.Sprintf("Delivered message to %v\n", ev.TopicPartition))
				}
			}
		}
	}()
	kc.KafkaProducer = kc
	kc.ClientConnector = p
}

func (kc *KafkaClient) ProduceMessages(msgs []string) {
	topic := KafkaTopic
	p := kc.ClientConnector
	for _, msg := range msgs {
		p.Produce(
			&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(msg),
			}, nil)
	}
	p.Flush(15 * 1000)
}
