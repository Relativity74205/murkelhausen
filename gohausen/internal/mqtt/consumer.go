package mqtt

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Relativity74205/murkelhausen/gohausen/internal/common"
	"github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var queueChannelReference chan common.ChannelPayload

func getKafkaTopic(currentMqttTopic string) (string, error) {
	for _, mapping := range common.Conf.MqttKafkaMappings {
		for _, mqttTopic := range mapping.MqttTopics {
			if currentMqttTopic == mqttTopic {
				return mapping.KafkaTopic, nil
			}
		}
	}
	return "", errors.New("currentMqttTopic not found in config, should not happen")
}

func getMqttTopics() map[string]byte {
	mqttTopics := make(map[string]byte)
	for _, mapping := range common.Conf.MqttKafkaMappings {
		for _, topic := range mapping.MqttTopics {
			if common.Conf.App.DebugMode && mapping.DebugMode { // only debug topics
				mqttTopics[topic] = byte(mapping.Qos)
			} else if !common.Conf.App.DebugMode && !mapping.DebugMode { // only productive topics
				mqttTopics[topic] = byte(mapping.Qos)
			}
		}
	}

	return mqttTopics
}

func getPayloadType(currentMqttTopic string) (string, error) {
	for _, mapping := range common.Conf.MqttKafkaMappings {
		for _, mqttTopic := range mapping.MqttTopics {
			if currentMqttTopic == mqttTopic {
				return mapping.PayloadType, nil
			}
		}
	}
	return "", errors.New("currentMqttTopic not found in config, should not happen")
}

func unmarshalPayload(msgPayload []byte, data common.KafkaValue) {
	err := json.Unmarshal(msgPayload, &data)
	if err != nil {
		log.WithField("error", err).Error("Error with unmarshalling message payloadType.")
	}
}

func getKafkaValue(mqttTopic string, msgPayload []byte) common.KafkaValue {
	payloadType, _ := getPayloadType(mqttTopic)

	switch payloadType {
	case "MQTTTestData":
		var data common.MQTTTestData
		unmarshalPayload(msgPayload, &data)
		return data
	case "XiaomiMiSensorData":
		var data common.XiaomiMiSensorData
		unmarshalPayload(msgPayload, &data)
		data.SensorName = mqttTopic
		data.Tstamp = time.Now().Local()
		return data
	case "AqaraSensorData":
		var data common.AqaraSensorData
		unmarshalPayload(msgPayload, &data)
		data.SensorName = mqttTopic
		data.Tstamp = time.Now().Local()
		return data
	case "PowerData":
		var data common.PowerData
		var rawData common.PowerDataRaw
		err := json.Unmarshal(msgPayload, &rawData)
		if err != nil {
			log.WithField("error", err).Error("Error with unmarshalling message payloadType.")
		}
		data.SensorName = mqttTopic
		location, _ := time.LoadLocation("Europe/Berlin")
		timeParsed, _ := time.ParseInLocation("2006-01-02T15:04:05", rawData.Time, location)
		if err != nil {
			log.WithField("error", err).Error("error parsing time for PowerData")
		}
		data.Tstamp = timeParsed
		data.PowerTotal = float64(rawData.Usage.Total)
		data.PowerCurrent = float64(rawData.Usage.Current)
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
	channelPayload := common.ChannelPayload{
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

func Start(messageQueue chan common.ChannelPayload, osSignalChannel chan os.Signal) {
	log.Info("Starting MQTT consumer...")
	queueChannelReference = messageQueue

	opts := mqtt.NewClientOptions()
	opts.AddBroker(common.Conf.Mqtt.Broker)
	clientId := common.Conf.Mqtt.ClientId
	if common.Conf.App.DebugMode {
		clientId = fmt.Sprintf("%s_debug", clientId)
	}
	opts.SetClientID(clientId)
	opts.SetCleanSession(common.Conf.Mqtt.CleanSession)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.WithField("error", token.Error()).Fatal("Could not connect to Mqtt server.")
	}
	log.Info("Connected to MQTT broker.")

	mqttTopics := getMqttTopics()
	if token := client.SubscribeMultiple(mqttTopics, onMessageReceived); token.Wait() && token.Error() != nil {
		log.WithFields(log.Fields{
			"error": token.Error(),
		}).Fatal("Subscription to topic failed")
	}
	log.WithField("mqttTopics", mqttTopics).Info("Subscribed to MQTT topics.")

	<-osSignalChannel
	log.Info("Disconnecting mqtt client  ...")
	client.Disconnect(250)
	log.Info("Mqtt client disconnected.")
}
