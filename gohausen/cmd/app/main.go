package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

func main() {
	setupConfig()
	setupLogger()
	startModules()
}

var modulesMapping = map[string]interface{}{
	"scheduler":     gohausenScheduler,
	"kafkaProducer": kafkaProducer,
	"dispatcher":    dispatcher,
	"mqttConsumer":  mqttConsumer,
}

func startModules() {
	log.Info("Creating channels...")

	var messageQueue = make(chan ChannelPayload, Conf.app.queueChannelSize)
	osSignalChannel := make(chan os.Signal, 1)
	signal.Notify(osSignalChannel, os.Interrupt, syscall.SIGTERM)

	log.WithField("moduleList", Conf.app.modules).Info("Starting modulesMapping...")
	// TODO start also kafkaProducer as go routine and end main function when all go routines close.
	// TODO go routines shall close on system call
	//go dispatcher(messageQueue)
	//go mqttConsumer(messageQueue)
	for _, moduleName := range Conf.app.modules {
		moduleCallable, ok := modulesMapping[moduleName]
		if !ok {
			log.WithField("module", moduleName).Error("Module not found.")
		}

		log.WithField("module", moduleName).Info("Starting module.")
		f := reflect.ValueOf(moduleCallable)
		in := make([]reflect.Value, 2)
		in[0] = reflect.ValueOf(messageQueue)
		in[1] = reflect.ValueOf(osSignalChannel)
		go f.Call(in)
	}

	log.Info("Started everything")
	<-osSignalChannel
	close(osSignalChannel)
	close(messageQueue)
	log.Info("Ending everything ...")
	time.Sleep(1 * time.Second) // TODO move to config
	log.Info("Sleep complete. Goodbye!")
}
