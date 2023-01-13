package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/znip-in/tiddi/src/api"
	_ "github.com/znip-in/tiddi/src/db"
)

var DOMAIN = os.Getenv("DOMAIN")
var PORT = os.Getenv("PORT")

func init() {

	if DOMAIN == "" {
		log.Println("Using default domain http://localhost")
		DOMAIN = "http://localhost"
		os.Setenv("DOMAIN", DOMAIN)
	}

	if PORT == "" {
		log.Println("Using default port :5656")
		PORT = "5656"
		os.Setenv("PORT", PORT)
	}

}

func main() {

	/*
		Http Server With Timeout
	*/
	SERVER := http.Server{
		Addr:              fmt.Sprintf(":%s", PORT),
		ReadHeaderTimeout: 3 * time.Second,
	}

	/*
		Endpoint: GET https://your-domain.com/
	*/
	http.HandleFunc("/", api.Home)

	/*
		Endpoint: POST https://your-domain.com/upload-image
		Create a non blocking channel to handle the request

	*/
	http.HandleFunc("/upload-image/", api.UploadImage)

	/*
		Start The Server
	*/
	log.Printf("Started server at %s:%s", DOMAIN, PORT)
	SERVER.ListenAndServe()
}
