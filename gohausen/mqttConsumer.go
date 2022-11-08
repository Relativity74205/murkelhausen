package main

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var messageQueue chan ChannelPayload

const qos = 0

var topicsTest = map[string]byte{
	"test_topic": byte(qos),
}
var topicsXiaomiMiSensor = map[string]byte{
	"zigbee2mqtt/XaomiTempCellarHobby":      byte(qos),
	"zigbee2mqtt/XaomiTempCellarVersorgung": byte(qos),
}

const broker = "tcp://192.168.1.69:1883"
const id = "gohausen_client"
const cleanSession = false
const xiaomiMiSensorKafkaTopic = "xiaomi_mi_sensor"

func onMessageReceivedTest(_ mqtt.Client, message mqtt.Message) {
	var data MQTTTestData
	msgPayload := message.Payload()
	err := json.Unmarshal(msgPayload, &data)
	if err != nil {
		fmt.Println(err)
	}

	messageQueue <- ChannelPayload{
		Topic: "test_topic",
		Value: data,
	}

	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}

func onMessageReceivedXiaomiMiSensor(_ mqtt.Client, message mqtt.Message) {
	var data XiaomiMiSensorData
	mqttTopic := message.Topic()
	msgPayload := message.Payload()
	err := json.Unmarshal(msgPayload, &data)
	if err != nil {
		fmt.Println(err)
	}

	messageQueue <- ChannelPayload{
		Topic: xiaomiMiSensorKafkaTopic,
		Key:   mqttTopic,
		Value: data,
	}

	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}

func mqttConsumer(queueChannel chan ChannelPayload) {
	messageQueue = queueChannel
	osSignalChannel := make(chan os.Signal, 1)
	signal.Notify(osSignalChannel, os.Interrupt, syscall.SIGTERM)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(id)
	opts.SetCleanSession(cleanSession)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.SubscribeMultiple(topicsTest, onMessageReceivedTest); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	if token := client.SubscribeMultiple(topicsXiaomiMiSensor, onMessageReceivedXiaomiMiSensor); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	<-osSignalChannel
	fmt.Println("Disconnecting mqtt client and closing queueChannel ...")
	client.Disconnect(250)
	close(messageQueue)
	fmt.Println("Mqtt client disconnected and queueChannel closed. Sleep for 1 second to give other routines chance to stop.")
	time.Sleep(1 * time.Second)
	fmt.Println("Sleep complete. Goodbye!")
}
