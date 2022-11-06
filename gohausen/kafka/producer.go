package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/avro"
	"gohausen/dto"
	"os"
)

const kafkaServer = "192.168.1.69:19092"
const schemaRegistryUrl = "http://192.168.1.69:8081"

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

func Producer(queueChannel chan dto.ChannelPayload) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServer})
	if err != nil {
		panic(err)
	}

	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(schemaRegistryUrl))
	if err != nil {
		fmt.Printf("Failed to create schema registry client: %s\n", err)
		os.Exit(1)
	}

	ser, err := avro.NewGenericSerializer(client, serde.ValueSerde, avro.NewSerializerConfig())
	if err != nil {
		fmt.Printf("Failed to create serializer: %s\n", err)
		os.Exit(1)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go handleProducerEvents(p)

	// Produce messages to topic (asynchronously)
	for msg := range queueChannel {
		topic := msg.Topic
		key := msg.Key
		value := msg.Value

		if err != nil {
			fmt.Println(err)
		}
		valueSerialized, err := ser.Serialize(topic, &value)
		if err != nil {
			fmt.Println(err)
		}

		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          valueSerialized,
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
