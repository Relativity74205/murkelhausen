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
