package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var queueChannelReference chan ChannelPayload

func getKafkaTopic(currentMqttTopic string) (string, error) {
	for _, mapping := range Conf.mqttKafkaMappings {
		for _, mqttTopic := range mapping.mqttTopics {
			if currentMqttTopic == mqttTopic {
				return mapping.kafkaTopic, nil
			}
		}
	}
	return "", errors.New("currentMqttTopic not found in config, should not happen")
}

func getMqttTopics() map[string]byte {
	mqttTopics := make(map[string]byte)
	for _, mapping := range Conf.mqttKafkaMappings {
		for _, topic := range mapping.mqttTopics {
			mqttTopics[topic] = byte(mapping.qos)
		}
	}

	log.WithField("mqttTopics", mqttTopics).Info("Connecting to the following topics")

	return mqttTopics
}

func getPayloadType(currentMqttTopic string) (string, error) {
	for _, mapping := range Conf.mqttKafkaMappings {
		for _, mqttTopic := range mapping.mqttTopics {
			if currentMqttTopic == mqttTopic {
				return mapping.payloadType, nil
			}
		}
	}
	return "", errors.New("currentMqttTopic not found in config, should not happen")
}

func unmarshalPayload(msgPayload []byte, data KafkaValue) {
	err := json.Unmarshal(msgPayload, &data)
	if err != nil {
		log.WithField("error", err).Error("Error with unmarshalling message payloadType.")
	}
}

func getKafkaValue(mqttTopic string, msgPayload []byte) KafkaValue {
	payloadType, _ := getPayloadType(mqttTopic)

	switch payloadType {
	case "MQTTTestData":
		var data MQTTTestData
		unmarshalPayload(msgPayload, &data)
		return data
	case "XiaomiMiSensorData":
		var data XiaomiMiSensorData
		unmarshalPayload(msgPayload, &data)
		data.SensorName = mqttTopic
		data.Tstamp = time.Now().Local()
		return data
	case "AqaraSensorData":
		var data AqaraSensorData
		unmarshalPayload(msgPayload, &data)
		data.SensorName = mqttTopic
		data.Tstamp = time.Now().Local()
		return data
	}

	return nil
}

func onMessageReceived(_ mqtt.Client, message mqtt.Message) {
	mqttTopic := message.Topic()
	msgPayload := message.Payload()

	kafkaTopic, err := getKafkaTopic(mqttTopic)
	if err != nil {
		log.WithField("error", err).Error("Could not find the assigned KafkaTopic to a MqttTopic.")
	}
	channelPayload := ChannelPayload{
		Topic: kafkaTopic,
		Key:   mqttTopic,
		Value: getKafkaValue(mqttTopic, msgPayload),
	}
	queueChannelReference <- channelPayload

	log.WithFields(log.Fields{
		"topic": channelPayload.Topic,
		"key":   channelPayload.Key,
		"value": fmt.Sprintf("%+v", channelPayload.Value),
	}).Info("Received MQTT message from Broker and sent to queueChannelReference.")
}

func mqttConsumer(queueChannel chan ChannelPayload) {
	queueChannelReference = queueChannel
	osSignalChannel := make(chan os.Signal, 1)
	signal.Notify(osSignalChannel, os.Interrupt, syscall.SIGTERM)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(Conf.mqtt.broker)
	opts.SetClientID(Conf.mqtt.clientId)
	opts.SetCleanSession(Conf.mqtt.cleanSession)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.WithField("error", token.Error()).Fatal("Could not connect to Mqtt server.")
	}

	if token := client.SubscribeMultiple(getMqttTopics(), onMessageReceived); token.Wait() && token.Error() != nil {
		log.WithFields(log.Fields{
			"error": token.Error(),
		}).Fatal("Subscription to topic failed")
	}

	<-osSignalChannel
	log.Info("Disconnecting mqtt client and closing queueChannel ...")
	client.Disconnect(250)
	close(queueChannelReference)
	log.Info("Mqtt client disconnected and queueChannel closed. Sleep for 1 second to give other routines chance to stop.")
	time.Sleep(1 * time.Second) // TODO move to config
	log.Info("Sleep complete. Goodbye!")
}
