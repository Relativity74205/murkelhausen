import csv
import logging
from pathlib import Path

import paho.mqtt.client
from paho.mqtt.client import Client
import pendulum

log = logging.getLogger(__name__)


FILENAME = "mqtt_data.csv"
p = Path(FILENAME)


def create_data_file():
    if not p.exists():
        with open(p, "a") as f:
            # TODO write get_csv_writer function
            csv_writer = csv.writer(f, delimiter=',')
            csv_writer.writerow(["tstamp", "topic", "message"])


def on_connect(client, userdata, flags, rc):
    log.info(f"Connected to MQTT broker with result code {str(rc)}.")

    # Subscribing in on_connect() means that if we lose the connection and
    # reconnect then subscriptions will be renewed.
    # client.subscribe("$SYS/#")
    client.subscribe(topic="zigbee2mqtt/+", qos=1)
    # TODO find out to which zigbee channels the client subscribed
    log.info(f"Subscribed to zigbee channels...")


# The callback for when a PUBLISH message is received from the server.
def on_message(client, userdata, msg):
    # TODO get timestamp from msg?
    now = pendulum.now("Europe/Berlin")
    msg_payload = msg.payload.decode("utf-8")
    with open(p, "a") as f:
        csv_writer = csv.writer(f, delimiter=',')
        csv_writer.writerow([now, msg.topic, msg_payload])
    log.info(f"new message - {now}: {msg.topic} - {msg_payload}")


def create_client() -> paho.mqtt.client.Client:
    client = Client(client_id="murkelhausen_mqtt_to_kafka", clean_session=False)
    client.on_connect = on_connect
    client.on_message = on_message

    client.connect(host="localhost", port=1883, keepalive=60)

    return client


def main():
    create_data_file()
    client = create_client()
    try:
        client.loop_forever()
    except Exception:
        log.exception("Error during listening to MQTT topic. Aborting")


if __name__ == "__main__":
    main()
