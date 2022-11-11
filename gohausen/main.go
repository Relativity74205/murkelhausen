package main

import log "github.com/sirupsen/logrus"

const queueChannelSize = 100

func main() {
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
