package main

import (
	"gohausen/dispatcher"
	"gohausen/kafka"
	"gohausen/mqtt"
)

const queueChannelSize = 100

func main() {
	var c = make(chan string, queueChannelSize)
	go kafka.Producer(c)
	go dispatcher.Main(c)
	mqtt.Consumer(c)
}
