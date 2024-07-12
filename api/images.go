package handler

import (
	"encoding/json"
	"net/http"

	"github.com/PetrusAriaa/web-margot-backend/lib/gateway"
)

func Image(w http.ResponseWriter, r *http.Request) {
	g := gateway.NewGateway("images")
	res, err := g.Call()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	json.NewEncoder(w).Encode(res)
}
