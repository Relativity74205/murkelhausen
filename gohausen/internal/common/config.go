package common

var Conf gohausenConfig

type appConfig struct {
	Modules          []string
	DebugMode        bool
	QueueChannelSize int
}

type mqttConfig struct {
	Broker       string
	ClientId     string
	CleanSession bool
	Qos          int
}

type kafkaConfig struct {
	Broker            string
	SchemaRegistryUrl string
}

type dispatcherConfig struct {
	Port               int
	ShellyHTKafkaTopic string
	ShellyFloodTopic   string
	TestTopic          string
}

type mqttKafkaMapping struct {
	Qos         int
	MqttTopics  []string
	KafkaTopic  string
	PayloadType string
	DebugMode   bool
}

type gohausenConfig struct {
	App               appConfig
	Mqtt              mqttConfig
	Kafka             kafkaConfig
	Dispatcher        dispatcherConfig
	MqttKafkaMappings []mqttKafkaMapping
}
