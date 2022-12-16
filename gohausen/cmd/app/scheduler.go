package main

import (
	"github.com/go-co-op/gocron"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func gohausenScheduler(queueChannel chan ChannelPayload, osSignalChannel chan os.Signal) {
	log.Info("Setting up schedules...")
	s := gocron.NewScheduler(time.UTC)
	_, _ = s.Every(15).Seconds().Do(func() { getStats(queueChannel) })

	log.Info("Starting scheduler asynchronous...")
	s.StartAsync()
	log.Info("Scheduler started...")

	<-osSignalChannel
	log.Info("Got message to say goodbye, stopping scheduler...")
	s.Stop()
	log.Info("Scheduler stopped!")
}
