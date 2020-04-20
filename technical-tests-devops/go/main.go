package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"techtestgo/datastore"

	"os"
)

var (
	redisHost string
	redisPort string
	applicationHost string
	applicationPort string
	ds *datastore.Datastore
)

const (
	defaultRedisHost = "localhost"
	defaultRedisPort = "6379"
	defaultApplicationHost = "0.0.0.0"
	defaultApplicationPort = "8080"
)

func init() {
	redisHost = os.Getenv("REDIS_HOST")
	redisPort = os.Getenv("REDIS_PORT")
	
	if redisHost == "" {
		redisHost = defaultRedisHost
	}
	if redisPort == "" {
		redisPort = defaultRedisPort
	}

	applicationHost = os.Getenv("APPLICATION_HOST")
	applicationPort = os.Getenv("APPLICATION_PORT")

	if applicationHost == "" {
		applicationHost = defaultApplicationHost
	}
	if applicationPort == "" {
		applicationPort = defaultApplicationPort
	}
}

func main() {
	// Open connectoin to datastore
	var err error
	ds, err = datastore.New(redisHost, redisPort)
	if err != nil {
		log.Fatalf("Unable to connect to redis: %v", err)
	}

	e := echo.New()
	e.File("/", "public/index.html")
	e.GET("/api/clicks", getCountHandler)
	e.POST("/api/clicks", clickHandler)
	e.GET("/api/healthcheck", healthCheckHandler)

	connectionString := fmt.Sprintf("%s:%s", applicationHost, applicationPort)
	e.Logger.Fatal(e.Start(connectionString))
}

func getCountHandler(c echo.Context) error {
	currentCount, err := ds.GetCount()
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]int{"count": currentCount})
}

func clickHandler(c echo.Context) error {
	// Open connection to redis
	currentCount, err := ds.GetCount()
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Increment count
	newCount := currentCount + 1
	if err := ds.SetCount(newCount); err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]int{"count": newCount})
}

func healthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"ping": "pong"})
}
