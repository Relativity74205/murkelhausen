package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const kafkaServer = "192.168.1.69:19092"
const topicConst string = "test_topic"

func handleProducerEvents(producer *kafka.Producer) {
	for producerEvent := range producer.Events() {
		switch kafkaEvent := producerEvent.(type) {
		case *kafka.Message:
			if kafkaEvent.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", kafkaEvent.TopicPartition)
			} else {
				fmt.Printf("Delivered message to %v\n", kafkaEvent.TopicPartition)
			}
		}
	}
}

func Producer(queueChannel chan string) {
	topic := topicConst
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServer})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go handleProducerEvents(p)

	// Produce messages to topic (asynchronously)
	for msg := range queueChannel {
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(msg),
		}, nil)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("queueChannel was closed. Flushing...")

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)

	fmt.Println("queueChannel was closed. Flushing complete.")
}
