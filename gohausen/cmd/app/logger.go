package main

import log "github.com/sirupsen/logrus"

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
