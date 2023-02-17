package main

import (
	"net/http"
)

func (app *Config) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.handle)
	return mux
}
