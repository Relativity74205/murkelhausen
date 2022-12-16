#

## Build & Deploy

```bash
go build . && scp gohausen beowulf:~ 
```

Build for raspberry
```bash
env GOOS=linux GOARCH=arm GOARM=7 go build ./cmd/app
```

## Start as background

```bash
nohup ./gohausen &>gohausen.logs &
```

## Links

[Go frameworks](https://github.com/mingrammer/go-web-framework-stars)


## Dispatcher microservice

### Shelly sensor

http://192.168.1.75/

### Shelly flood

192.168.1.79

nach

http://192.168.1.69:8123/shelly_flood/flood1

## MQTT consumer

### test messages

```bash
mosquitto_pub -h localhost -p 1883 -t test_topic -m "{\"val1\": 44}"

mosquitto_pub -h localhost -p 1883 -t "zigbee2mqtt/XiaomiTempCellarHobby" -m "{\"battery\":50,\"humidity\":53.9,\"linkquality\":61,\"power_outage_count\":17,\"temperature\":20.83,\"voltage\":3025}"
```

```bash
mosquitto_sub -v -h localhost -p 1883 -t '#'
```


## config

```yaml
app:
  queueChannelSize: 100
mqtt:
  broker: "tcp://192.168.1.69:1883"
  clientId: gohausen
  cleanSession: false
  qos: 1
kafka:
  broker: 192.168.1.69:19092
  schemaRegistryUrl: http://192.168.1.69:8081
dispatcher:
  port: 8123
  shellyHTKafkaTopic: shelly_ht_sensor
  shellyFloodTopic: shelly_flood
mappingMqttKafka:
  test:
    qos: 1
    mqttTopics:
      - test_topic
    kafkaTopic: test_topic
    payloadType: MQTTTestData
  xiaomi:
    qos: 1
    mqttTopics:
      - zigbee2mqtt/XiaomiTempCellarHobby
      - zigbee2mqtt/XiaomiTempCellarVersorgung
    kafkaTopic: xiaomi_mi_sensor
    payloadType: XiaomiMiSensorData
  aqara:
    qos: 1
    mqttTopics:
      - zigbee2mqtt/AqaraTempCellar
    kafkaTopic: aqara_sensor
    payloadType: AqaraSensorData

```