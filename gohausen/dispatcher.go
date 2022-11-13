package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

func dispatcher(queueChannel chan ChannelPayload) {
	router := gin.New()

	router.Use(Logger())

	addShellyHT(router, queueChannel)

	addr := fmt.Sprintf(":%d", Conf.dispatcher.port)
	err := router.Run(addr)
	if err != nil {
		log.WithField("error", err).Fatal("Could not start gin server!")
	}
}

func addShellyHT(router *gin.Engine, queueChannel chan ChannelPayload) {
	router.GET("/shelly/:sensor", func(c *gin.Context) {
		humidity, err := strconv.ParseFloat(c.Query("hum"), 32)
		if err != nil {
			log.WithField("error", err).Error("No hum query parameter in url query string!")
		}
		temperature, _ := strconv.ParseFloat(c.Query("temp"), 32)
		if err != nil {
			log.WithField("error", err).Error("No temp query parameter in url query string!")
		}
		shellyHTData := ShellyHTData{
			Humidity:    humidity,
			Temperature: temperature,
			Id:          c.Query("id"),
		}
		channelPayload := ChannelPayload{
			Topic: Conf.dispatcher.shellyHTKafkaTopic,
			Key:   c.Param("sensor"),
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

func Logger(notLogged ...string) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	var skip map[string]struct{}

	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, p := range notLogged {
			skip[p] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		if _, ok := skip[path]; ok {
			return
		}

		entry := log.WithFields(log.Fields{
			"hostname":   hostname,
			"statusCode": statusCode,
			"latency":    latency, // time to process
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("\"%s %s\" %d (%dms)", c.Request.Method, path, statusCode, latency)
			if statusCode >= http.StatusInternalServerError {
				entry.Error(msg)
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
