package controller

import (
	"encoding/json"
	"net/http"
)

func (c *controllerData) ImageLists(w http.ResponseWriter, r *http.Request) {
	res := c.service.GetImageList()
	json.NewEncoder(w).Encode(res)
}
