package common

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

type DispatcherTestData struct {
	Value  int       `json:"value"`
	Tstamp time.Time `json:"tstamp"`
}

func (data DispatcherTestData) data() KafkaValue { return data }

type SystemState struct {
	Tstamp              time.Time `json:"tstamp"`
	Hostname            string    `json:"hostname"`
	Uptime              int64     `json:"uptime"`
	MemoryTotal         int64     `json:"memory_total"`
	MemoryAvailable     int64     `json:"memory_available"`
	MemoryUsed          int64     `json:"memory_used"`
	MemoryUsedPercent   float64   `json:"memory_used_percent"`
	MemoryFree          int64     `json:"memory_free"`
	CpuCores            int       `json:"cpu_cores"`
	CpuLogical          int       `json:"cpu_logical"`
	CpuUsageAvg         float64   `json:"cpu_usage_avg"`
	RootDiskTotal       int64     `json:"root_disk_total"`
	RootDiskFree        int64     `json:"root_disk_free"`
	RootDiskUsed        int64     `json:"root_disk_used"`
	RootDiskUsedPercent float64   `json:"root_disk_used_percent"`
	Load01              float64   `json:"load01"`
	Load05              float64   `json:"load05"`
	Load15              float64   `json:"load15"`
	NetworkBytesSent    int64     `json:"network_bytes_sent"`
	NetworkBytesRecv    int64     `json:"network_bytes_recv"`
	ProcessCount        int       `json:"process_count"`
}

func (data SystemState) data() KafkaValue { return data }

type Usage struct {
	Total   float32
	Current float32
}

type PowerDataRaw struct {
	Time  string
	Usage Usage
}

type PowerData struct {
	SensorName   string    `json:"sensorname"`
	Tstamp       time.Time `json:"tstamp"`
	PowerTotal   float64   `json:"power_total"`
	PowerCurrent float64   `json:"power_current"`
}

func (data PowerData) data() KafkaValue { return data }
