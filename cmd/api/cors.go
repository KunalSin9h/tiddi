package api

import "net/http"

/*
EnableCors - Enable CORS for the given response writer
*/
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "https://i.imgur.com")
	(*w).Header().Set("Access-Control-Allow-Origin", "https://tiddi.kunalsin9h.com")
}
