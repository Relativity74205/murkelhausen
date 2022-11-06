
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

- https://docs.confluent.io/kafka-connectors/mqtt/current/mqtt-source-connector/index.html
- https://medium.com/python-point/mqtt-and-kafka-8e470eff606b
- https://www.kai-waehner.de/blog/2021/03/15/apache-kafka-mqtt-sparkplug-iot-blog-series-part-1-of-5-overview-comparison/
- https://github.com/kaiwaehner/kafka-connect-iot-mqtt-connector-example/blob/master/live-demo-kafka-connect-iot-mqtt-connector.adoc

## connect API

```bash
curl -s -X POST -H 'Content-Type: application/json' http://localhost:8083/connectors -d '{
    "name" : "mqtt-kafka-postgres-test",
"config" : {
    "connector.class" : "io.confluent.connect.mqtt.MqttSourceConnector",
    "tasks.max" : "1",
    "mqtt.server.uri" : "tcp://192.168.1.69:1883",
    "mqtt.topics" : "postgres_test",
    "kafka.topic" : "postgres_test"
    }
}'

```


```bash
curl -s -X POST -H 'Content-Type: application/json' http://localhost:8083/connectors -d '{
    "name" : "mqtt-kafka-postgres-test_jdbc",
"config" : {
    "connector.class" : "io.confluent.connect.jdbc.JdbcSinkConnector",
    "topics": "postgres_test_2",
    "connection.url": "jdbc:postgresql://192.168.1.69:5432/bar",
    "key.converter": "org.apache.kafka.connect.storage.StringConverter",
    "value.converter": "org.apache.kafka.connect.json.JsonConverter",
    "value.converter.schemas.enable": "false",
    "connection.user": "postgres",
    "connection.password": "foo",
    "dialect.name": "PostgreSqlDatabaseDialect"
    }
}'
```


```bash
curl -s -X DELETE http://localhost:8083/connectors/mqtt-kafka-postgres-test_jdbc
```

## connect config for mqtt

```json
{
  "name": "MqttSourceConnectorConnector_0",
  "config": {
    "name": "MqttSourceConnectorConnector_0",
    "connector.class": "io.confluent.connect.mqtt.MqttSourceConnector",
    "mqtt.server.uri": "tcp://192.168.1.69:1883",
    "mqtt.clean.session.enabled": "false",
    "kafka.topic": "mqtt",
    "mqtt.topics": "zigbee2mqtt/XaomiTempCellarHobby, zigbee2mqtt/XaomiTempCellarVersorgung"
  }
}
```


## connect config for postgres

# delete.enabled=false' and 'pk.mode=none

```json
{
  "name": "mqtt-kafka-postgres-test_jdbc",
  "config": {
    "name": "mqtt-kafka-postgres-test_jdbc",
    "connector.class": "io.confluent.connect.jdbc.JdbcSinkConnector",
    "topics": "postgres_test",
    "connection.url": "jdbc:postgresql://192.168.1.69:5432/bar",
    "connection.user": "postgres",
    "connection.password": "***",
    "dialect.name": "PostgreSqlDatabaseDialect"
  }
}
```