package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"cloud.google.com/go/storage"
)

func (c *controllerData) GetImageList(w http.ResponseWriter, r *http.Request) {
	res, err := c.service.GetImageList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
	json.NewEncoder(w).Encode(res)
}

func (c *controllerData) GetImage(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	img, t, err := c.service.GetImage(p[len(p)-1])
	if err == storage.ErrObjectNotExist {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Header().Add("Content-Type", t)
	w.Write(img)
}
