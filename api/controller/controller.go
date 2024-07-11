package controller

import "github.com/PetrusAriaa/web-margot-backend/api/service"

type controllerData struct {
	service *service.Service
}

func NewController(s *service.Service) *controllerData {
	return &controllerData{
		s,
	}
}
