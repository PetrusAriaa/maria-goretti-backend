package handler

import "net/http"

func (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("might be an image"))
}