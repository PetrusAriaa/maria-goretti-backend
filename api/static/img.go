package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorMsg []string

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("fileName")
	if q == "" {
		_e := map[string]ErrorMsg{
			"errors": {"Missing required query param 'fileName'."},
		}
		w.WriteHeader(http.StatusBadRequest)
		e, err := json.Marshal(_e)
		if err != nil {
			log.Fatalln(err.Error())
		}
		w.Write(e)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(q))
}
