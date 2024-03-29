package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/kunalsin9h/tiddi/internal/db"
)

type DeleteImageRequest struct {
	Uiid string `json:"uiid"`
}

/*
DeleteImage is the endpoint to delete the image from db
Endpoint: DELETE https://your-domain.com/delete-image/

@description

	Request.Body = {
		"uiid": "uiid of the image"
	}

	`uiid` -> Unique Image Id

@author Kunal Singh <kunalsin9h@gmail.com>
@type Api Handler Function
*/

func DeleteImage(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)

	if r.Method != "DELETE" {
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

	var res DeleteImageRequest

	err = json.Unmarshal(data, &res)

	if err != nil {
		log.Printf("[DELETE IMAGE] Unable to unmarshal data from request: %v", err)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	err = db.DeleteImageFormDB(res.Uiid)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
