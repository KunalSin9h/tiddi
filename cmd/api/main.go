package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	data "github.com/kunalsin9h/tiddi/cmd/data"
	redis "github.com/redis/go-redis/v9"
)

var client *redis.Client

type Config struct {
	PORT   string
	Models data.Models
}

func main() {

	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:",
		Password: "tiddi",
		DB:       0,
	})

	app := Config{
		PORT:   "5000",
		Models: data.New(client),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", app.PORT),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      app.routes(),
	}

	log.Printf("Started server on port: %s", app.PORT)
	log.Fatal(server.ListenAndServe())
}
