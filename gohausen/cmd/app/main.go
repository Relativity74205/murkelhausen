package main

import (
	"github.com/Relativity74205/murkelhausen/gohausen/internal/common"
	"github.com/Relativity74205/murkelhausen/gohausen/internal/dispatcher"
	"github.com/Relativity74205/murkelhausen/gohausen/internal/kafka"
	"github.com/Relativity74205/murkelhausen/gohausen/internal/mqtt"
	"github.com/Relativity74205/murkelhausen/gohausen/internal/tasks"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

func main() {
	common.SetupConfig()
	setupLogger()
	startModules()
}

var modulesMapping = map[string]interface{}{
	"scheduler":     tasks.Start,
	"kafkaProducer": kafka.Start,
	"dispatcher":    dispatcher.Start,
	"mqttConsumer":  mqtt.Start,
}

func startModules() {
	log.Info("Creating channels...")

	var messageQueue = make(chan common.ChannelPayload, common.Conf.App.QueueChannelSize)
	osSignalChannel := make(chan os.Signal, 1)
	signal.Notify(osSignalChannel, os.Interrupt, syscall.SIGTERM)

	log.WithField("moduleList", common.Conf.App.Modules).Info("Starting modulesMapping...")
	// TODO start also kafkaProducer as go routine and end main function when all go routines close.
	// TODO go routines shall close on system call
	//go dispatcher(messageQueue)
	//go mqttConsumer(messageQueue)
	for _, moduleName := range common.Conf.App.Modules {
		moduleCallable, ok := modulesMapping[moduleName]
		if !ok {
			log.WithField("module", moduleName).Error("Module not found.")
		}

		log.WithField("module", moduleName).Info("Starting module.")
		callable := reflect.ValueOf(moduleCallable)
		args := []reflect.Value{
			reflect.ValueOf(messageQueue), reflect.ValueOf(osSignalChannel),
		}
		go callable.Call(args)
	}

	log.Info("Started everything")
	<-osSignalChannel
	close(osSignalChannel)
	close(messageQueue)
	log.Info("Ending everything ...")
	time.Sleep(1 * time.Second) // TODO move to config
	log.Info("Sleep complete. Goodbye!")
}
