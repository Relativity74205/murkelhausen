package main

import (
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
)
import "github.com/spf13/viper"

// TODO move to config
const queueChannelSize = 100

type mqttConfig struct {
	broker       string
	clientId     string
	cleanSession bool
}

type kafkaConfig struct {
	broker string
}

type gohausenConfig struct {
	mqtt  mqttConfig
	kafka kafkaConfig
}

var Conf gohausenConfig

func main() {
	setupConfig()
	setupLogger()
	log.Info("Starting")
	var messageQueue = make(chan ChannelPayload, queueChannelSize)

	// TODO start also kafkaProducer as go routine and end main function when all go routines close.
	// TODO go routines shall close on system call
	go dispatcher(messageQueue)
	go mqttConsumer(messageQueue)
	kafkaProducer(messageQueue)

	//log.Info("Started everything")
}

func setupLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		ForceColors:            true,
		DisableLevelTruncation: false,
		PadLevelText:           true,
	})
	log.SetLevel(log.DebugLevel) // TODO config
	log.SetReportCaller(false)
}

func setupConfig() {
	viper.SetConfigName("config")
	viper.SetConfigName("yaml")
	viper.AddConfigPath("$HOME/.gohausen")
	viper.AddConfigPath(".")
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.WithField("e.Name", e.Name).Warning("Config file changed!")
	})
	viper.WatchConfig()

	err := viper.ReadInConfig()

	Conf = gohausenConfig{
		mqtt: mqttConfig{
			broker:       viper.GetString("mqtt.broker"),
			clientId:     viper.GetString("mqtt.clientId"),
			cleanSession: viper.GetBool("mqtt.cleanSession"),
		},
		kafka: kafkaConfig{
			broker: viper.GetString("kafka.broker"),
		},
	}

	if err != nil { // Handle errors reading the config file
		log.WithField("err", err).Fatal("fatal error config file:")
	}
}
