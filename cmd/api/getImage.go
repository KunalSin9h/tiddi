package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kunalsin9h/tiddi/internal/db"
)

type GetImageRequest struct {
	Uiid string `json:"uiid"`
}

type GetImageResponse struct {
	Title string `json:"title"`
	Image []byte `json:"image"`
}

/*
GetImage is the endpoint to get the image details
Endpoint: Get https://your-domain.com/get-image/{uuid}

@description

	`uiid` -> Unique Image Id

@author Kunal Singh <kunalsin9h@gmail.com>
@type Api Handler Function
*/
func GetImage(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	if r.Method != "GET" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	uiid, err := getUiid(r.URL.Path)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	if uiid == "" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	title, image, err := db.GetImageDetails(uiid)

	if err != nil {
		log.Printf("[GET-IMAGE] Unable to get image details: %v\n", err)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	response := GetImageResponse{
		Title: title,
		Image: image,
	}

	resBytes, err := json.Marshal(response)

	if err != nil {
		log.Printf("[GET-IMAGE] Unable to marshal response: %v\n", err)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resBytes)
}
