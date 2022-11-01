## Install local mosquitto clients

http://www.steves-internet-guide.com/install-mosquitto-linux/

```bash
sudo apt-add-repository ppa:mosquitto-dev/mosquitto-ppa
sudo apt-get update
sudo apt-get install mosquitto-clients
```


commands:

```bash
mosquitto_sub -h localhost -p 1883 -t postgres_test
```

```bash
mosquitto_pub -h localhost -p 1883 -t postgres_test -m "{\"val1\": 1}"
```