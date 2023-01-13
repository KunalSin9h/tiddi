package api

import (
	"net/http"
	"time"
)

/*
Home
Endpoint: https://your-domain.com/

Renders the sample client

	from ../frontend

@author Kunal Singh <kunalsin9h@gmail.com>
@type  Api Handler Function
*/
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
	time.Sleep(time.Second * 5)
}
