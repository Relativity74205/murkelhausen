package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	setupConfig()
	setupLogger()
	startModules()
}

func startModules() {
	log.Info("Creating messageQueue...")

	var messageQueue = make(chan ChannelPayload, Conf.app.queueChannelSize)

	log.Info("Starting modules...")
	// TODO start also kafkaProducer as go routine and end main function when all go routines close.
	// TODO go routines shall close on system call
	//go dispatcher(messageQueue)
	//go mqttConsumer(messageQueue)

	osSignalChannel := make(chan os.Signal, 1)
	signal.Notify(osSignalChannel, os.Interrupt, syscall.SIGTERM)

	if inModulesToLoad("kafkaProducer") {
		go kafkaProducer(messageQueue)
	}
	if inModulesToLoad("scheduler") {
		go gohausenScheduler(messageQueue, osSignalChannel)
	}

	log.Info("Started everything")
	<-osSignalChannel
	close(osSignalChannel)
	close(messageQueue)
	log.Info("Ending everything ...")
	time.Sleep(1 * time.Second) // TODO move to config
	log.Info("Sleep complete. Goodbye!")
}

func inModulesToLoad(moduleName string) bool {
	for _, entry := range Conf.app.modules {
		if entry == moduleName {
			log.Infof("Module '%s' is to be loaded.", moduleName)
			return true
		}
	}

	return false
}
