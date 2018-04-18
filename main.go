package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO load data and set up timer task
	// TODO static info page on /

	r := gin.Default()

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
