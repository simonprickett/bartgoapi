package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	cron "github.com/robfig/cron"
)

func updateStations() {
	// TODO update cached station data.
	fmt.Println("Updating stations cache...")

	httpClient := &http.Client{
		Timeout: time.Second * 4,
	}

	// TODO: Get the key out of here and don't use a demo key...
	response, err := httpClient.Get("http://api.bart.gov/api/stn.aspx?cmd=stns&key=MW9S-E7SL-26DU-VV8V&json=y")

	if err != nil {
		fmt.Printf("Error accessing BART API: %v\n", err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		// Look at this for JSON:
		// https://stackoverflow.com/questions/17156371/how-to-get-json-response-in-golang
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Printf("Error accessing BART API: %v\n", err)
			return
		}

		fmt.Println(string(body))
	} else {
		fmt.Printf("Error accessing BART API: %v\n", response.StatusCode)
		return
	}
}

func main() {
	updateStations()

	cr := cron.New()
	cr.AddFunc("@every 10s", updateStations)
	cr.Start()

	r := gin.Default()

	// TODO static info page on / served from a file
	// https://github.com/gin-gonic/gin/issues/75
	// r.Static("/assets", "./assets")

	// AppEngine health
	r.GET("/_ah/health", func(c *gin.Context) {
		c.String(200, "ok")
	})

	// System status
	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"method": "status"})
	})

	// Elevator status
	r.GET("/elevatorstatus", func(c *gin.Context) {
		c.JSON(200, gin.H{"method": "elevatorstatus"})
	})

	// Service announcements
	r.GET("/serviceannouncements", func(c *gin.Context) {
		c.JSON(200, gin.H{"method": "serviceannouncements"})
	})

	// Stations
	r.GET("/stations", func(c *gin.Context) {
		c.JSON(200, gin.H{"method": "stations"})
	})

	// One specific station
	r.GET("/stations/:stationAbbr", func(c *gin.Context) {
		stationAbbr := strings.ToLower(c.Params.ByName("stationAbbr"))
		c.JSON(200, gin.H{"method": "stations", "abbr": stationAbbr})
	})

	// All departures from all stations
	r.GET("/departures", func(c *gin.Context) {
		c.JSON(200, gin.H{"method": "departures"})
	})

	// Departures from one specific station
	r.GET("/departures/:stationAbbr", func(c *gin.Context) {
		stationAbbr := strings.ToLower(c.Params.ByName("stationAbbr"))
		c.JSON(200, gin.H{"method": "departures", "abbr": stationAbbr})
	})

	// Access details for all stations
	r.GET("/stationaccess", func(c *gin.Context) {
		c.JSON(200, gin.H{"method": "access"})
	})

	// Access details for one specific station
	r.GET("/stationaccess/:stationAbbr", func(c *gin.Context) {
		stationAbbr := strings.ToLower(c.Params.ByName("stationAbbr"))
		c.String(200, "access "+stationAbbr)
	})

	// Station information for all stations
	r.GET("/stationinfo", func(c *gin.Context) {
		c.JSON(200, gin.H{"method": "stationinfo"})
	})

	// Station information for one specific station
	r.GET("/stationinfo/:stationAbbr", func(c *gin.Context) {
		stationAbbr := strings.ToLower(c.Params.ByName("stationAbbr"))
		c.JSON(200, gin.H{"method": "stationinfo", "abbr": stationAbbr})
	})

	// Ticket cost and route between two stations
	r.GET("/tickets/:fromStation/:toStation", func(c *gin.Context) {
		fromStation := strings.ToLower(c.Params.ByName("fromStation"))
		toStation := strings.ToLower(c.Params.ByName("toStation"))
		c.JSON(200, gin.H{"method": "tickets", "fromStation": fromStation, "toStation": toStation})
	})

	// Station closest to a given latitude/longitude
	r.GET("/station/:latitude/:longitude", func(c *gin.Context) {
		// TODO ensure these are both floats...
		latitude := c.Params.ByName("latitude")
		longitude := c.Params.ByName("longitude")
		c.JSON(200, gin.H{"method": "station", "latitude": latitude, "toStation": longitude})
	})

	// Resolve issue https://github.com/gin-gonic/gin/issues/205
	// All stations ordered by distance from a given latitude/longitude
	// r.GET("/stations/:latitude/:longitude", func(c *gin.Context) {
	// 	latitude := c.Params.ByName("latitude")
	// 	longitude := c.Params.ByName("longitude")
	// 	c.String(200, "stations "+latitude+", "+longitude)
	// })

	r.Run(":9999")
}
