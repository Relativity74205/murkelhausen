package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const shellyHTKafkaTopic = "shelly_ht_sensor"

func dispatcher(queueChannel chan ChannelPayload) {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/shelly/:sensor", func(c *gin.Context) {
		humidity, err := strconv.ParseFloat(c.Query("hum"), 32)
		if err != nil {
			fmt.Println(err)
		}
		temperature, _ := strconv.ParseFloat(c.Query("temp"), 32)
		if err != nil {
			fmt.Println(err)
		}
		shellyHTData := ShellyHTData{
			Humidity:    humidity,
			Temperature: temperature,
			Id:          c.Query("id"),
		}
		channelPayload := ChannelPayload{
			Topic: shellyHTKafkaTopic,
			Key:   c.Param("sensor"),
			Value: shellyHTData,
		}

		fmt.Printf("Prepared data for sending to queueChannel: %v \n", channelPayload)
		queueChannel <- channelPayload
		fmt.Println("Send data to queueChannel.")

		c.String(http.StatusOK, "OK")
	})

	err := router.Run(":8123") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(err)
	}
}
