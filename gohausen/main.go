package main

const queueChannelSize = 100

func main() {
	var c = make(chan ChannelPayload, queueChannelSize)
	go kafkaProducer(c)
	go dispatcher(c)
	mqttConsumer(c)
}
