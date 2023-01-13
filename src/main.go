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

var HOST = os.Getenv("HOST")
var PORT = os.Getenv("PORT")

func init() {

	if PORT == "" {
		log.Println("[MAIN] Using default port (5656)")
		PORT = "5656"
		os.Setenv("PORT", PORT)
	}

	if HOST == "" {
		log.Printf("[MAIN] Using default host (http://localhost:%s)", PORT)
		HOST = "http://localhost:" + PORT
		os.Setenv("HOST", HOST)
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
		Endpoint: POST https://your-domain.com/get-image/
		Get the full image, title and other details .
	*/
	http.HandleFunc("/get-image/", api.GetImage)

	/*
		Endpoint: POST https://your-domain.com/upload-image/
		Create a non blocking channel to handle the request

	*/
	http.HandleFunc("/upload-image/", api.UploadImage)

	/*
		Start The Server
	*/
	log.Printf("Started server at port %s", PORT)
	SERVER.ListenAndServe()
}
