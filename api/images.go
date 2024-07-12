package handler

import "net/http"

func Image(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is array of image"))
}
