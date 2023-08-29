package api

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kunalsin9h/tiddi/internal/db"
)

type FetchURL struct {
	FetchURL string
}

/*
Home
Endpoint: https://your-domain.com/

Renders the sample client

	from ../frontend

@author Kunal Singh <kunalsin9h@gmail.com>
@type  Api Handler Function
*/
func Home(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	/*
		Home, Show the sample client
	*/
	if r.URL.Path == "/" {
		// show the sample client
		// from ../frontend
		t, err := template.ParseFiles("./cmd/frontend/index.html")

		if err != nil {
			log.Printf("[HOME HTML-PARSING] Unable to parse html: %v", err)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Invalid Server Error"))
			return
		}

		fetchURL := fmt.Sprintf("%s/upload-image/", os.Getenv("HOST"))

		var htmlTemplateValues = FetchURL{
			FetchURL: fetchURL,
		}

		t.Execute(w, htmlTemplateValues)

	} else {
		// GET https://your-domain.com/93s9x_
		// Give the image with the id

		uiid, err := getUiid(r.URL.Path)

		if uiid == "favicon.ico" {
			return
		}

		if err != nil {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}

		image, err := db.GetImageBytes(uiid)

		if err != nil {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(image)

	}
}

func getUiid(url string) (string, error) {

	tokens := strings.Split(url, "/")

	paths := make([]string, 0)

	for _, token := range tokens {
		if token != "/" && token != " " {
			paths = append(paths, token)
		}
	}

	if len(paths) >= 2 {
		return paths[len(paths)-1], nil
	}
	return "", errors.New("invalid url")

}
