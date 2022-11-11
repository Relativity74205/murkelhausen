package main

import (
	"encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var queueChannelReference chan ChannelPayload

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
	mqttTopic := message.Topic()
	msgPayload := message.Payload()
	err := json.Unmarshal(msgPayload, &data)
	if err != nil {
		log.WithField("error", err).Error("Error with unmarshalling message payload.")
	}

	channelPayload := ChannelPayload{
		Topic: mqttTopic,
		Value: data,
	}
	queueChannelReference <- channelPayload

	log.WithFields(log.Fields{
		"topic": channelPayload.Topic,
		"key":   channelPayload.Key,
		"value": channelPayload.Value,
	}).Info("Received MQTT message from Broker and sent to queueChannelReference.")
}

func onMessageReceivedXiaomiMiSensor(_ mqtt.Client, message mqtt.Message) {
	var data XiaomiMiSensorData
	mqttTopic := message.Topic()
	msgPayload := message.Payload()
	err := json.Unmarshal(msgPayload, &data)
	if err != nil {
		log.WithField("error", err).Error("Error with unmarshalling message payload.")
	}

	channelPayload := ChannelPayload{
		Topic: xiaomiMiSensorKafkaTopic,
		Key:   mqttTopic,
		Value: data,
	}
	queueChannelReference <- channelPayload

	log.WithFields(log.Fields{
		"topic": channelPayload.Topic,
		"key":   channelPayload.Key,
		"value": channelPayload.Value,
	}).Info("Received MQTT message from Broker and sent to queueChannelReference.")
}

func mqttConsumer(queueChannel chan ChannelPayload) {
	queueChannelReference = queueChannel
	osSignalChannel := make(chan os.Signal, 1)
	signal.Notify(osSignalChannel, os.Interrupt, syscall.SIGTERM)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(id)
	opts.SetCleanSession(cleanSession)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.WithField("error", token.Error()).Fatal("Could not connect to Mqtt server.")
	}

	if token := client.SubscribeMultiple(topicsTest, onMessageReceivedTest); token.Wait() && token.Error() != nil {
		log.WithFields(log.Fields{
			"error": token.Error(),
			"topic": topicsTest,
		}).Fatal("Subscription to topic failed")
	}
	if token := client.SubscribeMultiple(topicsXiaomiMiSensor, onMessageReceivedXiaomiMiSensor); token.Wait() && token.Error() != nil {
		log.WithFields(log.Fields{
			"error": token.Error(),
			"topic": topicsXiaomiMiSensor,
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
