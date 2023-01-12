package api

import "net/http"

/*
UploadImage endpoint to upload an image with image data and title
Endpoint: POST https://your-domain.com/upload-image

@description
	Request.Body = {
		data: "[]byte",
		title: "String"
	}

@author Kunal Singh <kunalsin9h@gmail.com>
@type  Api Handler Function
*/
func UploadImage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Uploading..."))
}
