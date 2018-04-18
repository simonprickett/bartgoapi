package main

import (
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
		c.String(200, "status")
	})

	// Elevator status
	r.GET("/elevatorstatus", func(c *gin.Context) {
		c.String(200, "elevatorstatus")
	})

	// Service announcements
	r.GET("/serviceannouncements", func(c *gin.Context) {
		c.String(200, "serviceannouncements")
	})

	// Stations
	r.GET("/stations", func(c *gin.Context) {
		c.String(200, "stations")
	})

	// One specific station
	r.GET("/stations/:stationAbbr", func(c *gin.Context) {
		stationAbbr := c.Params.ByName("stationAbbr")
		c.String(200, "stations "+stationAbbr)
	})

	// All departures from all stations
	r.GET("/departures", func(c *gin.Context) {
		c.String(200, "stations")
	})

	// Departures from one specific station
	r.GET("/departures/:stationAbbr", func(c *gin.Context) {
		stationAbbr := c.Params.ByName("stationAbbr")
		c.String(200, "departures "+stationAbbr)
	})

	// Access details for all stations
	r.GET("/stationaccess", func(c *gin.Context) {
		c.String(200, "access")
	})

	// Access details for one specific station
	r.GET("/stationaccess/:stationAbbr", func(c *gin.Context) {
		stationAbbr := c.Params.ByName("stationAbbr")
		c.String(200, "access "+stationAbbr)
	})

	// Station information for all stations
	r.GET("/stationinfo", func(c *gin.Context) {
		c.String(200, "info")
	})

	// Station information for one specific station
	r.GET("/stationinfo/:stationAbbr", func(c *gin.Context) {
		stationAbbr := c.Params.ByName("stationAbbr")
		c.String(200, "info "+stationAbbr)
	})

	// Ticket cost and route between two stations
	r.GET("/tickets/:fromStation/:toStation", func(c *gin.Context) {
		fromStation := c.Params.ByName("fromStation")
		toStation := c.Params.ByName("toStation")
		c.String(200, "tickets "+fromStation+", "+toStation)
	})

	// Station closest to a given latitude/longitude
	r.GET("/station/:latitude/:longitude", func(c *gin.Context) {
		latitude := c.Params.ByName("latitude")
		longitude := c.Params.ByName("longitude")
		c.String(200, "station "+latitude+", "+longitude)
	})

	// Resolve issue https://github.com/gin-gonic/gin/issues/205
	// All stations ordered by distance from a given latitude/longitude
	r.GET("/stations/:latitude/:longitude", func(c *gin.Context) {
		latitude := c.Params.ByName("latitude")
		longitude := c.Params.ByName("longitude")
		c.String(200, "stations "+latitude+", "+longitude)
	})

	r.Run(":9999")
}
