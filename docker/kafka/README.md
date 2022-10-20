
## confluent

- https://docs.confluent.io/platform/current/platform-quickstart.html#ce-docker-quickstart

Docker compose:

- https://github.com/confluentinc/cp-all-in-one/blob/7.2.2-post/cp-all-in-one/docker-compose.yml

Docker config reference:

- https://docs.confluent.io/platform/current/installation/docker/config-reference.html

## kafka connect connectors

```bash
confluent-hub install confluentinc/kafka-connect-mqtt:1.5.2
confluent-hub install confluentinc/kafka-connect-jdbc:10.6.0
```
For kafka-connect-jdbs see: https://www.confluent.io/hub/confluentinc/kafka-connect-jdbc


## mqtt to kafka:

- https://medium.com/python-point/mqtt-and-kafka-8e470eff606b
- https://www.kai-waehner.de/blog/2021/03/15/apache-kafka-mqtt-sparkplug-iot-blog-series-part-1-of-5-overview-comparison/
- 