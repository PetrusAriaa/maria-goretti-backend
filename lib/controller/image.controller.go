package controller

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (c *controllerData) GetImageList(w http.ResponseWriter, r *http.Request) {
	res := c.service.GetImageList()
	json.NewEncoder(w).Encode(res)
}

func (c *controllerData) GetImage(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	img, t := c.service.GetImage(p[len(p)-1])
	if img == nil {
		w.WriteHeader(404)
		w.Write([]byte("Not found"))
	}
	w.Header().Add("Content-Type", t)
	w.Write(img)
}