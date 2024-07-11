package service

import (
	"context"

	"github.com/PetrusAriaa/web-margot-backend/src/db"
	"github.com/PetrusAriaa/web-margot-backend/src/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService(db *db.DBConnections, ctx context.Context) *Service {
	r := repository.NewRepository(db, ctx)
	return &Service{
		r,
	}
}
