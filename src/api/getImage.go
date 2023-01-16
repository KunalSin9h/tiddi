package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/kunalsin9h/tiddi/src/db"
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
Endpoint: POST https://your-domain.com/get-image/

@description

	Request.Body = {
		"uiid": "uiid of the image"
	}

	`uiid` -> Unique Image Id

@author Kunal Singh <kunalsin9h@gmail.com>
@type Api Handler Function
*/
func GetImage(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	data, err := io.ReadAll(r.Body)

	if err != nil || len(data) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	var request GetImageRequest
	if err := json.Unmarshal(data, &request); err != nil {
		log.Printf("[GET-IMAGE] unable to unmarshal data received: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	if request.Uiid == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	title, image, err := db.GetImageDetails(request.Uiid)

	if err != nil {
		log.Printf("[GET-IMAGE] Unable to get image details: %v\n", err)
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resBytes)
}
