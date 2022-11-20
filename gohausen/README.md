#

## Build & Deploy

```bash
go build . && scp gohausen beowulf:~ 
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

## MQTT consumer

### test messages

```bash
mosquitto_pub -h localhost -p 1883 -t test_topic -m "{\"val1\": 44}"

mosquitto_pub -h localhost -p 1883 -t "zigbee2mqtt/XaomiTempCellarHobby" -m "{\"battery\":50,\"humidity\":53.9,\"linkquality\":61,\"power_outage_count\":17,\"temperature\":20.83,\"voltage\":3025}"
```


## config

```yaml
app:
  queueChannelSize: 100
mqtt:
  broker: "tcp://192.168.1.69:1883"
  clientId: gohausen
  cleanSession: false
  qos: 0
kafka:
  broker: 192.168.1.69:19092
  schemaRegistryUrl: http://192.168.1.69:8081
dispatcher:
  port: 8123
  shellyHTKafkaTopic: shelly_ht_sensor
mappingMqttKafka:
  test:
    qos: 0
    mqttTopics:
      - test_topic
    kafkaTopic: test_topic
    payloadType: MQTTTestData
  xiaomi:
    qos: 0
    mqttTopics:
      - zigbee2mqtt/XiaomiTempCellarHobby
      - zigbee2mqtt/XiaomiTempCellarVersorgung
    kafkaTopic: xiaomi_mi_sensor
    payloadType: XiaomiMiSensorData
  aqara:
    qos: 0
      - zigbee2mqtt/AqaraTempCellar
    kafkaTopic: aqara_sensor
    payloadType: AqaraSensorData
```