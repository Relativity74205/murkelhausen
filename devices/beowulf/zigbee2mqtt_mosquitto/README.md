# Zigbee2MQTT and Mosquitto

## Links

- Guide: <https://www.libe.net/zigbee2mqtt>
- MQTT Broker: <https://mosquitto.org/>
- MQTT Python Client: <https://github.com/eclipse/paho.mqtt.python>

Additional:

MQTT topics:

- <https://www.hivemq.com/blog/mqtt-essentials-part-5-mqtt-topics-best-practices/>

Quality of service:

- <https://www.hivemq.com/blog/mqtt-essentials-part-6-mqtt-quality-of-service-levels/>

## Zigbee

- Zigbee2MQTT Frontend on <http://192.168.1.69:8080/>

## Mosquitto

### Install local mosquitto clients

<http://www.steves-internet-guide.com/install-mosquitto-linux/>

```bash
sudo apt-add-repository ppa:mosquitto-dev/mosquitto-ppa
sudo apt-get update
sudo apt-get install mosquitto-clients
```

### Commands

Subscribe to specific topic

```bash
mosquitto_sub -h localhost -p 1883 -t postgres_test
```

Subscribe to all topics:

```bash
mosquitto_sub -v -h localhost -p 1883 -t '#'
```

Publish to topic:

```bash
mosquitto_pub -h localhost -p 1883 -t postgres_test -m "{\"val1\": 1}"
```

### Password

To change password, exec into mosquitto container and run in path `mosquitto/config/` the following command:

```bash
mosquitto_passwd -c passwd USER
```

TODO: How to add?
