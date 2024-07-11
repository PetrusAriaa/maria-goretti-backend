package controller

import (
	"encoding/json"
	"net/http"

	"github.com/PetrusAriaa/web-margot-backend/src/service"
)

type controllerData struct {
	service *service.Service
}

func NewController(s *service.Service) *controllerData {
	return &controllerData{
		s,
	}
}

func (c *controllerData) ImageLists(w http.ResponseWriter, r *http.Request) {
	res := c.service.GetImageList()
	json.NewEncoder(w).Encode(res)
}
