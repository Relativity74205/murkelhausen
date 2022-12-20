package common

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetupConfig() {
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
			Qos:         subTree.GetInt("qos"),
			MqttTopics:  subTree.GetStringSlice("mqttTopics"),
			KafkaTopic:  subTree.GetString("kafkaTopic"),
			PayloadType: subTree.GetString("payloadType"),
			DebugMode:   subTree.GetBool("debugMode"),
		})
	}
	Conf = gohausenConfig{
		App: appConfig{
			DebugMode:        viper.GetBool("app.debugMode"),
			Modules:          viper.GetStringSlice("app.modules"),
			QueueChannelSize: viper.GetInt("app.queueChannelSize"),
		},
		Tasks: tasksConfig{
			Psutil: taskConfig{viper.GetInt("tasks.psutil.schedule")},
		},
		Mqtt: mqttConfig{
			Broker:       viper.GetString("mqtt.broker"),
			ClientId:     viper.GetString("mqtt.clientId"),
			CleanSession: viper.GetBool("mqtt.cleanSession"),
			Qos:          viper.GetInt("mqtt.qos"),
		},
		Kafka: kafkaConfig{
			Broker:            viper.GetString("kafka.broker"),
			SchemaRegistryUrl: viper.GetString("kafka.schemaRegistryUrl"),
		},
		Dispatcher: dispatcherConfig{
			Port:               viper.GetInt("dispatcher.port"),
			ShellyHTKafkaTopic: viper.GetString("dispatcher.shellyHTKafkaTopic"),
			ShellyFloodTopic:   viper.GetString("dispatcher.shellyFloodTopic"),
			TestTopic:          viper.GetString("dispatcher.testTopic"),
		},
		MqttKafkaMappings: mappings,
	}

	if err != nil { // Handle errors reading the config file
		log.WithField("err", err).Fatal("fatal error config file:")
	}
}
