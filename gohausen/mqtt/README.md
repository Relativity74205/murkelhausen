# test messages

```bash
mosquitto_pub -h localhost -p 1883 -t test_topic -m "{\"val1\": 44}"

mosquitto_pub -h localhost -p 1883 -t "zigbee2mqtt/XaomiTempCellarHobby" -m "{\"battery\":50,\"humidity\":53.9,\"linkquality\":61,\"power_outage_count\":17,\"temperature\":20.83,\"voltage\":3025}"
```

