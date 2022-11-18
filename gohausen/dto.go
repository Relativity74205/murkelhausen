package main

import "time"

type ChannelPayload struct {
	Topic string
	Key   string
	Value KafkaValue
}

type KafkaValue interface {
	data() KafkaValue
}

// XiaomiMiSensorData {"battery":50,"humidity":53.9,"linkquality":61,"power_outage_count":17,"temperature":20.83,"voltage":3025}
type XiaomiMiSensorData struct {
	SensorName       string    `json:"sensorname"`
	Tstamp           time.Time `json:"tstamp"`
	Battery          int       `json:"battery"`
	Humidity         float64   `json:"humidity"`
	LinkQuality      int       `json:"linkquality"`
	PowerOutageCount int       `json:"power_outage_count"`
	Temperature      float64   `json:"temperature"`
	Voltage          int       `json:"voltage"`
}

func (data XiaomiMiSensorData) data() KafkaValue { return data }

type MQTTTestData struct {
	Val1 int `json:"val1"`
}

func (data MQTTTestData) data() KafkaValue { return data }

type ShellyHTData struct {
	Humidity    float64
	Temperature float64
	Id          string
}

func (data ShellyHTData) data() KafkaValue { return data }
