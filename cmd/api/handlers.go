package main

import (
	"errors"
	"net/http"
	"strings"
)

func (app *Config) handle(w http.ResponseWriter, r *http.Request) {
	app.enableCors(&w)

	if r.Method == "POST" {

	} else if r.Method == "GET" {
		app.getImage(w, r)
	} else {
		app.errorJson(&w, "Method Not Allowed", http.StatusMethodNotAllowed, "application/json")
	}
}

func (app *Config) getImage(w http.ResponseWriter, r *http.Request) {
	uuid, err := extractUuid(r.URL.Path)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(uuid))

	// TODO: get the images data from redis
	// and write
}

func extractUuid(url string) (string, error) {
	uuid := strings.Split(url, "/")

	numberOfPotentialUuid := 0
	for _, val := range uuid {
		if len(val) > 0 {
			numberOfPotentialUuid++
		}
	}

	if numberOfPotentialUuid != 1 {
		return "", errors.New("bad request")
	}

	return uuid[1], nil
}
