package main

import log "github.com/sirupsen/logrus"

// TODO log rotation
func setupLogger() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat:        "2006-01-02T15:04:05.999Z07:00",
		FullTimestamp:          true,
		ForceColors:            true,
		DisableLevelTruncation: false,
		PadLevelText:           true,
	})
	log.SetLevel(log.DebugLevel) // TODO config
	log.SetReportCaller(false)   // otherwise it is very verbose...
}
