package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Conf gohausenConfig

type appConfig struct {
	queueChannelSize int
	modules          []string
}

type mqttConfig struct {
	broker       string
	clientId     string
	cleanSession bool
	qos          int
}

type kafkaConfig struct {
	broker            string
	schemaRegistryUrl string
}

type dispatcherConfig struct {
	port               int
	shellyHTKafkaTopic string
	shellyFloodTopic   string
}

type mqttKafkaMapping struct {
	qos         int
	mqttTopics  []string
	kafkaTopic  string
	payloadType string
}

type gohausenConfig struct {
	app               appConfig
	mqtt              mqttConfig
	kafka             kafkaConfig
	dispatcher        dispatcherConfig
	mqttKafkaMappings []mqttKafkaMapping
}

func setupConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.app")
	viper.AddConfigPath(".")
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.WithField("e.Name", e.Name).Warning("Config file changed!")
	})
	viper.WatchConfig()

	err := viper.ReadInConfig()

	var mappings []mqttKafkaMapping
	for k := range viper.GetStringMap("mappingMqttKafka") {
		subTreeName := fmt.Sprintf("mappingMqttKafka.%s", k)
		subTree := viper.Sub(subTreeName)

		mappings = append(mappings, mqttKafkaMapping{
			qos:         subTree.GetInt("qos"),
			mqttTopics:  subTree.GetStringSlice("mqttTopics"),
			kafkaTopic:  subTree.GetString("kafkaTopic"),
			payloadType: subTree.GetString("payloadType"),
		})
	}
	Conf = gohausenConfig{
		app: appConfig{
			queueChannelSize: viper.GetInt("app.queueChannelSize"),
			modules:          viper.GetStringSlice("app.modules"),
		},
		mqtt: mqttConfig{
			broker:       viper.GetString("mqtt.broker"),
			clientId:     viper.GetString("mqtt.clientId"),
			cleanSession: viper.GetBool("mqtt.cleanSession"),
			qos:          viper.GetInt("mqtt.qos"),
		},
		kafka: kafkaConfig{
			broker:            viper.GetString("kafka.broker"),
			schemaRegistryUrl: viper.GetString("kafka.schemaRegistryUrl"),
		},
		dispatcher: dispatcherConfig{
			port:               viper.GetInt("dispatcher.port"),
			shellyHTKafkaTopic: viper.GetString("dispatcher.shellyHTKafkaTopic"),
			shellyFloodTopic:   viper.GetString("dispatcher.shellyFloodTopic"),
		},
		mqttKafkaMappings: mappings,
	}

	if err != nil { // Handle errors reading the config file
		log.WithField("err", err).Fatal("fatal error config file:")
	}
}
