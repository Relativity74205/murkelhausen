package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var c chan string

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	c <- string(message.Payload())
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}

func listen_mqtt() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

	topic := "test_topic"
	broker := "tcp://192.168.1.69:1883"
	id := "gohausen_client"
	cleansess := false
	qos := 0

	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(id)
	opts.SetCleanSession(cleansess)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe(topic, byte(qos), onMessageReceived); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	<-sigchan
	client.Disconnect(250)
	close(c)
	fmt.Println("Disconnecting")
}

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

func kafka_producer(c chan string) {
	kafkaServer := "192.168.1.69:19092"
	topic := "test_topic"

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaServer})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go handleProducerEvents(p)

	// Produce messages to topic (asynchronously)
	for msg := range c {
		time.Sleep(5 * time.Second)
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(msg),
		}, nil)
		print("foo")
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

func main() {
	c = make(chan string, 3)
	go kafka_producer(c)
	listen_mqtt()
}
