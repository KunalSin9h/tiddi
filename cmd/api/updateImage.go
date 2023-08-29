package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/kunalsin9h/tiddi/internal/db"
)

type UpdateImageRequest struct {
	Uiid  string `json:"uiid"`
	Title string `json:"title"`
	Image []byte `json:"image"`
}

/*
UpdateImage is the endpoint to update image details the image from db
Endpoint: PUT https://your-domain.com/update-image/

@description

	Request.Body = {
		"uiid": "uiid of the image"
		"title": "title for the image"
		"data": "[]byte of the image"
	}

	`uiid` -> Unique Image Id

@author Kunal Singh <kunalsin9h@gmail.com>
@type Api Handler Function
*/

func UpdateImage(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	if r.Method != "PUT" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	data, err := io.ReadAll(r.Body)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	var req UpdateImageRequest

	err = json.Unmarshal(data, &req)

	if err != nil {
		log.Printf("[UPDATE IMAGE] Unable to unmarshal data from request: %v", err)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	if req.Title != "" {
		err = db.UpdateTitle(req.Uiid, req.Title)
		if err != nil {
			log.Printf("[UPDATE TITLE] Error while updating title: %v", err)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
	}

	if req.Image != nil {
		err = db.UpdateImageData(req.Uiid, req.Image)
		if err != nil {
			log.Printf("[UPDATE TITLE] Error while updating image data: %v", err)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
