package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/znip-in/tiddi/src/db"
	"github.com/znip-in/tiddi/src/uuid"
)

type UploadImageRequest struct {
	Data  []byte `json:"data"`
	Title string `json:"title"`
}

type UploadImageResponse struct {
	URL string `json:"url"`
}

/*
UploadImage endpoint to upload an image with image Data and title
Endpoint: POST https://your-domain.com/upload-image

@description

	Request.Body = {
		Data: "[]byte",
		title: "String"
	}

@author Kunal Singh <kunalsin9h@gmail.com>
@type  Api Handler Function
*/
func UploadImage(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	if r.Method == "POST" {
		Data, err := io.ReadAll(r.Body)

		if err != nil {
			log.Printf("[UPLOAD-IMAGE] Error Reading Request Body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		var request UploadImageRequest
		err = json.Unmarshal(Data, &request)

		if err != nil {
			log.Printf("[UPLOAD-IMAGE] Error Parsing Request Body: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		if request.Data == nil {
			log.Printf("[UPLOAD-IMAGE] Missing Data (image []byte) Request Body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}

		uniqueImageId, err := uuid.New(7)

		if err != nil {
			log.Printf("[UPLOAD-IMAGE] Error Generating Unique Image Id: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		err = db.StoreImage(uniqueImageId, request.Title, request.Data)

		if err != nil {
			log.Printf("[STORE IMAGE] Error storing Data into db, %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		if os.Getenv("DOMAIN") == "" {
			log.Println("[UPLOAD-IMAGE] Domain is absent while preparing image url")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		if os.Getenv("PORT") == "" {
			log.Println("[UPLOAD-IMAGE] Port is absent while preparing image url")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		response := UploadImageResponse{
			URL: fmt.Sprintf("%s:%s/%s", os.Getenv("DOMAIN"), os.Getenv("PORT"), uniqueImageId),
		}

		jsonResponse, err := json.Marshal(response)

		if err != nil {
			log.Printf("[UPLOAD-IMAGE] Error while sending response Data: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
