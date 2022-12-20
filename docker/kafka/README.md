<http://192.168.1.69:9021/>

## confluent

- <https://docs.confluent.io/platform/current/platform-quickstart.html#ce-docker-quickstart>

Docker compose:

- <https://github.com/confluentinc/cp-all-in-one/blob/7.2.2-post/cp-all-in-one/docker-compose.yml>

Docker config reference:

- <https://docs.confluent.io/platform/current/installation/docker/config-reference.html>

## kafka connect connectors

```bash
confluent-hub install confluentinc/kafka-connect-mqtt:1.5.2
confluent-hub install confluentinc/kafka-connect-jdbc:10.6.0
```

For kafka-connect-jdbs see: <https://www.confluent.io/hub/confluentinc/kafka-connect-jdbc>

## mqtt to kafka

- <https://docs.confluent.io/kafka-connectors/mqtt/current/mqtt-source-connector/index.html>
- <https://medium.com/python-point/mqtt-and-kafka-8e470eff606b>
- <https://www.kai-waehner.de/blog/2021/03/15/apache-kafka-mqtt-sparkplug-iot-blog-series-part-1-of-5-overview-comparison/>
- <https://github.com/kaiwaehner/kafka-connect-iot-mqtt-connector-example/blob/master/live-demo-kafka-connect-iot-mqtt-connector.adoc>

## connect API

<https://docs.confluent.io/platform/current/connect/references/restapi.html#tasks>

Change config:

```bash

curl -X PUT -H 'Content-Type: application/json' http://192.168.1.69:8083/connectors/PostgresSinkTest2/config --data-binary "@postgres_sink.json" | jq

```

Restart:

```bash
curl -X POST http://192.168.1.69:8083/connectors/PostgresSinkTest2/restart | jq
```

```bash
curl -X POST -H 'Content-Type: application/json' http://localhost:8083/connectors -d '{
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
curl -X POST -H 'Content-Type: application/json' http://localhost:8083/connectors -d '{
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
curl -X DELETE http://localhost:8083/connectors/mqtt-kafka-postgres-test_jdbc
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
  "name": "PostgresSinkTest2",
  "config": {
    "value.converter.schema.registry.url": "http://schema-registry:8081",
    "name": "PostgresSinkTest2",
    "connector.class": "io.confluent.connect.jdbc.JdbcSinkConnector",
    "key.converter": "org.apache.kafka.connect.storage.StringConverter",
    "value.converter": "io.confluent.connect.avro.AvroConverter",
    "transforms": "insertTS",
    "errors.tolerance": "none",
    "topics": "xiaomi_mi_sensor, shelly_ht_sensor, aqara_sensor",
    "transforms.insertTS.type": "org.apache.kafka.connect.transforms.InsertField$Value",
    "transforms.insertTS.timestamp.field": "messageTS",
    "connection.url": "jdbc:postgresql://192.168.1.69:5432/bar",
    "connection.user": "postgres",
    "connection.password": "***",
    "dialect.name": "PostgreSqlDatabaseDialect",
    "delete.enabled": "false",
    "pk.mode": "kafka",
    "auto.create": "true",
    "auto.evolve": "true"
  }
}
```

### kafka producers/connect sinks change offsets

<https://rmoff.net/2019/10/15/skipping-bad-records-with-the-kafka-connect-jdbc-sink-connector/>

```bash
docker-compose exec broker bash
kafka-consumer-groups --bootstrap-server broker:29092 --list
kafka-consumer-groups --bootstrap-server broker:29092 --describe --group connect-PostgresSinkTest2
kafka-consumer-groups --bootstrap-server broker:29092 --group connect-PostgresSinkTest2 --reset-offsets --topic power_data --to-offset 43220 --execute
```
