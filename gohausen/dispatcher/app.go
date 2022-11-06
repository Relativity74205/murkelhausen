package dispatcher

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ShellyHTData struct {
	SensorName  string
	Humidity    float64
	Temperature float64
	Id          string
	QueryPath   string
	FullPath    string
}

func Main(queueChannel chan string) {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//shelly1?hum=65&temp=21.50&id=shellyht-C8E78E"
	router.GET("/shelly/:sensor", func(c *gin.Context) {
		queryPath := ""
		for k, v := range c.Request.URL.Query() {
			queryPath += fmt.Sprintf("%s=%s", k, v)
		}

		humidity, _ := strconv.ParseFloat(c.Query("hum"), 32)
		temperature, _ := strconv.ParseFloat(c.Query("temp"), 32)
		shellyHTData := ShellyHTData{
			c.Param("sensor"),
			humidity,
			temperature,
			c.Query("id"),
			queryPath,
			c.FullPath(),
		}
		jsonData, _ := json.Marshal(shellyHTData)
		fmt.Printf("Sending shelly message (%s) to Kafka.\n", string(jsonData))
		queueChannel <- string(jsonData)

		c.String(http.StatusOK, "OK")
	})

	err := router.Run(":8123") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(err)
	}
}
