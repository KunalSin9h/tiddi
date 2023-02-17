package main

import (
	"encoding/json"
	"net/http"
)

func (app *Config) enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST")
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	URL     string `json:"url,omitempty"`
}

func (app *Config) errorJson(w *http.ResponseWriter, message string, statusCode int, contentType string) {
	response := jsonResponse{
		Error:   true,
		Message: message,
	}

	jsonEncoded, err := json.Marshal(response)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		return
	}

	(*w).Header().Set("Content-Type", contentType)
	(*w).WriteHeader(statusCode)
	(*w).Write(jsonEncoded)
}
