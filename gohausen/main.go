package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
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
	fmt.Println("Disconnecting")
}

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func kafka_producer() {
	// kafkaURL := "localhost:9092"
	kafkaURL := "192.168.1.69:9092"
	topic := "test_topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaURL, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	// conn, err := kafka.Dial("tcp", kafkaURL)
	// if err != nil {
	// 	log.Fatal("failed to dial leader:", err)
	// }
	// defer conn.Close()

	// controller, err := conn.Controller()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// var controllerConn *kafka.Conn
	// controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer controllerConn.Close()

	// topicConfigs := []kafka.TopicConfig{
	// 	{
	// 		Topic:             topic,
	// 		NumPartitions:     1,
	// 		ReplicationFactor: 1,
	// 	},
	// }

	// err = controllerConn.CreateTopics(topicConfigs...)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	// _, err = conn.WriteMessages(
	// 	kafka.Message{Value: []byte("one!")},
	// 	kafka.Message{Value: []byte("two!")},
	// 	kafka.Message{Value: []byte("three!")},
	// )
	// if err != nil {
	// 	log.Fatal("failed to write messages:", err)
	// } else {
	// 	print("Written message.")
	// }

	// if err := conn.Close(); err != nil {
	// 	log.Fatal("failed to close writer:", err)
	// } else {
	// 	print("Closed connection")
	// }
	// kafkaWriter := getKafkaWriter(kafkaURL, topic)
	// defer kafkaWriter.Close()
	// key := "Key-1"
	// msg := kafka.Message{
	// 	Key:   []byte(key),
	// 	Value: []byte(fmt.Sprint(uuid.New())),
	// }
	// err := kafkaWriter.WriteMessages(context.Background(), msg)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("produced", key)
	// }
}

func kafka_producer2() {
	kafkaURL := "192.168.1.69:19092"
	topic := "test_topic"
	writer := newKafkaWriter(kafkaURL, topic)
	defer writer.Close()
	fmt.Println("start producing ... !!")
	for i := 0; ; i++ {
		key := fmt.Sprintf("Key-%d", i)
		msg := kafka.Message{
			Key:   []byte(key),
			Value: []byte(fmt.Sprint(uuid.New())),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("produced", key)
		}
		time.Sleep(1 * time.Second)
	}
}

func kafka_list_topics() {
	conn, err := kafka.Dial("tcp", "192.168.1.69:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}

func kafka_reader() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"192.168.1.69:9092"},
		Topic:     "mqtt",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(0)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func main() {
	kafka_producer2()
}
