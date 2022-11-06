package mqtt

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var c chan string

const topic = "test_topic"
const broker = "tcp://192.168.1.69:1883"
const id = "gohausen_client"
const cleanSession = false

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	c <- string(message.Payload())
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}

func Consumer(queueChannel chan string) {
	c = queueChannel
	osSignalChannel := make(chan os.Signal, 1)
	signal.Notify(osSignalChannel, os.Interrupt, syscall.SIGTERM)

	qos := 0

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(id)
	opts.SetCleanSession(cleanSession)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe(topic, byte(qos), onMessageReceived); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	<-osSignalChannel
	fmt.Println("Disconnecting mqtt client and closing queueChannel ...")
	client.Disconnect(250)
	close(c)
	fmt.Println("Mqtt client disconnected and queueChannel closed. Sleep for 1 second to give other routines chance to stop.")
	time.Sleep(1 * time.Second)
	fmt.Println("Sleep complete. Goodbye!")
}
