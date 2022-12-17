package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

func addShellyHT(router *gin.Engine, queueChannel chan ChannelPayload) {
	router.GET("/shelly/:sensor", func(c *gin.Context) {
		humidity, err := strconv.ParseFloat(c.Query("hum"), 32)
		if err != nil {
			log.WithField("error", err).Error("No hum query parameter in url query string!")
		}
		temperature, err := strconv.ParseFloat(c.Query("temp"), 32)
		if err != nil {
			log.WithField("error", err).Error("No temp query parameter in url query string!")
		}
		sensorName := c.Param("sensor")
		shellyHTData := ShellyHTData{
			SensorName:  sensorName,
			Tstamp:      time.Now().Local(),
			Humidity:    humidity,
			Temperature: temperature,
			Id:          c.Query("id"),
		}
		channelPayload := ChannelPayload{
			Topic: Conf.Dispatcher.ShellyHTKafkaTopic,
			Key:   sensorName,
			Value: shellyHTData,
		}

		queueChannel <- channelPayload
		log.WithFields(log.Fields{
			"topic": channelPayload.Topic,
			"key":   channelPayload.Key,
			"value": fmt.Sprintf("%+v", channelPayload.Value),
		}).Info("Received dispatcher message and send data to queueChannel.")

		c.String(http.StatusOK, "OK")
	})
}

func addShellyFlood(router *gin.Engine, queueChannel chan ChannelPayload) {
	router.GET("/shelly_flood/:sensor", func(c *gin.Context) {

		temperature, err := strconv.ParseFloat(c.Query("temp"), 32)
		if err != nil {
			log.WithField("error", err).Error("No temp query parameter in url query string!")
		}
		flood, err := strconv.Atoi(c.Query("flood"))
		if err != nil {
			log.WithField("error", err).Error("No temp query parameter in url query string!")
		}
		sensorName := c.Param("sensor")
		shellyFloodData := ShellyFloodData{
			SensorName:  sensorName,
			Tstamp:      time.Now().Local(),
			Temperature: temperature,
			Flood:       flood,
			Id:          c.Query("id"),
		}
		channelPayload := ChannelPayload{
			Topic: Conf.Dispatcher.ShellyFloodTopic,
			Key:   sensorName,
			Value: shellyFloodData,
		}

		queueChannel <- channelPayload
		log.WithFields(log.Fields{
			"topic": channelPayload.Topic,
			"key":   channelPayload.Key,
			"value": fmt.Sprintf("%+v", channelPayload.Value),
		}).Info("Received dispatcher message and send data to queueChannel.")

		c.String(http.StatusOK, "OK")
	})
}
