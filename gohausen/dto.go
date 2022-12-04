package main

import (
	"time"
)

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

// AqaraSensorData {"battery":93,"humidity":51.94,"linkquality":109,"power_outage_count":6,"pressure":998.5,"temperature":21.01,"voltage":3175}
type AqaraSensorData struct {
	SensorName       string    `json:"sensorname"`
	Tstamp           time.Time `json:"tstamp"`
	Battery          int       `json:"battery"`
	Humidity         float64   `json:"humidity"`
	LinkQuality      int       `json:"linkquality"`
	PowerOutageCount int       `json:"power_outage_count"`
	Pressure         float64   `json:"pressure"`
	Temperature      float64   `json:"temperature"`
	Voltage          int       `json:"voltage"`
}

func (data AqaraSensorData) data() KafkaValue { return data }

type MQTTTestData struct {
	Val1 int `json:"val1"`
}

func (data MQTTTestData) data() KafkaValue { return data }

type ShellyHTData struct {
	SensorName  string    `json:"sensorname"`
	Tstamp      time.Time `json:"tstamp"`
	Humidity    float64   `json:"humidity"`
	Temperature float64   `json:"temperature"`
	Id          string    `json:"id"`
}

func (data ShellyHTData) data() KafkaValue { return data }

type ShellyFloodData struct {
	SensorName  string    `json:"sensorname"`
	Tstamp      time.Time `json:"tstamp"`
	Temperature float64   `json:"temperature"`
	Flood       int       `json:"flood"`
	Id          string    `json:"id"`
}

func (data ShellyFloodData) data() KafkaValue { return data }

// {"Time":"2022-11-29T22:02:05","Usage":{"Total":11381.07,"Current":2655.5}}
// {"Time":"2022-12-03T21:42:39","Usage":{"Total":11836.01,"Current":2943.0,"Current_p1":-2.0,"Current_p2":2.0,"Current_p3":2942.0}}

type Usage struct {
	Total   float32
	Current float32
}

type PowerDataRaw struct {
	Time  string
	Usage Usage
}

type PowerData struct {
	SensorName   string  `json:"sensorname"`
	Tstamp       string  `json:"tstamp"`
	PowerTotal   float32 `json:"powertotal"`
	PowerCurrent float32 `json:"powercurrent"`
}

func (data PowerData) data() KafkaValue { return data }
