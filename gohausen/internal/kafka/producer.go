package kafka

import (
	"github.com/Relativity74205/murkelhausen/gohausen/internal/common"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/avro"
	log "github.com/sirupsen/logrus"
	"os"
)

func handleProducerEvents(producer *kafka.Producer) {
	for producerEvent := range producer.Events() {
		switch kafkaEvent := producerEvent.(type) {
		case *kafka.Message:
			if kafkaEvent.TopicPartition.Error != nil {
				log.WithField("topicPartition", kafkaEvent.TopicPartition).Error("Delivery failed to Kafka!")
			} else {
				log.WithField("topicPartition", kafkaEvent.TopicPartition).Info("Successfully delivered message to Kafka!")
			}
		}
	}
}

func Start(messageQueue chan common.ChannelPayload, _ chan os.Signal) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": common.Conf.Kafka.Broker})
	if err != nil {
		panic(err)
	}

	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(common.Conf.Kafka.SchemaRegistryUrl))
	if err != nil {
		log.WithField("error", err).Fatal("Failed to create schema registry client!")
	}

	ser, err := avro.NewGenericSerializer(client, serde.ValueSerde, avro.NewSerializerConfig())
	if err != nil {
		log.WithField("error", err).Fatal("Failed to create serializer!")
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go handleProducerEvents(p)

	// Produce messages to topic (asynchronously)
	for msg := range messageQueue {
		topic := msg.Topic
		key := msg.Key
		value := msg.Value

		valueSerialized, err := ser.Serialize(topic, &value)
		if err != nil {
			log.WithFields(log.Fields{"messageValue": value, "error": err}).Error("Payload value could not be serialized!")
		}

		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          valueSerialized,
		}, nil)
		if err != nil {
			log.WithField("error", err).Error("Failed to sent message.")
		}
	}

	log.Info("queueChannel was closed. Flushing...")

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000) // TODO move to config and research the meaning

	log.Info("queueChannel was closed. Flushing complete.")
}
