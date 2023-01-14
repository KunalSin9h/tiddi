package api

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/znip-in/tiddi/src/db"
)

type PostURL struct {
	postURL string
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
		http.ServeFile(w, r, "./src/frontend/index.html")
		/*
			html, err := template.ParseFiles("./src/frontend/index.html")

			if err != nil {
				log.Printf("[HOME HTML] Unable to parse html from ./src/frontend/index.html: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}

			pr := PostURL{
				postURL: os.Getenv("HOST"),
			}

			html.Execute(w, pr)
		*/
	} else {
		// GET https://your-domain.com/93s9x_
		// Give the image with the id

		uiid, err := getUiid(r.URL.Path)

		if uiid == "favicon.ico" {
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}

		image, err := db.GetImageBytes(uiid)

		if err != nil {
			log.Printf("[GET-IMAGE] unable to get image form db: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(image)

	}
}

func getUiid(url string) (string, error) {

	uiid := strings.Split(url, "/")
	n := len(uiid)

	if n == 2 {
		return uiid[1], nil
	} else if n == 3 {
		if uiid[2] == "" {
			return uiid[1], nil
		} else {
			return "", errors.New("invalid url")
		}
	} else {
		return "", errors.New("invalid url")
	}
}
